// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"golang-boilerplate/ent"
)

type CreateUserInput struct {
	Name string `json:"name"`
}

type UserQueries struct {
	List *ent.UserConnection `json:"list"`
}
