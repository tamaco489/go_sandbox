package model

import "time"

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	ProviderUID string    `json:"provider_uid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUser(id int64, name, email, role, status, providerUID string) *User {
	now := time.Now()
	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		Role:        role,
		Status:      status,
		ProviderUID: providerUID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
