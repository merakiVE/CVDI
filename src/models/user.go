package models

import (
	"time"

	arangoDB "github.com/diegogub/aranGO"

	//"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
)

const (
	collectionName = "users"
)

type UserModel struct {
	arangoDB.Document

	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Token     string    `json:"token"`
	LastLogin time.Time `json:"last_login"`

	types.Timestamps
}

func (this UserModel) GetKey() string {
	return this.Key
}

func (this UserModel) GetCollection() string {
	return collectionName
}

func (this UserModel) GetError() (string, bool) {
	return this.Message, this.Error
}
