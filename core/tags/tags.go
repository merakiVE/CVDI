package tags

import (
	"errors"

	"gopkg.in/jeevatkm/go-model.v1"
	"github.com/fatih/structs"
	"github.com/fatih/structtag"
)

type HandleTag func(i FieldParam)

type StructProcessorTag struct {
	rulesTag map[string]HandleTag
}

func New() (*StructProcessorTag) {
	sp := &StructProcessorTag{
		rulesTag: make(map[string]HandleTag, len(defaultTagsRules)),
	}

	for _key, _fn := range defaultTagsRules {
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

	isTagRestricted := false

	//Verify if new tag handle rule is reserved
	for _, tag_value := range reservedTags {
		if tag_value == _tag {
			isTagRestricted = true
			break
		}
	}

	if isTagRestricted {
		return errors.New("Name tag: " + _tag + " is reserved")
	}

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

	tag, _ := model.Tag(_model, _fieldName)

	tags, err := structtag.Parse(string(tag))

	if err != nil {
		panic(err)
	}

	return tags.Keys()
}


func GetMapTagField(_model interface{}, _fieldName string) (map[string]*structtag.Tag) {

	map_field := make(map[string]*structtag.Tag, 0)

	tag, _ := model.Tag(_model, _fieldName)

	tags, err := structtag.Parse(string(tag))

	if err != nil {
		panic(err)
	}

	for _, v := range tags.Tags() {
		map_field[v.Key] = v
	}

	return map_field
}

func (this StructProcessorTag) ProcessTags(_model interface{}) {

	modelFields := structs.Fields(_model)

	for _, field := range modelFields {

		if field.IsEmbedded() {
			for _, fieldE := range field.Fields() {

				data_tags := GetMapTagField(_model, fieldE.Name())

				for key, value_tag := range data_tags {

					cb := this.GetHandleRule(key)

					if cb != nil {
						cb(ModelParam{
							field: fieldE,
							nameAction: value_tag.Name,
							params: value_tag.Options,
						})
					}
				}
			}
		} else {
			data_tags := GetMapTagField(_model, field.Name())

			for key, value_tag := range data_tags {

				cb := this.GetHandleRule(key)

				if cb != nil {
					cb(ModelParam{
						field: field,
						nameAction: value_tag.Name,
						params: value_tag.Options,
					})
				}
			}

		}
	}
}
