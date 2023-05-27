package repos

import (
	"time"

	"github.com/happilymarrieddad/nats-api-playground/api/types"
	"xorm.io/xorm"
)

//go:generate mockgen -destination=./mocks/Users.go -package=mock_repos github.com/happilymarrieddad/nats-api-playground/api/internal/repos Users
type Users interface {
	Find(limit, offset int) ([]*types.User, int64, error)
	Get(id int64) (*types.User, bool, error)
	GetByEmail(email string) (*types.User, bool, error)
	Create(usr *types.User) error
	CreateTx(sesh *xorm.Session, usr *types.User) error
	Update(usr types.UserUpdate) (*types.User, error)
	UpdateTx(sesh *xorm.Session, usr types.UserUpdate) (*types.User, error)
	Delete(id int64) error
	DeleteTx(sesh *xorm.Session, id int64) error
}

func NewUsers(db *xorm.Engine) Users {
	return &users{db: db}
}

type users struct {
	db *xorm.Engine
}

func (r *users) Find(limit, offset int) ([]*types.User, int64, error) {
	usrs := []*types.User{}

	count, err := r.db.Limit(limit, offset).FindAndCount(&usrs)
	if err != nil {
		return nil, 0, err
	}

	return usrs, count, nil
}

func (r *users) Get(id int64) (*types.User, bool, error) {
	usr := &types.User{}

	exists, err := r.db.ID(id).Get(usr)
	if err != nil {
		return nil, false, err
	} else if !exists {
		usr = nil
	}

	return usr, exists, nil
}

func (r *users) GetByEmail(email string) (*types.User, bool, error) {
	usr := &types.User{}

	exists, err := r.db.Where("email = ?", email).Get(usr)
	if err != nil {
		return nil, false, err
	} else if !exists {
		usr = nil
	}

	return usr, exists, nil
}

func (r *users) Create(usr *types.User) error {
	return wrapInTx(r.db, func(s *xorm.Session) error {
		return r.CreateTx(s, usr)
	})
}

func (r *users) CreateTx(sesh *xorm.Session, usr *types.User) error {
	if sesh == nil {
		return errTransaction
	}

	usr.CreatedAt = time.Now()

	if _, err := sesh.Insert(usr); err != nil {
		return err
	}

	return nil
}

func (r *users) Update(usr types.UserUpdate) (rtUsr *types.User, err error) {
	err = wrapInTx(r.db, func(s *xorm.Session) error {
		rtUsr, err = r.UpdateTx(s, usr)
		return err
	})
	return
}

func (r *users) UpdateTx(sesh *xorm.Session, usr types.UserUpdate) (*types.User, error) {
	if sesh == nil {
		return nil, errTransaction
	}

	usr.UpdatedAt = time.Now()

	if _, err := r.db.ID(usr.ID).Update(&usr); err != nil {
		return nil, err
	}

	// Don't bother returning the exists... we already know it does. Update will fail if not
	newUsr, _, err := r.Get(usr.ID)
	return newUsr, err
}

func (r *users) Delete(id int64) error {
	return wrapInTx(r.db, func(s *xorm.Session) error {
		return r.DeleteTx(s, id)
	})
}

func (r *users) DeleteTx(sesh *xorm.Session, id int64) error {
	if sesh == nil {
		return errTransaction
	}

	_, err := sesh.ID(id).Delete(&types.User{})
	return err
}
