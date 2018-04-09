package project

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func SetDatabase(db *gorm.DB) {
	m.DB = db
	m.DB.AutoMigrate(&Project{})
}

type Manager struct {
	DB *gorm.DB
}

type Project struct {
	ID int `json:"id";gorm:"primary_key"`
	Title string `json:"title";gorm:"type:varchar(100);not null"`
	Description string `json:"description";gorm:"type:varchar(255);default:''"`
	Link string `json:"link";gorm:"type:varchar(100);default:''"`
}

var m Manager

func (m Manager) GetProjectById(id int) (*Project, error) {
	project := &Project{}
	m.DB.Find(&project, id)
	if project.ID == id {
		return project, nil
	}
	return project, errors.New("Project not found. ")
}

func (m Manager) GetAllProjects() (*[]Project, error) {
	projects := &[]Project{}
	m.DB.Find(&projects)
	return projects, nil
}

func (m Manager) CreateProject(p *Project) (*Project, error) {
	created := m.DB.NewRecord(p)
	if !created {
		return p, errors.New("Project is not created. ")
	}
	m.DB.Create(&p)
	return p, nil
}

func (m Manager) UpdateProject(p *Project) {
	m.DB.Save(&p)
}

func (m Manager) DeleteProject(p *Project) {
	m.DB.Delete(&p)
}
