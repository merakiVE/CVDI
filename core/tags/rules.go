package tags

import (
	"github.com/fatih/structs"
	"reflect"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/merakiVE/CVDI/core/utils"
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

/*
	Regla Default Tag
	Con esta regla se pueden setear valores por defecto en los campos que se le especifique dicha tag
	Uso:
		NameField string `default:"valor por defecto"`
 */
func RuleDefault(f FieldParam) () {
	if f.GetAction() == "auto_now" {
		if f.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {
			f.GetField().Set(time.Now())
		}
	}
}

/*
	Rule que se ejecutan al momento de guardar el registro por primera vez en al db
 */
func RuleOnCreate(f FieldParam) () {

	switch f.GetAction() {

	case "execute":
		{
			// Testing function
		}
	case "set":
		{
			params := f.GetParams()
			/*
				Accion que auto genera un timestamps y lo setea al field
			 */
			if params[0] == "auto_now" {
				if f.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {
					f.GetField().Set(time.Now())
				}
			}
			/*
				Accion que auto genera un uuid y lo setea al field
			 */

			if params[0] == "auto_uuid" {

				if f.GetField().Kind() == reflect.String {
					//Generate uuid and set value to field
					f.GetField().Set(utils.GenerateUUIDV4())
				}
			}
		}
	/*
		Accion que hashea el password y lo setea
	 */
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
