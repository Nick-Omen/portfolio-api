package tag

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type Tag struct {
	ID int `json:"id";gorm:"primary_key"`
	Title string `json:"title";gorm:"type:varchar(100);not null"`
	Link string `json:"link";gorm:"type:varchar(100);default:''"`
}

type Manager struct {
	DB *gorm.DB
}

type managerInterface interface {
	GetAll() (*[]Tag, error)
	GetOne(int) (*Tag, error)
	Create(*Tag) (*Tag, error)
	Update(*Tag) (*Tag, error)
	Delete(*Tag) bool
}

var m Manager

func SetDatabase(db *gorm.DB) {
	m.DB = db
	m.DB.AutoMigrate(&Tag{})
}

func (m Manager) GetOne(id int) (*Tag, error) {
	tag := &Tag{}
	m.DB.Find(&tag, id)

	if tag.ID == id {
		return tag, nil
	}
	return tag, errors.New("Tag not found. ")
}

func (m Manager) GetAll() (*[]Tag, error) {
	tags := &[]Tag{}
	m.DB.Find(&tags)
	return tags, nil
}

func (m Manager) Create(t *Tag) (*Tag, error) {
	created := m.DB.NewRecord(t)
	if !created {
		return t, errors.New("Tag is not created. ")
	}
	m.DB.Create(&t)
	return t, nil
}

func (m Manager) Update(t *Tag) (*Tag, error) {
	m.DB.Save(&t)
	return t, nil
}

func (m Manager) Delete(t *Tag) bool {
	m.DB.Delete(&t)
	return true
}

var _ managerInterface = (*Manager)(nil)
