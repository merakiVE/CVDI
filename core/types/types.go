package types

import (
	arangoDB "github.com/hostelix/aranGO"
	"time"
)

type JsonObject map[string]interface{}

type JsonArray []JsonObject

type Timestamps struct {
	CreatedAt time.Time `json:"created_at" on_create:"set,auto_now"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this *Timestamps) AutoNowCreate() {
	this.CreatedAt = time.Now()
}

func (this *Timestamps) AutoNowUpdate() {
	this.UpdatedAt = time.Now()
}

type ResponseAPI struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type Citizen struct {
	arangoDB.Document

	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	IdentityDocument string `json:"identity_document"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
