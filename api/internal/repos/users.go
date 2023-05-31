package repos

import (
	"context"
	"time"

	"github.com/happilymarrieddad/nats-api-playground/api/types"
	"xorm.io/xorm"
)

type UserFindOpts struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	// Specific fields
	ID        *int64  `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
}

//go:generate mockgen -destination=./mocks/Users.go -package=mock_repos github.com/happilymarrieddad/nats-api-playground/api/internal/repos Users
type Users interface {
	Find(ctx context.Context, opts *UserFindOpts) ([]*types.User, int64, error)
	Get(ctx context.Context, id int64) (*types.User, bool, error)
	GetByEmail(ctx context.Context, email string) (*types.User, bool, error)
	Create(ctx context.Context, usr *types.User) error
	CreateTx(ctx context.Context, sesh *xorm.Session, usr *types.User) error
	Update(ctx context.Context, usr types.UserUpdate) (*types.User, error)
	UpdateTx(ctx context.Context, sesh *xorm.Session, usr types.UserUpdate) (*types.User, error)
	Delete(ctx context.Context, id int64) error
	DeleteTx(ctx context.Context, sesh *xorm.Session, id int64) error
}

func NewUsers(db *xorm.Engine) Users {
	return &users{db: db}
}

type users struct {
	db *xorm.Engine
}

func (r *users) Find(ctx context.Context, opts *UserFindOpts) (usrs []*types.User, count int64, err error) {
	if opts == nil {
		opts = &UserFindOpts{Limit: 10}
	}

	// Just a safety check
	if opts.Limit < 1 || opts.Limit > 100 {
		opts.Limit = 10
	}

	sesh := r.db.NewSession()

	// handle pagination
	if opts.Limit > 0 {
		if opts.Offset > 0 {
			sesh = sesh.Limit(opts.Limit, opts.Offset)
		} else {
			sesh = sesh.Limit(opts.Limit)
		}
	}

	if opts.ID != nil {
		sesh = sesh.ID(*opts.ID)
	}
	if opts.FirstName != nil {
		sesh = sesh.And("first_name = ?", *opts.FirstName)
	}
	if opts.LastName != nil {
		sesh = sesh.And("last_name = ?", *opts.LastName)
	}
	if opts.Email != nil {
		sesh = sesh.And("email = ?", *opts.Email)
	}

	usrs = []*types.User{}
	if count, err = sesh.Context(ctx).FindAndCount(&usrs); err != nil {
		return nil, 0, err
	}

	return usrs, count, nil
}

func (r *users) Get(ctx context.Context, id int64) (*types.User, bool, error) {
	usr := &types.User{}

	exists, err := r.db.Context(ctx).ID(id).Get(usr)
	if err != nil {
		return nil, false, err
	} else if !exists {
		usr = nil
	}

	return usr, exists, nil
}

func (r *users) GetByEmail(ctx context.Context, email string) (*types.User, bool, error) {
	usr := &types.User{}

	exists, err := r.db.Context(ctx).Where("email = ?", email).Get(usr)
	if err != nil {
		return nil, false, err
	} else if !exists {
		usr = nil
	}

	return usr, exists, nil
}

func (r *users) Create(ctx context.Context, usr *types.User) error {
	return wrapInTx(r.db, func(s *xorm.Session) error {
		return r.CreateTx(ctx, s, usr)
	})
}

func (r *users) CreateTx(ctx context.Context, sesh *xorm.Session, usr *types.User) error {
	if sesh == nil {
		return errTransaction
	}

	usr.CreatedAt = time.Now()

	if _, err := sesh.Context(ctx).Insert(usr); err != nil {
		return err
	}

	return nil
}

func (r *users) Update(ctx context.Context, usr types.UserUpdate) (rtUsr *types.User, err error) {
	err = wrapInTx(r.db, func(s *xorm.Session) error {
		rtUsr, err = r.UpdateTx(ctx, s, usr)
		return err
	})
	return
}

func (r *users) UpdateTx(ctx context.Context, sesh *xorm.Session, usr types.UserUpdate) (*types.User, error) {
	if sesh == nil {
		return nil, errTransaction
	}

	usr.UpdatedAt = time.Now()

	if _, err := r.db.Context(ctx).ID(usr.ID).Update(&usr); err != nil {
		return nil, err
	}

	// Don't bother returning the exists... we already know it does. Update will fail if not
	newUsr, _, err := r.Get(ctx, usr.ID)
	return newUsr, err
}

func (r *users) Delete(ctx context.Context, id int64) error {
	return wrapInTx(r.db, func(s *xorm.Session) error {
		return r.DeleteTx(ctx, s, id)
	})
}

func (r *users) DeleteTx(ctx context.Context, sesh *xorm.Session, id int64) error {
	if sesh == nil {
		return errTransaction
	}

	_, err := sesh.Context(ctx).ID(id).Delete(&types.User{})
	return err
}
