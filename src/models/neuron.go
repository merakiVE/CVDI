package models

import (
	arangoDB "github.com/hostelix/aranGO"
)

type Neuron struct {
	arangoDB.Document

	ID        string         `json:"id"`
	Host      string         `json:"host"`
	Port      int            `json:"port"`
	Name      string         `json:"name"`
	Actions   []ActionNeuron `json:"actions"`
	PublicKey string         `json:"public_key"`
}

func (this Neuron) GetKey() string {
	return this.Key
}

func (this Neuron) GetCollection() string {
	return "neuron"
}

func (this Neuron) GetError() (string, bool) {
	return this.Message, this.Error
}


type ActionNeuron struct {
	Name        string                 `json:"name"`
	EndPoint    string                 `json:"end_point"`
	Params      map[string]interface{} `json:"params"`
	Method      string                 `json:"method"`
	Description string                 `json:"description"`
}