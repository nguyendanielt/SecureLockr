package handler

import (
	"net/http"

	"passwordservice/internal/api/helper"
	"passwordservice/pkg/database"

	"github.com/google/uuid"
)

func GetPassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	password := database.GetPwd(reqBody.PasswordId.String(), userId)
	if password == nil {
		http.Error(w, "Could not find password", http.StatusNotFound)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"password": password,
	})
	helper.SendActivityProducerMessage(userId, "GET ONE PASSWORD")
}

func GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	passwords := database.GetAllPwds(userId)
	if passwords == nil {
		http.Error(w, "Could not find all passwords associated with user", http.StatusNotFound)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"passwords": passwords,
	})
	helper.SendActivityProducerMessage(userId, "GET ALL PASSWORDS")
}

func AddPassword(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(r.Header.Get("userid"))
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	reqBody.UserId = userId
	if err := database.AddPwd(reqBody); err != nil {
		http.Error(w, "Failed to add password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully added password",
	})
	helper.SendActivityProducerMessage(userId.String(), "ADD PASSWORD")
}

func DeletePassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	if err := database.DeletePwd(reqBody.PasswordId.String(), userId); err != nil {
		http.Error(w, "Failed to delete password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully deleted password",
	})
	helper.SendActivityProducerMessage(userId, "DELETE PASSWORD")
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(r.Header.Get("userid"))
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	reqBody.UserId = userId
	if err := database.UpdatePwd(reqBody); err != nil {
		http.Error(w, "Failed to update password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully updated password",
	})
	helper.SendActivityProducerMessage(userId.String(), "UPDATE PASSWORD")
}
