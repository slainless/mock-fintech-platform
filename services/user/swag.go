package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var stub *platform.TransactionHistory
var stub2 *core.HistoryParams

// @title Account Manager Service
// @contact.name Aiman Fauzy
// @BasePath /

// @accept json
// @accept mpfd
// @accept x-www-form-urlencoded

// ====
// @summary Register user with JWT token
// @router /register [post]
// @accept json
// @accept mpfd
// @param payload body RegisterPayload true "JWT token"
// @produce json
// @failure default {string} string
// @success 201 {object} RegisterResponse
func stubRegister(c *gin.Context) {}

// ===
// @summary Get user's payment accounts
// @router /account [get]
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} AccountsResponse
func stubAccounts(c *gin.Context) {}

// ===
// @summary Get user's payment account by UUID
// @router /account/{account_uuid} [get]
// @param account_uuid path string true "Account UUID"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} AccountResponse
func stubAccount(c *gin.Context) {}

// ===
// @summary Register user's payment account
// @router /account [post]
// @param payload body CreatePayload true "Account data"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 201 {object} AccountResponse
func stubCreateAccount(c *gin.Context) {}

// ===
// @summary Get user's payment account histories
// @router /history [get]
// @param params query core.HistoryParams true "History params"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} HistoriesResponse
func stubHistories(c *gin.Context) {}

// ===
// @summary Get user's recurring payments
// @router /subscription [get]
// @param params query SubscriptionParams true "Subscription params"
// @param Authorization header string true "Authentication token"
// @produce json
// @failure default {string} string
// @success 200 {object} SubscriptionResponse
func stubSubscriptions(c *gin.Context) {}
