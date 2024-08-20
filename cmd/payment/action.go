package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/payment_service"
	"github.com/slainless/mock-fintech-platform/pkg/tracker"
	"github.com/slainless/mock-fintech-platform/services/payment"

	"github.com/slainless/mock-fintech-platform/cmd/payment/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	docs.SwaggerInfo.Version = version

	db, err := util.NewDB(flagPostgresURL)
	if err != nil {
		return err
	}

	tracker := &tracker.LogTracker{}
	paymentServices := payment_service.InitiatePaymentServices()
	service := payment.NewService(flagAuthSecret, db, paymentServices, tracker)

	app := gin.Default()
	service.Mount(app)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app.Run(flagAddress.Value()...)
}
