package entities

import "time"

type Product struct {
  ID          uint
  Name        string
  Category    Category
  Stock       uint
  Price       uint
  Description string
  CreatedAt   time.Time
  UpdatedAt   time.Time
}
