package project

import (
	"errors"
	"github.com/jinzhu/gorm"
	"nick_omen_api/tag"
)

func SetDatabase(db *gorm.DB) {
	m.DB = db
	m.DB.AutoMigrate(&Project{})
}

type Manager struct {
	DB *gorm.DB
}

type managerInterface interface {
	GetAll() (*[]Project, error)
	GetOne(int) (*Project, error)
	Create(*Project) (*Project, error)
	Update(*Project) (*Project, error)
	Delete(*Project) bool
}

type Project struct {
	ID int `json:"id";gorm:"primary_key"`
	Title string `json:"title";gorm:"type:varchar(100);not null"`
	Description string `json:"description";gorm:"type:varchar(255);default:''"`
	Link string `json:"link";gorm:"type:varchar(100);default:''"`
	Tags []tag.Tag `json:"tags";gorm:"many2many:tags"`
}

var m Manager

func (m Manager) GetOne(id int) (*Project, error) {
	project := &Project{}
	m.DB.Find(&project, id)
	if project.ID == id {
		return project, nil
	}
	return project, errors.New("Project not found. ")
}

func (m Manager) GetAll() (*[]Project, error) {
	projects := &[]Project{}
	m.DB.Find(&projects)
	return projects, nil
}

func (m Manager) Create(p *Project) (*Project, error) {
	created := m.DB.NewRecord(p)
	if !created {
		return p, errors.New("Project is not created. ")
	}
	m.DB.Create(&p)
	return p, nil
}

func (m Manager) Update(p *Project) (*Project, error) {
	m.DB.Save(&p)
	return p, nil
}

func (m Manager) Delete(p *Project) bool {
	m.DB.Delete(&p)
	return true
}

var _ managerInterface = (*Manager)(nil)
