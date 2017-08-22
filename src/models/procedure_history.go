package models

import (
	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

type Stage struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCurrent   bool    `json:"is_current"`
}

type ProcedureHistoryModel struct {
	arangoDB.Document

	ProcedureID string        `json:"id" validate:"required" on_create:"set,auto_uuid"`
	UserID      string        `json:"user_id,omitempty" validate:"required"`
	Stages      []Stage      `json:"stage"`

	types.Timestamps
	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this ProcedureHistoryModel) GetKey() string {
	return this.Key
}

func (this ProcedureHistoryModel) GetCollection() string {
	return "procedure_history"
}

func (this ProcedureHistoryModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this ProcedureHistoryModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *ProcedureHistoryModel) PreSave(c *arangoDB.Context) {

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

/*
func (this ProcedureHistoryModel) GetCurrentStage() Stage {
	for _, value := range this.Stages {
		if value.IsCurrent {
			return value
		}
	}
}*/
