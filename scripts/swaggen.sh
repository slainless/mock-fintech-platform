#!/usr/bin/env bash
bin/swag/swag init -g ./services/user/swag.go -o ./cmd/user/docs
bin/swag/swag init -g ./services/payment/swag.go -o ./cmd/payment/docs