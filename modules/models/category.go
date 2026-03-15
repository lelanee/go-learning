package models

import "time"

type Category struct {
	ID        uint
	Name      string
	Products  []Product
	CreatedAt time.Time
}
