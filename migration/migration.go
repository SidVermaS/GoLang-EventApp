package migration

import (
	"fmt"
	// "github.com/jinzhu/gorm"

	"../config"
	"../entities"
)

func InitialMigration()	{
	db,err:=config.GetDB()
	if err!=nil	{
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("InitialMigration")
	db.AutoMigrate(entities.User{})	
	db.AutoMigrate(entities.Post{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT","RESTRICT")
	db.AutoMigrate(entities.Comment{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT","RESTRICT").AddForeignKey("user_id","users(user_id)","RESTRICT","RESTRICT")
	db.AutoMigrate(entities.Reaction{}).AddForeignKey("post_id","posts(post_id)","RESTRICT","RESTRICT").AddForeignKey("user_id","users(user_id)","RESTRICT","RESTRICT")
}