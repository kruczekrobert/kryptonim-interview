package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"kryptonim-interview/app/exchanges"
	"net/http"
	"strings"
)

type CoreHandler struct {
	ExchangeService *exchanges.Service
}

func NewCoreHandler(exchangesService *exchanges.Service) *CoreHandler {
	return &CoreHandler{
		ExchangeService: exchangesService,
	}
}

func (h *CoreHandler) RegisterRoutes(r *gin.Engine) {
	r.Group("core/").
		GET("exchange", h.Exchange).
		GET("rates", h.Rates)
}

func (h *CoreHandler) Exchange(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	amountStr := ctx.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	data, err := h.ExchangeService.Exchange(strings.ToUpper(from), strings.ToUpper(to), amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *CoreHandler) Rates(ctx *gin.Context) {
	currencies := strings.Split(ctx.Query("currencies"), ",")
	if len(currencies) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	data, err := h.ExchangeService.Rates(currencies)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(200, data)
}
