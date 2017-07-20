package tags

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"time"
)

type ModelParam struct {
	field *structs.Field
	param string
}

func (this ModelParam) FieldName() (string) {
	return this.field.Name()
}

func (this ModelParam) GetParam() (string) {
	return this.param
}

func (this ModelParam) GetField() (*structs.Field) {
	return this.field
}

type FieldParam interface {
	FieldName() string
	GetParam() string
	GetField() *structs.Field
}

var (
	defaultTagsRules = map[string]HandleTag{
		"default": RuleDefault,
	}

	reservedTags = []string{
		"json",
		"validate",
	}
)

func RuleDefault(_model FieldParam) () {

	if _model.GetParam() == "auto_now" {
		if _model.GetField().Kind() == reflect.ValueOf(time.Time{}).Kind() {

			_model.GetField().Set(time.Now())
		}
	}

	fmt.Println("Handle default " + _model.GetParam())
}
