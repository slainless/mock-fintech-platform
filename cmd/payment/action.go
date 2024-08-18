package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
	"github.com/slainless/mock-fintech-platform/pkg/tracker"
	"github.com/slainless/mock-fintech-platform/services/payment"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	db, err := util.NewDB(flagPostgresURL)
	if err != nil {
		return err
	}

	tracker := &tracker.NilTracker{}
	service := payment.NewService(flagAuthSecret, db, map[string]platform.PaymentService{}, tracker)

	app := gin.Default()
	service.Mount(app)

	return app.Run()
}
