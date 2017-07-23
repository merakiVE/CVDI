package tags

import (
	"github.com/fatih/structs"
	"reflect"
	"time"
	"golang.org/x/crypto/bcrypt"

)

type ModelParam struct {
	field      *structs.Field
	nameAction string
	params     []string
}

func (this ModelParam) FieldName() (string) {
	return this.field.Name()
}

func (this ModelParam) GetParams() ([]string) {
	return this.params
}

func (this ModelParam) GetAction() (string) {
	return this.nameAction
}

func (this ModelParam) GetField() (*structs.Field) {
	return this.field
}

type FieldParam interface {
	FieldName() string
	GetParams() []string
	GetAction() string
	GetField() *structs.Field
}

var (
	defaultTagsRules = map[string]HandleTag{
		"default":   RuleDefault,
		"on_create": RuleOnCreate,
	}

	reservedTags = []string{
		"json",
		"validate",
	}
)

func RuleDefault(_model FieldParam) () {

	if _model.GetAction() == "auto_now" {
		if _model.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {

			_model.GetField().Set(time.Now())
		}
	}
}

func RuleOnCreate(f FieldParam) () {

	switch f.GetAction() {

		case "execute" : {

			// Testing function
		}

		case "set" : {

			params := f.GetParams()

			if params[0] == "auto_now" {

				if f.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {

					f.GetField().Set(time.Now())
				}
			}
		}

		case "make_password" : {

			if f.GetField().Kind() == reflect.String {
				bytes, err := bcrypt.GenerateFromPassword([]byte(f.GetField().Value().(string)), 14)
				
				if err != nil {
					panic(err)
				}

				f.GetField().Set(string(bytes))
			}
		}
	}
}
