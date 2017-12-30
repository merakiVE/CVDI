package tags

import (
	"errors"

	goModel "gopkg.in/jeevatkm/go-model.v1"
	"github.com/fatih/structs"
	"github.com/fatih/structtag"
)

type HandleFuncTag func(i FieldParam)

type StructProcessorTag struct {
	rulesTag map[string]HandleFuncTag
}

/*
	Funcion para crear un nuevo procesor tag
	registra todas las rules dentro de la estructura de datos
 */
func New() (*StructProcessorTag) {
	sp := &StructProcessorTag{
		rulesTag: make(map[string]HandleFuncTag, len(defaultTagsRules)),
	}

	for key, fn := range defaultTagsRules {
		sp.RegisterHandleRule(key, fn)
	}

	return sp
}

/*
	Funcion que obtiene el manejador de la regla, osea la funcion que procesa la regla
 */
func (this StructProcessorTag) GetHandleRule(tag string) (HandleFuncTag) {

	v, ok := this.rulesTag[tag]

	if ok {
		return v
	}
	return nil
}

/*
	Funcion para registrar las reglas en el processor tag
 */
func (this *StructProcessorTag) RegisterHandleRule(tag string, fn HandleFuncTag) (error) {

	isTagRestricted := false

	//Verify if new tag handle rule is reserved
	for _, tag_value := range reservedTags {
		if tag_value == tag {
			isTagRestricted = true
			break
		}
	}

	if isTagRestricted {
		return errors.New("Name tag: " + tag + " is reserved")
	}

	if len(tag) == 0 {
		return errors.New("Function Key cannot be empty")
	}

	if fn == nil {
		return errors.New("Function cannot be empty")
	}

	this.rulesTag[tag] = fn

	return nil
}

func GetKeysTagField(model interface{}, fieldName string) ([]string) {

	tag, _ := goModel.Tag(model, fieldName)

	tags, err := structtag.Parse(string(tag))

	if err != nil {
		panic(err)
	}

	return tags.Keys()
}


func GetMapTagField(model interface{}, fieldName string) (map[string]*structtag.Tag) {

	map_field := make(map[string]*structtag.Tag, 0)

	tag, _ := goModel.Tag(model, fieldName)

	tags, err := structtag.Parse(string(tag))

	if err != nil {
		panic(err)
	}

	for _, v := range tags.Tags() {
		map_field[v.Key] = v
	}

	return map_field
}

func (this StructProcessorTag) ProcessTags(model interface{}) {

	modelFields := structs.Fields(model)

	for _, field := range modelFields {

		if field.IsEmbedded() {
			for _, fieldE := range field.Fields() {

				data_tags := GetMapTagField(model, fieldE.Name())

				for key, value_tag := range data_tags {

					cb := this.GetHandleRule(key)

					if cb != nil {
						cb(FieldParamStruct{
							field: fieldE,
							actionName: value_tag.Name,
							params: value_tag.Options,
						})
					}
				}
			}
		} else {
			data_tags := GetMapTagField(model, field.Name())

			for key, value_tag := range data_tags {

				cb := this.GetHandleRule(key)

				if cb != nil {
					cb(FieldParamStruct{
						field: field,
						actionName: value_tag.Name,
						params: value_tag.Options,
					})
				}
			}

		}
	}
}
