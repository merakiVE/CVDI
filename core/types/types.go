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

type ResponseAPI struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
