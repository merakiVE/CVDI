package tags

import (
	"fmt"
	"errors"
	"github.com/fatih/structs"
	//"reflect"
)

type HandleTag func(i interface{}) interface{}

type StructProcessorTag struct {
	rulesTag map[string]HandleTag
}

/*
func Tags(s interface{}) (map[string]reflect.StructTag, error) {
	sv, err := structValue(s)
	if err != nil {
		return nil, err
	}

	tags := map[string]reflect.StructTag{}

	fields := modelFields(sv)
	for _, f := range fields {
		tags[f.Name] = f.Tag
	}

	return tags, nil
}*/


func (this StructProcessorTag) GetHandleRule(_tag string) (HandleTag) {

	v, ok := this.rulesTag[_tag]

	if ok {
		return v
	}
	return nil
}

func (this *StructProcessorTag) RegisterHandleRule(_tag string, _fn HandleTag) (error) {

	if len(_tag) == 0 {
		return errors.New("Function Key cannot be empty")
	}

	if _fn == nil {
		return errors.New("Function cannot be empty")
	}

	this.rulesTag[_tag] = _fn

	return nil
}

func ProcessTags(_model interface{}) {

	modelFields := structs.Fields(_model)

	for _, field := range modelFields {

		if field.IsEmbedded() {
			for _, fieldE := range field.Fields() {
				fmt.Println(fieldE.Name())
			}
		}else {
		}
	}
}

/*func ExecuteHandleRuleTag() {}*/
