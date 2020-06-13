package entities

type User struct	{
	User_id int `json:"user_id"gorm:"auto_increment;not null;primary_key"`
	Mobile_no string `json:"mobile_no"gorm:"type:bigint(12);unique"`
	Name string `json:"name"gorm:"type:varchar(32)"`
	Role string `json:"role"gorm:"type:varchar(32)"`
	Photo_url string `json:"photo_url"gorm:"type:varchar(60)"`
	Password string `json:"password"gorm:"type:text"`
	Device_token string `json:"device_token"gorm:"type:text"`
}
type Users struct	{
	User_id int `json:"user_id"`
	Name string `json:"name"`
	Photo_url string `json:"Photo_url"`
}
type UserLogin struct	{
	User_id int `json:"user_id"gorm:"auto_increment;not null;primary_key"`
	Mobile_no string `json:"mobile_no"gorm:"type:bigint(12);unique"`
	Name string `json:"name"gorm:"type:varchar(32)"`
	Role string `json:"role"gorm:"type:varchar(32)"`
	Photo_url string `json:"photo_url"gorm:"type:varchar(60)"`
}
type ResultUserLogin struct	{
	Message string `json:"message"`
	Status string `json:"status"`
	UserLogin UserLogin `json:"user"`		
}
type UserPassword struct	{
	User_id int `json:"user_id"`
	Old_password string `json:"old_password"`
	New_password string `json:"new_password"`
}