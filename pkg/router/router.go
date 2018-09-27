package router

import (
	"net/http"

	"github.com/containerum/events-api/pkg/eaerrors"

	"time"

	"github.com/containerum/cherry/adaptors/cherrylog"
	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/events-api/pkg/db"
	h "github.com/containerum/events-api/pkg/router/handlers"
	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/containerum/events-api/pkg/server/impl"
	"github.com/containerum/events-api/pkg/util/validation"
	"github.com/containerum/events-api/static"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/utils/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func CreateRouter(mongo *db.MongoStorage, status *model.ServiceStatus, tv *m.TranslateValidate, enableCORS bool) http.Handler {
	e := gin.New()
	systemHandlersSetup(e, status, enableCORS)
	initMiddlewares(e, tv)
	eventsHandlersSetup(e, tv, impl.NewEventsActionsImpl(mongo))

	return e
}

func initMiddlewares(e gin.IRouter, tv *m.TranslateValidate) {
	e.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	binding.Validator = &validation.GinValidatorV9{Validate: tv.Validate} // gin has no local validator
	e.Use(httputil.SaveHeaders)
	e.Use(httputil.PrepareContext)
	e.Use(httputil.RequireHeaders(eaerrors.ErrValidation, httputil.UserIDXHeader, httputil.UserRoleXHeader))
	e.Use(tv.ValidateHeaders(map[string]string{
		httputil.UserIDXHeader:   "uuid",
		httputil.UserRoleXHeader: "eq=admin|eq=user",
	}))
	e.Use(httputil.SubstituteUserMiddleware(tv.Validate, tv.UniversalTranslator, eaerrors.ErrValidation))
	e.Use(m.RequiredUserHeaders())
}

func systemHandlersSetup(router gin.IRouter, status *model.ServiceStatus, enableCORS bool) {
	if enableCORS {
		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		cfg.AddAllowMethods(http.MethodDelete)
		cfg.AddAllowHeaders(httputil.UserRoleXHeader, httputil.UserIDXHeader, httputil.UserNamespacesXHeader)
		router.Use(cors.New(cfg))
	}
	router.Group("/static").
		StaticFS("/", static.HTTP)
	router.Use(gonic.Recovery(eaerrors.ErrInternal, cherrylog.NewLogrusAdapter(logrus.WithField("component", "gin_recovery"))))

	router.GET("/status", httputil.ServiceStatus(status))
}

func eventsHandlersSetup(router gin.IRouter, tv *m.TranslateValidate, backend server.EventsActions) {
	eventsHandlers := h.EventsHandlers{EventsActions: backend, TranslateValidate: tv}

	mainGroup := router.Group("/events")
	{
		mainGroup.GET("/all", eventsHandlers.AllResourcesChangesEventsHandler)           //Websockets only
		mainGroup.GET("/selected", eventsHandlers.SelectedResourcesChangesEventsHandler) //Websockets only

		allEvents := mainGroup.Group("/namespaces/:namespace")
		{
			allEvents.GET("/all", eventsHandlers.AllNamespaceResourcesChangesEventsHandler)           //Websockets only
			allEvents.GET("/selected", eventsHandlers.SelectedNamespaceResourcesChangesEventsHandler) //Websockets only
		}
	}
	containerumEvents := router.Group("/events/containerum")
	{
		containerumEvents.POST("/users", eventsHandlers.AddUserEventHandler)
		containerumEvents.GET("/users", eventsHandlers.GetUsersEventsListHandler)

		containerumEvents.POST("/system", eventsHandlers.AddSystemEventHandler)
		containerumEvents.GET("/system", eventsHandlers.GetSystemEventsListHandler)
	}
	events := router.Group("/events/namespaces/:namespace")
	{
		events.GET("/pods/:pod", eventsHandlers.GetPodEventsListHandler)
		events.GET("/pods", eventsHandlers.GetNamespacePodsEventsListHandler)

		events.GET("/pvc/:pvc", eventsHandlers.GetPVCEventsListHandler)
		events.GET("/pvc", eventsHandlers.GetNamespacePVCsEventsListHandler)
	}
	changes := router.Group("/changes/namespaces/:namespace")
	{
		changes.GET("", eventsHandlers.GetNamespaceChangesListHandler)

		changes.GET("/deployments/:deployment", eventsHandlers.GetDeploymentChangesListHandler)
		changes.GET("/deployments", eventsHandlers.GetNamespaceDeploymentsChangesListHandler)

		changes.GET("/services/:service", eventsHandlers.GetServiceChangesListHandler)
		changes.GET("/services", eventsHandlers.GetNamespaceServicesChangesListHandler)

		changes.GET("/ingresses/:ingress", eventsHandlers.GetIngressChangesListHandler)
		changes.GET("/ingresses", eventsHandlers.GetNamespaceIngressesChangesListHandler)

		changes.GET("/pvc/:pvc", eventsHandlers.GetPVCChangesListHandler)
		changes.GET("/pvc", eventsHandlers.GetNamespacePVCsChangesListHandler)

		changes.GET("/secrets/:secret", eventsHandlers.GetSecretChangesListHandler)
		changes.GET("/secrets", eventsHandlers.GetNamespaceSecretsChangesListHandler)

		changes.GET("/configmaps/:configmap", eventsHandlers.GetConfigMapChangesListHandler)
		changes.GET("/configmaps", eventsHandlers.GetNamespaceConfigMapsChangesListHandler)
	}
}
