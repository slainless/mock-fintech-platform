package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/payment_service"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
	"github.com/slainless/mock-fintech-platform/pkg/tracker"
	"github.com/slainless/mock-fintech-platform/services/user"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	db, err := util.NewDB(flagPostgresURL)
	if err != nil {
		return err
	}

	tracker := &tracker.NilTracker{}
	mockService := payment_service.NewMockPaymentService()
	service := user.NewService(flagAuthSecret, db,
		map[string]platform.PaymentService{
			"bank_of_the_xyz": mockService,
			"infinite_loan":   mockService,
			"fishtech":        mockService,
		},
		tracker,
	)

	app := gin.Default()
	service.Mount(app)

	return app.Run(flagAddress.Value()...)
}
