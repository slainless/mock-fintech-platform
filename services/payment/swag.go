package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var stub *platform.TransactionHistory
var stub2 *core.HistoryParams

// @title Payment Manager Service
// @contact.name Aiman Fauzy
// @BasePath /

// @accept json
// @accept mpfd
// @accept x-www-form-urlencoded

// ===
// @summary Send amount to account UUID
// @router /send [post]
// @param payload body SendPayload true "Send payload"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} SendResponse
func stubSend(c *gin.Context) {}

// ===
// @summary Withdraw amount from account
// @router /withdraw [post]
// @param payload body WithdrawPayload true "Withdraw payload"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} WithdrawResponse
func stubWithdraw(c *gin.Context) {}
