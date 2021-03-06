package tag

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type Tag struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"type:varchar(100);not null" valid:"required~title|This field is required,length(2|100)"`
	Link  string `json:"link" gorm:"type:varchar(100);default:''" valid:"optional,url~link|This field only accept links. For example: https://www.example.com/"`
}

type Manager struct {
	DB *gorm.DB
}

type Filter map[string][]int

type managerInterface interface {
	GetAll(Filter) (*[]Tag, error)
	GetOne(int) (*Tag, error)
	Create(*Tag) (*Tag, error)
	Update(*Tag) (*Tag, error)
	Delete(*Tag) bool
}

var M Manager

func SetDatabase(db *gorm.DB) {
	M.DB = db
	M.DB.AutoMigrate(&Tag{})
}

func (m Manager) GetOne(id int) (*Tag, error) {
	tag := &Tag{}
	M.DB.Find(&tag, id)

	if tag.ID == id {
		return tag, nil
	}
	return tag, errors.New("Tag not found. ")
}

func (m Manager) GetAll(filter Filter) (*[]Tag, error) {
	tags := &[]Tag{}

	idList, idListOk := filter["tag_ids"]

	if idListOk {
		if len(idList) > 0 {
			M.DB.Where(idList).Find(&tags)
		}
	}

	if !idListOk {
		M.DB.Find(&tags)
	}

	return tags, nil
}

func (m Manager) Create(t *Tag) (*Tag, error) {
	created := M.DB.NewRecord(t)
	if !created {
		return t, errors.New("Tag is not created. ")
	}
	M.DB.Create(&t)
	return t, nil
}

func (m Manager) Update(t *Tag) (*Tag, error) {
	M.DB.Save(&t)
	return t, nil
}

func (m Manager) Delete(t *Tag) bool {
	M.DB.Delete(&t)
	return true
}

var _ managerInterface = (*Manager)(nil)
