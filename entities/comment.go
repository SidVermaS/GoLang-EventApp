package entities

type Comment struct	{
	Comment_id int `json:"comment_id"gorm:"auto_increment;not null;primary_key"`
	Comment_text string `json:"comment_text"gorm:"type:varchar(100)"`
	Post_id int `json:"post_id"gorm:"not null"`
	User_id int `json:"user_id"gorm:"not null"`
}

















