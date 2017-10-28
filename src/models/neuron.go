package models

import (
	"errors"

	arangoDB "github.com/hostelix/aranGO"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
)

type NeuronModel struct {
	arangoDB.Document

	ID   string         `json:"id" validate:"required" on_create:"set,auto_uuid"`
	Host string         `json:"host" validate:"url,required"`
	Port int            `json:"port" validate:"required"`
	Name string         `json:"name" validate:"required"`
	// option dive is for validate element to array
	Actions   []ActionNeuron `json:"actions" validate:"required,dive,required"`
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

func (this NeuronModel) GetAction(id string) (ActionNeuron, error) {
	for _, an := range this.Actions {
		if an.ID == id {
			return an, nil
		}
	}
	return ActionNeuron{}, errors.New("Action ID not found")
}

func (this *NeuronModel) PreSave(c *arangoDB.Context) {

	v := validator.New()
	v.Validate(this)

	if v.IsValid() {
		//Tag Process for model
		t := tags.New()

		t.ProcessTags(this)

		//Process tag to struct slice
		for i := range this.Actions {
			t.ProcessTags(&this.Actions[i])
		}
	} else {
		c.Err["error_validation"] = "Error validating model"
		this.ErrorsValidation = v.GetMessagesValidation()
	}

	return
}

type ActionNeuron struct {
	ID          string                 `json:"id" on_create:"set,auto_uuid"`
	Name        string                 `json:"name" validate:"required"`
	EndPoint    string                 `json:"end_point" validate:"required"`
	Params      map[string]string      `json:"params"`
	Method      string                 `json:"method" validate:"required"`
	Description string                 `json:"description" validate:"required"`
	Help        string                 `json:"help"`
}

func (this *ActionNeuron) Validate() []map[string]string {

	v := validator.New()
	v.Validate(this)

	if !v.IsValid() {
		return v.GetMessagesValidation()
	}

	tags.New().ProcessTags(this)

	return nil
}
