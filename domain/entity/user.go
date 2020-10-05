package entity

import (
	"fmt"
	"regexp"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) (*User, error) {
	matchedStr, err := regexp.MatchString(`/^[a-z ,.'-]+$/i`, name)
	if false == matchedStr || err != nil {
		return nil, fmt.Errorf("invalid name")
	}
	matched, emailErr := regexp.MatchString(`\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`, email)
	if false == matched || emailErr != nil {
		return nil, fmt.Errorf("invalid email")
	}
	return &User{
		Name:  name,
		Email: email,
	}, nil
}
