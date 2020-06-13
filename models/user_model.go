package models

import	(
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	pagination "github.com/biezhi/gorm-paginator/pagination"
	"../entities"
)

type UserModel struct	{
	Db *gorm.DB	
}

func (userModel UserModel) Register(reqUser *entities.User) (isFailed bool, err error)	{
	data:=[]byte(reqUser.Password)
	h:=md5.Sum(data)
	s:=hex.EncodeToString(h[:])
	reqUser=&entities.User{Name: reqUser.Name, Role: reqUser.Role, Mobile_no: reqUser.Mobile_no, Password: s}
	err=userModel.Db.Create(reqUser).Error
	isFailed=userModel.Db.NewRecord(reqUser)
	return
}

func (userModel UserModel) Login(reqUser *entities.User) (resultUserLogin *entities.ResultUserLogin)	{
	data:=[]byte(reqUser.Password)
	h:=md5.Sum(data)
	s:=hex.EncodeToString(h[:])
	reqUser=&entities.User{Mobile_no: reqUser.Mobile_no, Password: s}	
	var user entities.User
	var userLogin entities.UserLogin
	userModel.Db.Where(&entities.User{Mobile_no: reqUser.Mobile_no}).Select("user_id").First(&reqUser).Scan(&user)
	var status string
	if user==entities.User{}	{
		status="unregistered"
	}	else	{
		
		userModel.Db.Where(&entities.User{Mobile_no: reqUser.Mobile_no, Password: s}).Select("user_id, mobile_no, name, role, photo_url").First(&reqUser).Scan(&userLogin)
		if(userLogin==entities.UserLogin{})	{
			status="invalid"
		}	else	{
			status="valid"
		}
	}
	resultUserLogin=&entities.ResultUserLogin{Status: status, UserLogin: userLogin}
	return
}

func (userModel UserModel) Users(page int) (users []entities.Users, err error)	{

	pagination.Paging(&pagination.Param{
		DB: userModel.Db,
		Page: page,
		Limit: 10,
		OrderBy: []string{"name asc"},
	},&users)

	return
}

func (userModel UserModel) UserUpdate(reqUser *entities.User) (isFailed bool, err error)	{
	reqUser=&entities.User{User_id: reqUser.User_id, Name: reqUser.Name, Role: reqUser.Role}
	var userMap=make(map[string]interface{})
	if(reqUser.Name!="")	{
		userMap["name"]=reqUser.Name
	}
	if(reqUser.Role!="")	{
		userMap["role"]=reqUser.Role
	}
	var count int
	err=userModel.Db.Model(&reqUser).Updates(userMap).Count(&count).Error

	fmt.Println("count: ",count)
	if err!=nil || count==0	{
		isFailed=true
	}	else	{
		isFailed=false
	}
	return
}

func (userModel UserModel) UserUpdatePassword(reqUserPassword *entities.UserPassword) (isFailed bool, err error)	{
	data:=[]byte(reqUserPassword.Old_password)
	h:=md5.Sum(data)
	reqUserPassword.Old_password=hex.EncodeToString(h[:])
	
	data=[]byte(reqUserPassword.New_password)
	h=md5.Sum(data)
	nps:=hex.EncodeToString(h[:])

	user:=&entities.User{User_id: reqUserPassword.User_id}

	var count int
	err=userModel.Db.Model(&user).Where("password=?",reqUserPassword.Old_password).Update("password", nps).Count(&count).Error
	
	fmt.Println("count", count, " ",nps," old: ",reqUserPassword.Old_password)

	if err!=nil || count==0	{
		isFailed=true
	}	else	{
		isFailed=false
	}
	return
}