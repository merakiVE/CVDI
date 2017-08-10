package models

import (
	arangoDB "github.com/hostelix/aranGO"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

type NeuronModel struct {
	arangoDB.Document

	ID        string         `json:"id"`
	Host      string         `json:"host"`
	Port      int            `json:"port"`
	Name      string         `json:"name"`
	Actions   []ActionNeuron `json:"actions"`
	PublicKey string         `json:"public_key"`

	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this NeuronModel) GetKey() string {
	return this.Key
}

func (this NeuronModel) GetCollection() string {
	return "neurons"
}

func (this NeuronModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this NeuronModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *NeuronModel) PreSave(c *arangoDB.Context) {

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

type ActionNeuron struct {
	ID          string                 `json:"id" on_create:"set,auto_uuid"`
	Name        string                 `json:"name"`
	EndPoint    string                 `json:"end_point"`
	Params      map[string]interface{} `json:"params"`
	Method      string                 `json:"method"`
	Description string                 `json:"description"`
}
