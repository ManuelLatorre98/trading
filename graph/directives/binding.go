package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	err = validate.Var(val, constraint)

	if err != nil {
		fieldName := *graphql.GetPathContext(ctx).Field
		validationErrors := err.(validator.ValidationErrors)
		transErr := fmt.Errorf("%s%+v", fieldName, validationErrors[0].Translate(trans))
		return val, transErr
	}
	return val, nil
}

func ValidateAddTranslation(tag string, message string) {
	validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		if fe.Param() != "" {
			message = fmt.Sprintf(message, fe.Param())
		}
		return message
	})
}
