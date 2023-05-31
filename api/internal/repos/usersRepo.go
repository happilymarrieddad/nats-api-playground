package repos

import (
	"context"
	"fmt"
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

//go:generate mockgen -destination=./mocks/Users.go -package=mock_repos plant-ecommerce-api/internal/repos Users
type Users interface {
	Find(ctx context.Context, opts *UserFindOpts) ([]*types.User, int64, error)
	Create(context.Context, *types.User) error
	FindByID(ctx context.Context, id int64) (*types.User, bool, error)
	FindByEmail(ctx context.Context, email string) (*types.User, bool, error)
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

	handleOpts := func(useLimit bool) *xorm.Session {
		sesh := r.db.NewSession()

		// handle pagination
		if useLimit && opts.Limit > 0 {
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

		return sesh
	}

	usrs = []*types.User{}
	if err = handleOpts(true).Find(&usrs); err != nil {
		return nil, 0, err
	}

	if count, err = handleOpts(false).Count(&types.User{}); err != nil {
		return nil, 0, err
	}

	return usrs, count, nil
}

func (r *users) FindByEmail(ctx context.Context, email string) (*types.User, bool, error) {
	user := &types.User{Email: email}

	if exists, err := r.db.Context(ctx).Get(user); err != nil {
		return nil, false, err
	} else if !exists {
		return nil, false, types.NewNotFoundError(fmt.Sprintf("user with email '%s' not found", email))
	}

	return user, true, nil
}

func (r *users) FindByID(ctx context.Context, id int64) (*types.User, bool, error) {
	user := &types.User{}
	if exists, err := r.db.Context(ctx).ID(id).Get(user); err != nil {
		return nil, false, err
	} else if !exists {
		return nil, false, types.NewNotFoundError(fmt.Sprintf("user with id '%d' not found", id))
	}

	return user, true, nil
}

func (r *users) Create(ctx context.Context, newUsr *types.User) error {
	if err := types.Validate(newUsr); err != nil {
		return err
	}

	t := time.Now()
	newUsr.CreatedAt = t
	newUsr.Type = types.UserTypeStandard

	if _, err := r.db.Context(ctx).Insert(newUsr); err != nil {
		return err
	}

	return nil
}
