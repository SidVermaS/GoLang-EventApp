package api

import (
	"net/http"
	"encoding/json"

	"../../config"
	"../../models"
	"../../entities"
	"../../json_response"
)

func PostCreate(response http.ResponseWriter, request *http.Request)	{
	var post entities.Post
	err:=json.NewDecoder(request.Body).Decode(&post)
	db, err1:=config.GetDB()

	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err1.Error())
	}
	defer db.Close()
	postModel:=models.PostModel{
		Db: db,
	}
	isFailed, err2:=postModel.PostCreate(&post)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully posted",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}
}

func PostDelete(response http.ResponseWriter, request *http.Request)	{
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
	isFailed, err2:=commentModel.CommentDeleteAll(comment.Post_id)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{	
		post:=entities.Post{Post_id: comment.Post_id}
		postModel:=models.PostModel{
			Db: db,
		}
		isFailed1, err3:=postModel.PostDelete(&post)
		if err3!=nil	{
			json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err3.Error())
		}	else if isFailed1	{
			json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
		}	else	{
			mapResponse:=map[string]string{
				"message": "Successfully deleted the post",
			}
			json_response.RespondWithJson(response, http.StatusOK, mapResponse)
		}
	}	
}


