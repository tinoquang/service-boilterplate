package router

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type customBinder struct {
	binder        echo.Binder
	validator     *validator.Validate
	errTranslator ut.Translator
}

func newCustomBinder() *customBinder {
	v := validator.New()

	en := en.New()
	uni := ut.New(en, en)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic(err)
	}

	return &customBinder{
		binder:        &echo.DefaultBinder{},
		validator:     v,
		errTranslator: trans,
	}
}

func (cb *customBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.binder.Bind(i, c); err != nil {
		return err
	}

	// if err.Error() == "EOF" {
	// 	return "missing request body"
	// }

	if err := cb.validator.Struct(i); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		errTrans := validationErr.Translate(cb.errTranslator)
		var errMsg string
		for _, e := range errTrans {
			if errMsg != "" {
				errMsg += ", "
			}
			errMsg += e
		}

		return fmt.Errorf("invalid request body: %s", errMsg)
	}

	return nil
}
