package JValidator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
	Uni      *ut.UniversalTranslator
)

func init() {
	Validate = validator.New()

	en := en.New()
	Uni = ut.New(en, en)

	Trans, _ = Uni.GetTranslator("en")
	for key, value := range errorMsg {
		AddTranslation(key, value)
	}
}

func AddTranslation(key, value string) {
	Validate.RegisterTranslation(key, Trans, func(ut ut.Translator) error {
		return ut.Add(key, value, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(key,
			fe.Field(),
			fe.Param(),
			fe.Tag(),
			fmt.Sprintf("%v", fe.Value()),
			fe.Kind().String(),
			fe.Type().String(),
			fe.Namespace(),
			fe.StructNamespace(),
			fe.StructField(),
			fe.ActualTag(),
		)
		return t
	})
}
