package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	// validasi error
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokenTokenToken")
	response := helper.APIResponse("login sukses", http.StatusOK, "sukses", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	// validasi error
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("email sudah digunakan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "server error"}

		response := helper.APIResponse("email sudah digunakan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_avaibility": isEmailAvailable,
	}

	metaMessage := "email terdaftar"

	if isEmailAvailable {
		metaMessage = "email tersedia"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "sukses", data)
	c.JSON(http.StatusOK, response)
	return

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("avatar gagal disimpan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("avatar gagal disimpan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("avatar gagal disimpan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("avatar berhasil disimpan", http.StatusOK, "sukses", data)
	c.JSON(http.StatusOK, response)
	return

}
