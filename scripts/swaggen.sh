#!/usr/bin/env bash
bin/swag/swag init -g ./services/user/swag.go -o ./services/user/docs
bin/swag/swag init -g ./services/payment/swag.go -o ./services/payment/docs