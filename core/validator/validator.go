package validator

import (
	"fmt"
	"strings"
	//"encoding/json"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	uniTranslator *ut.UniversalTranslator
	validate      *validator.Validate
	translate     ut.Translator
)

func init() {
	//Validator
	lang_en := en.New()
	uniTranslator = ut.New(lang_en, lang_en)
	translate, _ = uniTranslator.GetTranslator("en")
	validate = validator.New()

	en_translations.RegisterDefaultTranslations(validate, translate)
}

type StructValidator struct {
	MessagesValidation []map[string]string `json:"messages_validation"`
	ValidationSuccess  bool `json:"is_valid"`
}

func New() (*StructValidator) {
	return &StructValidator{}
}

func (this *StructValidator) Validate(_struct interface{}) (err error) {

	_errorValidationStruct := validate.Struct(_struct)
	_messages := make([]map[string]string, 0)
	this.ValidationSuccess = false

	if _errorValidationStruct != nil {

		if _, ok := _errorValidationStruct.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range _errorValidationStruct.(validator.ValidationErrors) {
			//Get namespace exmple:
			// NameStruct.Field => Person.Name
			// NameStruct.SliceField[0].Field => Person.Hobbies[0].Name
			struct_name := err.Namespace()

			//Split namespace
			//[0] => Name Struct
			//[1] => Name Field
			//[3] => Name Field Embeded
			key_field_error := strings.Join(strings.Split(struct_name, ".")[1:], ".")

			_messages = append(_messages, map[string]string{
				strings.ToLower(key_field_error): err.Translate(translate),
			})
		}

	} else {
		this.ValidationSuccess = true
	}

	this.MessagesValidation = _messages

	return _errorValidationStruct
}

func (this StructValidator) GetMessagesValidation() ([]map[string]string) {
	return this.MessagesValidation
}

func (this StructValidator) IsValid() (bool) {
	return this.ValidationSuccess
}
