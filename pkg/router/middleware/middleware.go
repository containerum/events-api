package middleware

import (
	"net/textproto"

	"github.com/containerum/events-api/pkg/eaerrors"

	"github.com/containerum/cherry"
	"github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type TranslateValidate struct {
	*ut.UniversalTranslator
	*validator.Validate
}

func (tv *TranslateValidate) HandleError(err error) (int, *cherry.Err) {
	switch err.(type) {
	case *cherry.Err:
		e := err.(*cherry.Err)
		return e.StatusHTTP, e
	default:
		return eaerrors.ErrInternal().StatusHTTP, eaerrors.ErrInternal().AddDetailsErr(err)
	}
}

func (tv *TranslateValidate) BadRequest(ctx *gin.Context, err error) (int, *cherry.Err) {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		ret := eaerrors.ErrValidation()
		for _, fieldErr := range validationErr {
			if fieldErr == nil {
				continue
			}
			t, _ := tv.FindTranslator(httputil.GetAcceptedLanguages(ctx.Request.Context())...)
			ret.AddDetailF("Field %s: %s", fieldErr.Namespace(), fieldErr.Translate(t))
		}
		return ret.StatusHTTP, ret
	}
	return eaerrors.ErrValidation().StatusHTTP, eaerrors.ErrValidation().AddDetailsErr(err)
}

func (tv *TranslateValidate) ValidateHeaders(headerTagMap map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerErr := make(map[string]validator.ValidationErrors)
		for header, tag := range headerTagMap {
			ferr := tv.VarCtx(ctx.Request.Context(), ctx.GetHeader(textproto.CanonicalMIMEHeaderKey(header)), tag)
			if ferr != nil {
				headerErr[header] = ferr.(validator.ValidationErrors)
			}
		}
		if len(headerErr) > 0 {
			ret := eaerrors.ErrValidation()
			for header, fieldErrs := range headerErr {
				for _, fieldErr := range fieldErrs {
					if fieldErr == nil {
						continue
					}
					t, _ := tv.FindTranslator(httputil.GetAcceptedLanguages(ctx.Request.Context())...)
					ret.AddDetailF("Header %s: %s", header, fieldErr.Translate(t))
				}
			}
			ctx.AbortWithStatusJSON(ret.StatusHTTP, ret)
			return
		}
	}
}
