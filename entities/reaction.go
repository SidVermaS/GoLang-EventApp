package entities

type Reaction struct	{
	Reaction_id int `json:"reaction_id"gorm:"auto_increment;not_null;primary_key"`
	Status string`json:"status"gorm:"varchar(10)"`
	Post_id int `json:"post_id"gorm:"auto_increment"`
	User_id int `json:"user_id"gorm:auto_increment"`
}	







