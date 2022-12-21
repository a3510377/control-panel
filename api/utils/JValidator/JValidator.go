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
			fe.Field(),                    // {0} Field
			fe.Param(),                    // {1} Param
			fe.Tag(),                      // {2} Tag
			fmt.Sprintf("%v", fe.Value()), // {3} Value
			fe.Kind().String(),            // {4} Kind
			fe.Type().String(),            // {5} Type
			fe.Namespace(),                // {6} Namespace
			fe.StructNamespace(),          // {7} StructNamespace
			fe.StructField(),              // {8} StructField
			fe.ActualTag(),                // {9} ActualTag
		)
		return t
	})
}
