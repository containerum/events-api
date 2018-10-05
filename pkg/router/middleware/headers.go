package middleware

import (
	"github.com/containerum/kube-client/pkg/model"

	"github.com/containerum/events-api/pkg/eaerrors"

	"net/textproto"

	"encoding/base64"

	"errors"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type UserHeaderDataMap map[string]model.UserHeaderData

const (
	UserNamespaces = "user-namespaces"

	RoleUser  = "user"
	RoleAdmin = "admin"
)

func RequiredUserHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.WithField("Headers", ctx.Request.Header).Debug("Header list")
		notFoundHeaders := requireHeaders(ctx, httputil.UserRoleXHeader)
		if len(notFoundHeaders) > 0 {
			gonic.Gonic(eaerrors.ErrRequiredHeadersNotProvided().AddDetails(notFoundHeaders...), ctx)
			return
		}
		// Check User-Role and User-Namespace
		if isUser, err := checkIsUserRole(GetHeader(ctx, httputil.UserRoleXHeader)); err != nil {
			log.WithField("Value", GetHeader(ctx, httputil.UserRoleXHeader)).WithError(err).Warn("check User-Role Error")
			gonic.Gonic(eaerrors.ErrInvalidRole(), ctx)
		} else {
			//User-Role: user, check User-Namespace
			if isUser {
				notFoundHeaders := requireHeaders(ctx, httputil.UserNamespacesXHeader, httputil.UserIDXHeader)
				if len(notFoundHeaders) > 0 {
					gonic.Gonic(eaerrors.ErrRequiredHeadersNotProvided().AddDetails(notFoundHeaders...), ctx)
					return
				}
				userNs, errNs := checkUserNamespace(GetHeader(ctx, httputil.UserNamespacesXHeader))
				if errNs != nil {
					log.WithField("Value", GetHeader(ctx, httputil.UserNamespacesXHeader)).WithError(errNs).Warn("Check User-Namespace header Error")
					gonic.Gonic(eaerrors.ErrValidation().AddDetailF("%v: %v", httputil.UserNamespacesXHeader, errNs), ctx)
					return
				}
				ctx.Set(UserNamespaces, userNs)
			}
		}
	}
}

func checkIsUserRole(userRole string) (bool, error) {
	switch userRole {
	case "", RoleAdmin:
		return false, nil
	case RoleUser:
		return true, nil
	}
	return false, errors.New("invalid user role")
}

func requireHeaders(ctx *gin.Context, headers ...string) (notFoundHeaders []string) {
	for _, v := range headers {
		if GetHeader(ctx, v) == "" {
			notFoundHeaders = append(notFoundHeaders, v)
		}
	}
	return
}

func GetHeader(ctx *gin.Context, header string) string {
	return ctx.GetHeader(textproto.CanonicalMIMEHeaderKey(header))
}

func checkUserNamespace(userNamespace string) (*UserHeaderDataMap, error) {
	return ParseUserHeaderData(userNamespace)
}

//ParseUserHeaderData decodes headers for substitutions
func ParseUserHeaderData(str string) (*UserHeaderDataMap, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.WithError(err).WithField("Value", str).Warn("unable to decode user header data")
		return nil, errors.New("unable to decode user header data")
	}
	var userData []model.UserHeaderData
	err = jsoniter.Unmarshal(data, &userData)
	if err != nil {
		log.WithError(err).WithField("Value", string(data)).Warn("unable to unmarshal user header data")
		return nil, errors.New("unable to unmarshal user header data")
	}
	result := UserHeaderDataMap{}
	for _, v := range userData {
		result[v.ID] = v
	}
	return &result, nil
}
