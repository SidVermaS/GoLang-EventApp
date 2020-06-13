package models

type Reaction struct	{
	Reaction_id `json:"reaction_id"gorm:"auto_increment;not_null;primary_key"`
	Status `json:"status"gorm:"varchar(10)"`
	Post_id `json:"post_id"gorm:"auto_increment"`
	User_id `json:"user_id"gorm:auto_increment"`
}	







