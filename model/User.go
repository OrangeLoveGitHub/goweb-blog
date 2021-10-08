package model

import (
	"encoding/base64"
	"goweb-blog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" `
	Password string `gorm:"type:varchar(20);not null" json:"password" `
	Role     int    `gorm:"type:int" json:"role"`
}

// query user if exist
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// add user
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// query user list
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// edit user

// encrypt password
func ScryptPw(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 5, 6, 21, 22, 9}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw

}
