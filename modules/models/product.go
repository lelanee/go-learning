package models

import (
	"time"
)

type Product struct {
	ID         uint
	Name       string
	Price      float64
	CategoryID uint
	Category   Category
	CreatedAt  time.Time
}
