package models

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"../entities"
)

type ReactionModel struct	{
	Db *gorm.DB 
}

func (reactionModel ReactionModel) reaction(reqReaction *entities.Reaction) (isFailed bool, err error)	{
	isFailed=true
	reqReaction=&entities.Reaction{Status: reqReaction.Status, Post_id: reqReaction.Post_id, User_id: reqReaction.User_id}

	var reaction entities.Reaction
	var count int
	reactionModel.Db.Where(&entities.Reaction{Post_id: reqReaction.Post_id, User_id: reqReaction.User_id}).Select("status").First(&reqReaction).Scan(&reaction)

	if reaction==entities.Reaction{}	{
		err=reactionModel.Db.Create(reqReaction).Count(&count).Error
	}	else if reaction.Status==reqReaction.Status	{
		isFailed=false
	}	else	{		
		if reaction.Status=="remove"	{
			err=reactionModel.Db.Where(map[string]string{"post_id": reqReaction.Post_id, "user_id": reqReaction.User_id}).Delete(&Reaction{}).Count(&count).Error			
		}	else	{
			err=reactionModel.Db.Model(&reqReaction).Where(map[string]string{"post_id":reqReaction.Post_id, "user_id":reqReaction.User_id}).Update("status",reqReaction.Status).Count(&count).Error
		}
	}
	if err!=nil		{
		isFailed=true
	}
	return
}