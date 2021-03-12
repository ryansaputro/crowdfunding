package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(UserID)

	if err != nil {
		response := helper.APIResponse("error ambil data campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("data campaign", http.StatusOK, "sukses", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("gagal mengambil detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignByID(input)

	if err != nil {
		response := helper.APIResponse("gagal mengambil detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("detail campaign", http.StatusOK, "sukses", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	// validasi error
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("campaign gagal dibuat", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)

	if err != nil {
		response := helper.APIResponse("gagal membuat campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("campaign sukses dibuat", http.StatusOK, "sukses", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}
