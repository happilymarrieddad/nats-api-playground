package repos

import (
	"sync"

	"xorm.io/xorm"
)

var singular GlobalRepo

//go:generate mockgen -destination=./mocks/GlobalRepo.go -package=mock_repos plant-ecommerce-api/internal/repos GlobalRepo
type GlobalRepo interface {
	// Categories() Categories
	// Customers() Customers
	// Products() Products
	// Users() Users
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

func (g *globalRepo) getFactory(key string, fn func() interface{}) interface{} {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	val, exists := g.fac[key]
	if exists {
		return val
	}

	newFac := fn()
	g.fac[key] = newFac

	return newFac
}

// func (g *globalRepo) Categories() Categories {
// 	return g.getFactory("Categories", func() interface{} { return NewCategories(g.db) }).(Categories)
// }

// func (g *globalRepo) Customers() Customers {
// 	return g.getFactory("Customers", func() interface{} { return NewCustomers(g.db) }).(Customers)
// }

// func (g *globalRepo) Products() Products {
// 	return g.getFactory("Products", func() interface{} { return NewProducts(g.db) }).(Products)
// }

// func (g *globalRepo) Users() Users {
// 	return g.getFactory("Users", func() interface{} { return NewUsers(g.db) }).(Users)
// }
