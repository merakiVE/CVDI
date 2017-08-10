package tags

import (
	"github.com/fatih/structs"
	"reflect"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

type FieldParamStruct struct {
	field      *structs.Field
	actionName string
	params     []string
}

func (this FieldParamStruct) FieldName() (string) {
	return this.field.Name()
}

func (this FieldParamStruct) GetParams() ([]string) {
	return this.params
}

func (this FieldParamStruct) GetAction() (string) {
	return this.actionName
}

func (this FieldParamStruct) GetField() (*structs.Field) {
	return this.field
}

type FieldParam interface {
	FieldName() string
	GetParams() []string
	GetAction() string
	GetField() *structs.Field
}

var (
	defaultTagsRules = map[string]HandleFuncTag{
		"default":   RuleDefault,
		"on_create": RuleOnCreate,
	}

	reservedTags = []string{
		"json",
		"validate",
	}
)

func RuleDefault(f FieldParam) () {

	if f.GetAction() == "auto_now" {
		if f.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {

			f.GetField().Set(time.Now())
		}
	}
}

func RuleOnCreate(f FieldParam) () {

	switch f.GetAction() {

	case "execute":
		{

			// Testing function
		}

	case "set":
		{

			params := f.GetParams()

			if params[0] == "auto_now" {

				if f.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {

					f.GetField().Set(time.Now())
				}
			}

			if params[0] == "auto_uuid" {

				if f.GetField().Kind() == reflect.String {

					//Generate uuid and set value to field
					f.GetField().Set(uuid.NewV4())
				}
			}

		}

	case "make_password":
		{

			if f.GetField().Kind() == reflect.String {
				bytes, err := bcrypt.GenerateFromPassword([]byte(f.GetField().Value().(string)), 10)

				if err != nil {
					panic(err)
				}

				f.GetField().Set(string(bytes))
			}
		}
	}
}
