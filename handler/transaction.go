package handler

import (
	"crowdfunding/helper"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetTransactionsInput
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("error ambil data transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionsByCampaignID(input)

	if err != nil {
		response := helper.APIResponse("error ambil data transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("list data transactions", http.StatusOK, "sukses", transaction.FormatterCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
	return

}
