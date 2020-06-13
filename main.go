package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"	

	user_api "./apis/user_api"
	post_api "./apis/post_api"
	comment_api "./apis/comment_api"
	"./migration"
)

func HandleRequests()	{
	router:=mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/user/register", user_api.Register).Methods("POST")
	router.HandleFunc("/api/user/login", user_api.Login).Methods("POST")
	router.HandleFunc("/api/user", user_api.Users).Methods("GET")
	router.HandleFunc("/api/user",user_api.UserUpdate).Methods("PATCH")
	router.HandleFunc("/api/user/password",user_api.UserUpdatePassword).Methods("PATCH")
	router.HandleFunc("/api/post",post_api.PostCreate).Methods("POST")
	router.HandleFunc("/api/post",post_api.PostDelete).Methods("DELETE")
	router.HandleFunc("/api/comment",comment_api.CommentCreate).Methods("POST")
	router.HandleFunc("/api/comment",comment_api.CommentDelete).Methods("DELETE")	
	

	err:=http.ListenAndServe(":5000", router)
	if err!=nil	{
		fmt.Println(err)
	}
}

func main()	{	
	migration.InitialMigration()
	HandleRequests()	
}
