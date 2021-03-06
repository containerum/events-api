package main

import (
	"net/http"
	"os"
	"time"

	"os/signal"

	"context"

	"fmt"
	"text/tabwriter"

	"github.com/containerum/events-api/pkg/router"
	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/util/validation"
	"github.com/containerum/kube-client/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func initServer(c *cli.Context) error {
	setupLogs(c)

	log.Infof("Starting %v %v", c.App.Name, c.App.Version)

	w := tabwriter.NewWriter(log.StandardLogger().Writer(), 0, 0, 2, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for _, f := range c.GlobalFlagNames() {
		fmt.Fprintf(w, "Flag: %s\t Value: %s\n", f, c.String(f))
	}
	if err := w.Flush(); err != nil {
		log.Debug(err)
	}

	translate := setupTranslator()
	validate := validation.StandardResourceValidator(translate)

	tv := &m.TranslateValidate{UniversalTranslator: translate, Validate: validate}

	mongo, err := setupMongo(c)
	exitOnError(err)
	defer mongo.Close()

	status := model.ServiceStatus{
		Name:     c.App.Name,
		Version:  c.App.Version,
		StatusOK: true,
	}

	app := router.CreateRouter(mongo, &status, tv, c.Duration("db_request_period"), c.Bool("cors"))

	srv := &http.Server{
		Addr:    ":" + c.String("port"),
		Handler: app,
	}

	// serve connections
	go exitOnError(srv.ListenAndServe())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // subscribe on interrupt event
	<-quit                            // wait for event
	log.Infoln("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func exitOnError(err error) {
	if err != nil {
		log.WithError(err).Fatalf("can`t setup events-api")
		os.Exit(1)
	}
}
