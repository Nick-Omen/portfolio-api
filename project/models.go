package project

import (
	"errors"
	"github.com/jinzhu/gorm"
	"nick_omen_api/tag"
)

func SetDatabase(db *gorm.DB) {
	M.DB = db
	M.DB.AutoMigrate(&Project{})
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
	ID int `gorm:"primary_key" json:"id"`
	Title string `gorm:"type:varchar(100);not null" json:"title" valid:"required~title|This field is required,length(2|100)"`
	Description string `gorm:"type:varchar(255);default:''" json:"description" valid:"optional,length(20|255)~description|This field length should be from 20 to 255"`
	Link string `gorm:"type:varchar(100);default:''" json:"link" valid:"optional,url~link|This field only accept links. For example: https://www.example.com/"`
	Tags []tag.Tag `gorm:"many2many:project_tags" json:"tags"`
	TagIDs []int `gorm:"-" json:"tag_ids"`
}

var M Manager

func (m Manager) GetOne(id int) (*Project, error) {
	project := &Project{}
	tags := &[]tag.Tag{}
	M.DB.Find(project, id).Related(tags, "Tags")
	if project.ID == id {
		project.Tags = *tags
		return project, nil
	}
	return project, errors.New("Project not found. ")
}

func (m Manager) GetAll() (*[]Project, error) {
	projects := &[]Project{}
	tags := &[]tag.Tag{}
	M.DB.Find(projects).Related(tags)
	return projects, nil
}

func (m Manager) Create(p *Project) (*Project, error) {
	created := M.DB.NewRecord(p)
	if !created {
		return p, errors.New("Project is not created. ")
	}
	M.DB.Create(p)

	tags, _ := tag.M.GetAll(tag.Filter{"tag_ids": p.TagIDs})
	for _, t := range *tags {
		M.DB.Model(p).Association("Tags").Append(t)
	}

	p, err := M.GetOne(p.ID)
	if err == nil {
		return p, nil
	}
	return p, err
}

func (m Manager) Update(p *Project) (*Project, error) {
	M.DB.Save(p)

	newTags, _ := tag.M.GetAll(tag.Filter{"tag_ids": p.TagIDs})
	M.DB.Model(p).Association("Tags").Clear()
	for _, t := range *newTags {
		M.DB.Model(p).Association("Tags").Append(t)
	}

	p, err := M.GetOne(p.ID)
	if err == nil {
		return p, nil
	}
	return p, err
}

func (m Manager) GetTags(p *Project) (*[]tag.Tag, error) {
	tags := &[]tag.Tag{}
	M.DB.Model(p).Related(tags, "Tags")
	return tags, nil
}

func (m Manager) Delete(p *Project) bool {
	M.DB.Model(p).Association("Tags").Clear()
	M.DB.Delete(&p)
	return true
}

var _ managerInterface = (*Manager)(nil)
