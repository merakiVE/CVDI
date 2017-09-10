package models

import (
	"time"

	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

type UserModel struct {
	arangoDB.Document

	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username,omitempty" validate:"required"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required" on_create:"make_password"`
	Token     string    `json:"token,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`

	types.Timestamps
	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this UserModel) GetKey() string {
	return this.Key
}

func (this UserModel) GetCollection() string {
	return "users"
}

func (this UserModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this UserModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *UserModel) PreSave(c *arangoDB.Context) {

	v := validator.New()

	v.Validate(this)

	if v.IsValid() {

		//Tag Process for model
		tags.New().ProcessTags(this)
	} else {

		c.Err["error_validation"] = "Error validating model"
		this.ErrorsValidation = v.GetMessagesValidation()
	}

	return
}
