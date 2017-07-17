package validator

import (
	"fmt"
	"strings"
	//"encoding/json"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"

	"meraki/CVDI/types"
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
	MessagesValidation types.JsonArray `json:"messages_validation"`
	ValidationSuccess  bool `json:"is_valid"`
}

func CreateValidator() (*StructValidator) {
	return &StructValidator{}
}

func (this *StructValidator) Validate(_struct interface{}) (err error) {

	_errorValidationStruct := validate.Struct(_struct)
	_messages := make(types.JsonArray, 0)
	this.ValidationSuccess = false

	if _errorValidationStruct != nil {

		if _, ok := _errorValidationStruct.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range _errorValidationStruct.(validator.ValidationErrors) {

			_messages = append(_messages, types.JsonObject{
				strings.ToLower(err.Field()): err.Translate(translate),
			})
		}

	} else {
		this.ValidationSuccess = true
	}

	this.MessagesValidation = _messages

	return _errorValidationStruct
}

func (this StructValidator) GetMessagesValidation() (types.JsonArray) {
	return this.MessagesValidation
}

func (this StructValidator) IsValid() (bool) {
	return this.ValidationSuccess
}
