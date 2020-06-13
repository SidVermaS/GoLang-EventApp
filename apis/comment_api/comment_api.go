package api

import (
	"net/http"
	"encoding/json"
	
	"../../config"
	"../../models"
	"../../entities"
	"../../json_response"
)

func CommentCreate(response http.ResponseWriter, request *http.Request)	{
	var comment entities.Comment
	err:=json.NewDecoder(request.Body).Decode(&comment)
	db, err1:=config.GetDB()

	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err1.Error())
	}
	defer db.Close()
	commentModel:=models.CommentModel{
		Db: db,
	}
	isFailed, err2:=commentModel.CommentCreate(&comment)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully commented",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}

}

func CommentDelete(response http.ResponseWriter, request *http.Request)	{
	var comment entities.Comment
	err:=json.NewDecoder(request.Body).Decode(&comment)
	db, err1:=config.GetDB()

	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err1.Error())
	}
	defer db.Close()
	commentModel:=models.CommentModel{
		Db: db,
	}
	isFailed, err2:=commentModel.CommentDelete(&comment)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully deleted the comment",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}
}






















































