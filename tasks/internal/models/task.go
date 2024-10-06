package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID          int64     `bun:"id,pk,autoincrement"`
	Titile      string    `bun:"title,notnull"`
	Description string    `bun:"description,notnull"`
	Status      string    `bun:"status,notnull"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
