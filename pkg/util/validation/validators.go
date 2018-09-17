package validation

import (
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

func StandardResourceValidator(uni *ut.UniversalTranslator) (ret *validator.Validate) {
	ret = validator.New()
	ret.SetTagName("binding")

	enTranslator, _ := uni.GetTranslator(en.New().Locale())
	enUSTranslator, _ := uni.GetTranslator(en_US.New().Locale())

	enTranslations.RegisterDefaultTranslations(ret, enTranslator)
	enTranslations.RegisterDefaultTranslations(ret, enUSTranslator)

	ret.RegisterStructValidation(eventValidate, kubtypes.Event{})

	return
}

func eventValidate(structLevel validator.StructLevel) {
	req := structLevel.Current().Interface().(kubtypes.Event)

	v := structLevel.Validator()

	if err := v.Var(req.ResourceType, "required"); err != nil {
		structLevel.ReportValidationErrors("ResourceType", "", err.(validator.ValidationErrors))
	}

	if err := v.Var(req.ResourceName, "required"); err != nil {
		structLevel.ReportValidationErrors("ResourceName", "", err.(validator.ValidationErrors))
	}
}
