package entities

import "time"

type Category struct {
  ID        uint
  Name      string
  CreatedAt time.Time
  UpdatedAt time.Time
}
