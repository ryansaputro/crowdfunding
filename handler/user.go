package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari inputan user
	// mapping input dr user ke struct input.go RegisterUserInput
	// struct diatas akan kita passing sebagai parameter service
	// dari service nanti akan di dependensi kan ke repository
	// dari repository akan di simpan ke database
	var input user.RegisterUserInput
	// proses validasi
	err := c.ShouldBindJSON(&input)
	// validasi error
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("akun gagal dibuat", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	NewUser, err := h.userService.RegisterUser(input)
	// jika user baru gagal dibuat
	if err != nil {
		response := helper.APIResponse("akun gagal dibuat", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(NewUser, "tokenTokenToken")
	response := helper.APIResponse("akun berhasil dibuat", http.StatusOK, "sukses", formatter)

	c.JSON(http.StatusOK, response)

}
