package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/notprQn/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func dataBaseUserToUser(DbUser database.User) User {
	return User{
		ID:        DbUser.ID,
		CreatedAt: DbUser.CreatedAt,
		UpdatedAt: DbUser.UpdatedAt,
		Name:      DbUser.Name,
	}
}
