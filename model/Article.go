package model

import (
	"goweb-blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category Category
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid" `
	Desc     string `gorm:"type:varchar(200)" json:"desc"`
	Content  string `gorm:"type:longtext" json:"content"`
	Img      string `gorm:"type:varchar(100)" json:"img"`
}

// add Article
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//TODO 查询单个文章
//TODO 查询文章列表
// query Article list
func GetArticle(pageSize int, pageNum int) []Article {
	var article []Article
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return article
}

// edit Article
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where(" id = ?", id).Update(maps).Error
	//fmt.Print("edit---------------")
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// delete Article
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
