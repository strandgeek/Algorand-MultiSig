package middlewares

import "gorm.io/gorm"

type Middlewares struct {
	db *gorm.DB
}

func NewMiddlewares(db *gorm.DB) *Middlewares {
	return &Middlewares{
		db: db,
	}
}
