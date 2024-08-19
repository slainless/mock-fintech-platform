package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/payment_service"
	"github.com/slainless/mock-fintech-platform/pkg/tracker"
	"github.com/slainless/mock-fintech-platform/services/user"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	db, err := util.NewDB(flagPostgresURL)
	if err != nil {
		return err
	}

	tracker := &tracker.LogTracker{}
	paymentServices := payment_service.InitiatePaymentServices()
	service := user.NewService(flagAuthSecret, db, paymentServices, tracker)

	app := gin.Default()
	service.Mount(app)

	return app.Run(flagAddress.Value()...)
}
