package types

import "time"

type userType string

const (
	UserTypeAdmin    userType = "admin"
	UserTypeStandard userType = "standard"
)

type User struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	FirstName string     `validate:"required" json:"first_name" xorm:"first_name"`
	LastName  string     `validate:"required" json:"last_name" xorm:"last_name"`
	Email     string     `validate:"required" json:"email" xorm:"email"`
	Password  string     `validate:"required" json:"-" xorm:"password"`
	Type      userType   `validate:"required" json:"-" xorm:"user_type"`
	CreatedAt time.Time  `json:"created_at" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" xorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted_at"`
}

func (*User) TableName() string {
	return `users`
}
