package validation

import (
	"regexp"

	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	dnsLabel    = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	dockerImage = regexp.MustCompile(`(?:(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])(?:(?:\.(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]))+)?(?::[0-9]+)?/)?[a-z0-9]+(?:(?:(?:[._]|__|[-]*)[a-z0-9]+)+)?(?:(?:/[a-z0-9]+(?:(?:(?:[._]|__|[-]*)[a-z0-9]+)+)?)+)?`)
)

func dnsValidationFunc(fl validator.FieldLevel) bool {
	return dnsLabel.MatchString(fl.Field().String())
}

func dockerImageValidationFunc(fl validator.FieldLevel) bool {
	return dockerImage.MatchString(fl.Field().String())
}

func kubeQuantityValidationFunc(fl validator.FieldLevel) bool {
	_, err := resource.ParseQuantity(fl.Field().String())
	return err == nil
}

func registerCustomTags(v *validator.Validate) {
	v.RegisterValidation("dns", dnsValidationFunc)
	v.RegisterValidation("docker_image", dockerImageValidationFunc)
	v.RegisterValidation("kube_quantity", kubeQuantityValidationFunc)
}

func registerCustomTagsENTranslation(v *validator.Validate, t ut.Translator) {
	v.RegisterTranslation("dns", t, func(ut ut.Translator) error {
		return ut.Add("dns-label", "{0} must be dns-label", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("dns-label", fe.Field())
		if err != nil {
			return err.Error()
		}
		return t
	})

	v.RegisterTranslation("docker_image", t, func(ut ut.Translator) error {
		return ut.Add("docker-image", "{0} must be valid docker image name", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("docker-image", fe.Field())
		if err != nil {
			return err.Error()
		}
		return t
	})

	v.RegisterTranslation("kube_quantity", t, func(ut ut.Translator) error {
		return ut.Add("kube-quantity", "{0} must be valid kubernetes quantity value", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("kube-quantity", fe.Field())
		if err != nil {
			return err.Error()
		}
		return t
	})
}
