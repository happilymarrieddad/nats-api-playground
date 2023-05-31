package repos

import (
	"sync"

	"xorm.io/xorm"
)

var singular GlobalRepo

//go:generate mockgen -destination=./mocks/GlobalRepo.go -package=mock_repos github.com/happilymarrieddad/nats-api-playground/api/internal/repos GlobalRepo
type GlobalRepo interface {
	Users() Users
}

func NewGlobalRepo(db *xorm.Engine) (GlobalRepo, error) {
	if singular == nil {
		singular = &globalRepo{
			db:    db,
			mutex: &sync.RWMutex{},
			fac:   make(map[string]interface{}),
		}
	}

	return singular, nil
}

type globalRepo struct {
	db    *xorm.Engine
	mutex *sync.RWMutex
	fac   map[string]interface{}
}

func (g *globalRepo) getFactory(key string, fn func(db *xorm.Engine) interface{}) interface{} {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	val, exists := g.fac[key]
	if exists {
		return val
	}

	newFac := fn(g.db)
	g.fac[key] = newFac

	return newFac
}

func (g *globalRepo) Users() Users {
	return g.getFactory("Users", func(db *xorm.Engine) interface{} { return NewUsers(db) }).(Users)
}
