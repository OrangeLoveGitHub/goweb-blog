package model

import (
	"goweb-blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Category struct {

	//gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// query category if exist
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_NAME_USED
	}
	return errmsg.SUCCESS
}

// add category
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// query category list
func GetCategory(pageSize int, pageNum int) []Category {
	var category []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return category
}

// edit category
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&category).Where(" id = ?", id).Update(maps).Error
	//fmt.Print("edit---------------")
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// delete category
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
