package models

import (
	arangoDB "github.com/hostelix/aranGO"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

type ProcedureSessionModel struct {
	arangoDB.Document

	ID           string        `json:"id" validate:"required" on_create:"set,auto_uuid"`
	ProcedureID  string        `json:"id" validate:"required" on_create:"set,auto_uuid"`
	UserID       string        `json:"user_id,omitempty" validate:"required"`
	CurrentStage int           `json:"stage"`

	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this ProcedureSessionModel) GetKey() string {
	return this.Key
}

func (this ProcedureSessionModel) GetCollection() string {
	return "procedure_session"
}

func (this ProcedureSessionModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this ProcedureSessionModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *ProcedureSessionModel) PreSave(c *arangoDB.Context) {

	v := validator.New()
	v.Validate(this)

	if v.IsValid() {
		//Tag Process for model
		t := tags.New()

		t.ProcessTags(this)
	} else {
		c.Err["error_validation"] = "Error validating model"
		this.ErrorsValidation = v.GetMessagesValidation()
	}

	return
}
