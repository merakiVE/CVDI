package types

import (
	"time"
	arangoDB "github.com/diegogub/aranGO"
)

func init() {}

type JsonObject map[string]interface{}
type JsonArray []JsonObject

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

type User struct {
	arangoDB.Document

	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Token     string    `json:"token"`
	LastLogin time.Time `json:"last_login"`
}

type Citizen struct {
	arangoDB.Document

	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	IdentityDocument string `json:"identity_document"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	User             User   `json:"user"`
}

type ResponseAPI struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
