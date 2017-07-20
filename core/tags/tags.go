package tags

import (
	"fmt"
	"errors"
	"strings"
	"github.com/fatih/structs"
	"gopkg.in/jeevatkm/go-model.v1"
)

var (
	reservedTags []string
)

type HandleTag func(i interface{})

type StructProcessorTag struct {
	rulesTag map[string]HandleTag
}

func New() (*StructProcessorTag) {
	sp := &StructProcessorTag{
		rulesTag: make(map[string]HandleTag, len(tagsDefaultRules)),
	}

	for _key, _fn := range tagsDefaultRules {
		sp.RegisterHandleRule(_key, _fn)
	}

	return sp
}

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

func GetKeysTagField(_model interface{}, _fieldName string) ([]string) {
	keys := make([]string, 0)

	tag, _ := model.Tag(_model, _fieldName)

	for _, v := range strings.Split(string(tag), " ") {
		value := strings.Split(v, ":")
		keys = append(keys, value[0])
	}

	return keys
}

func GetMapTagField(_model interface{}, _fieldName string) (map[string]string) {
	map_field := make(map[string]string, 0)

	tag, _ := model.Tag(_model, _fieldName)

	for _, v := range strings.Split(string(tag), " ") {
		value := strings.Split(v, ":")
		map_field[value[0]] = value[1]
	}

	return map_field
}

func (this StructProcessorTag) ProcessTags(_model interface{}) {

	modelFields := structs.Fields(_model)

	for _, field := range modelFields {

		if field.IsEmbedded() {
			for _, fieldE := range field.Fields() {

				//fmt.Println(fieldE.Name())

				tx, _ := model.Tag(_model, fieldE.Name())
				fmt.Println(fieldE.Name(), tx)
			}
		} else {
			keys_tag := GetKeysTagField(_model, field.Name())

			for _, key := range keys_tag {

				fmt.Println(GetMapTagField(_model, field.Name()))
				cb := this.GetHandleRule(key)

				if cb != nil {
					cb(_model)
				}

			}
		}
	}
}
