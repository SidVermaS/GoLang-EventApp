package user_apis

import(
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	// "github.com/gorilla/mux"

	"../../config"
	"../../models"
	"../../entities"
	"../../json_response"	
)


func Register(response http.ResponseWriter, request *http.Request)	{
	var user entities.User
	err:=json.NewDecoder(request.Body).Decode(&user)
	db, err1:=config.GetDB()
	
	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err1.Error())
	}
	defer db.Close()
	userModel:=models.UserModel{
		Db: db,
	}
	isFailed, err2:=userModel.Register(&user)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully registered",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}	
}

func Login(response http.ResponseWriter, request *http.Request)	{
	var user entities.User
	err:=json.NewDecoder(request.Body).Decode(&user)
	db, err1:=config.GetDB()

	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}
	defer db.Close()
	userModel:=models.UserModel{
		Db: db,
	}
	resultUserLogin:=userModel.Login(&user)
	if resultUserLogin.Status=="unregistered"	{
		json_response.RespondWithError(response, http.StatusUnauthorized, "Unregistered user", "")
	} else if resultUserLogin.Status=="invalid"	{
		json_response.RespondWithError(response, http.StatusUnauthorized, "Invalid credentials", "")
	} else	{
		json_response.RespondWithJson(response, http.StatusOK, resultUserLogin)
	}
}

func Users(response http.ResponseWriter, request *http.Request)	{	
	page, err:=strconv.Atoi(request.URL.Query().Get("page"))

	fmt.Println("page: ",request.FormValue("page"), " ",page)

	db, err1:=config.GetDB()
	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	
	defer db.Close()
	userModel:=models.UserModel{
		Db: db,
	}	
	users, err1:=userModel.Users(page)
	if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else	{
		json_response.RespondWithJson(response, http.StatusOK, users)
	}	
}

func UserUpdate(response http.ResponseWriter, request *http.Request)	{
	var user entities.User
	err:=json.NewDecoder(request.Body).Decode(&user)
	db, err1:=config.GetDB()
	
	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}
	defer db.Close()
	userModel:=models.UserModel{
		Db: db,
	}
	isFailed, err2:=userModel.UserUpdate(&user)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully updated the profile",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}	
}

func UserUpdatePassword(response http.ResponseWriter, request *http.Request)	{
	var userPassword entities.UserPassword
	err:=json.NewDecoder(request.Body).Decode(&userPassword)
	db, err1:=config.GetDB()
	if err!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}	else if err1!=nil	{
		json_response.RespondWithError(response, http.StatusInternalServerError, "Failed", err.Error())
	}
	defer db.Close()
	userModel:=models.UserModel{
		Db: db,
	}
	isFailed, err2:=userModel.UserUpdatePassword(&userPassword)
	if err2!=nil	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", err2.Error())
	}	else if isFailed	{
		json_response.RespondWithError(response, http.StatusBadRequest, "Failed", "")
	}	else	{
		mapResponse:=map[string]string{
			"message": "Successfully updated the password",
		}
		json_response.RespondWithJson(response, http.StatusOK, mapResponse)
	}	

}