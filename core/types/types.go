package types

import (
	arangoDB "github.com/diegogub/aranGO"
	"time"
)

type JsonObject map[string]interface{}

type JsonArray []JsonObject

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseAPI struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type NeuronCEHDUN struct {
	arangoDB.Document

	Host    string         `json:"host"`
	Port    int            `json:"port"`
	Name    string         `json:"name"`
	Actions []ActionNeuron `json:"actions"`
}

type ActionNeuron struct {
	EndPoint    string                 `json:"end_point"`
	Params      map[string]interface{} `json:"params"`
	Method      string                 `json:"method"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
}

type Citizen struct {
	arangoDB.Document

	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	IdentityDocument string `json:"identity_document"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
}
