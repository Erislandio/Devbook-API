package models

import "time"

// User representa usuário na rede social
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
