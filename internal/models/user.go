package models

import (
	"fmt"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:text;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: User{},
	}
}

func (b *UserBuilder) WithID(id uint) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) WithToken(token string) *UserBuilder {
	b.user.Token = token
	return b
}

func (b *UserBuilder) Build() (*User, error) {
	if b.user.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if b.user.Email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if b.user.Password == "" {
		return nil, fmt.Errorf("password is required")
	}
	return &b.user, nil
}
