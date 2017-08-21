package models

import (
	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

//Structs for BPMN
type Lane struct {
	Name       string        `json:"name"`
	InPool     bool          `json:"in_pool"`
	NamePool   string        `json:"name_pool"`
	Activities []Activity    `json:"activities,omitempty" validate:"required"`
}

type Bpmn struct {
	Lanes []Lane `json:"lanes"`
}

type Activity struct {
	NeuronKey string `json:"neuron_key"`
	ActionID  string `json:"action_id"`
	Sequence  int32    `json:"sequence"`
}

type Stage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProcedureModel struct {
	arangoDB.Document

	ID     string        `json:"id" validate:"required" on_create:"set,auto_uuid"`
	Owner  string        `json:"owner,omitempty" validate:"required"`
	Stages []Stage       `json:"stages"`

	types.Timestamps
	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this ProcedureModel) GetKey() string {
	return this.Key
}

func (this ProcedureModel) GetCollection() string {
	return "procedures"
}

func (this ProcedureModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this ProcedureModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *ProcedureModel) PreSave(c *arangoDB.Context) {

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
