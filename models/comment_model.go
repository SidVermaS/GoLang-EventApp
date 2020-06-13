package models

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"../entities"
)

type CommentModel struct	{
	Db *gorm.DB 
}

func (commentModel CommentModel) CommentCreate(reqComment *entities.Comment) (isFailed bool, err error)	{
	reqComment=&entities.Comment{Comment_text: reqComment.Comment_text, Post_id: reqComment.Post_id, User_id: reqComment.User_id}

	err=commentModel.Db.Create(reqComment).Error
	isFailed=commentModel.Db.NewRecord(reqComment)
	return
}

func (commentModel CommentModel) CommentDelete(reqComment *entities.Comment) (isFailed bool, err error)	{
	reqComment=&entities.Comment{Comment_id: reqComment.Comment_id}
	var count int

	err=commentModel.Db.Delete(&reqComment).Count(&count).Error
	if err!=nil || count==0	{
		isFailed=true
	}	else	{
		isFailed=false
	}
	return
}

func (commentModel CommentModel) CommentDeleteAll(Post_id int) (isFailed bool, err error)	{
	fmt.Println("pid",Post_id)
	err=commentModel.Db.Where("post_id=?",Post_id).Delete(&entities.Comment{}).Error
	
	if err!=nil	{
		isFailed=true
	}	else	{
		isFailed=false
	}
	return
}














































