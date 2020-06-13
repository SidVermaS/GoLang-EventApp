package entities

type Post struct	{
	Post_id int `json:"post_id"gorm:"auto_increment;not null;primary_key"`
	Caption string `json:"caption"gorm:"type:varchar(80)"`
	Photo_url string `json:"photo_url"gorm:"type:varchar(60)"`
	User_id int `json:"user_id"gorm:"not null"`
}










