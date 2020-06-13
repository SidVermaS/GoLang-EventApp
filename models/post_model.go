package models

import (
	"github.com/jinzhu/gorm"

	"../entities"
)

type PostModel struct	{
	Db *gorm.DB
}
 
func (postModel PostModel) PostCreate(reqPost *entities.Post) (isFailed bool, err error)	{
	reqPost=&entities.Post{Caption: reqPost.Caption, Photo_url: reqPost.Photo_url, User_id: reqPost.User_id}
	err=postModel.Db.Create(reqPost).Error
	isFailed=postModel.Db.NewRecord(reqPost)
	return
}

func (postModel PostModel) PostDelete(reqPost *entities.Post) (isFailed bool, err error)	{
	reqPost=&entities.Post{Post_id: reqPost.Post_id}
	var count int
	err=postModel.Db.Delete(&reqPost).Count(&count).Error
	
	if err!=nil || count==0	{
		isFailed=true
	}	else	{
		isFailed=false
	}
	return
}
















