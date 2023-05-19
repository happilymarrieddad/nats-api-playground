package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	FirstName string     `validate:"required" json:"first_name" xorm:"first_name"`
	LastName  string     `validate:"required" json:"last_name" xorm:"last_name"`
	Email     string     `validate:"required" json:"email" xorm:"email"`
	Password  string     `json:"-" xorm:"password"`
	CreatedAt time.Time  `json:"created_at" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (*User) TableName() string {
	return `users`
}

func (u *User) SetPassword(psw string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) PasswordMatches(psw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(psw)) == nil
}

type UserUpdate struct {
	ID        int64     `validate:"required" json:"id" xorm:"'id' pk autoincr"`
	FirstName *string   `json:"first_name" xorm:"first_name"`
	LastName  *string   `json:"last_name" xorm:"last_name"`
	UpdatedAt time.Time `json:"-" xorm:"updated_at"`
}

func (*UserUpdate) TableName() string {
	return `users`
}
