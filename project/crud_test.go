package project

import (
	"testing"
	"github.com/jinzhu/gorm"
	"nick_omen_api/storage"
	. "github.com/franela/goblin"
	"github.com/go-resty/resty"
	"net/http"
	"encoding/json"
	"fmt"
)

var db *gorm.DB
var createdProjectId int
var url = "http://localhost:9002/project/"

func TestCreateProjectWithTitleOnly(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create project with title only", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"New project"}`).
			Post(url)
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			createdProjectId = project.ID
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 201", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusCreated)
		})
		g.It("Should have title `New project`", func() {
			g.Assert(project.Title).Equal("New project")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(project.ID == 0).IsFalse()
		})
		g.It("Should have empty description", func() {
			g.Assert(project.Description).Equal("")
		})
		g.It("Should have empty link", func() {
			g.Assert(project.Link).Equal("")
		})
		g.It("Should have empty tags", func() {
			t.Log(project.Tags)
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestGetProject(t *testing.T) {
	g := Goblin(t)

	g.Describe("Get project", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(fmt.Sprintf("%s%d/", url, createdProjectId))
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
		g.It("Should have title `New project`", func() {
			g.Assert(project.Title).Equal("New project")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(project.ID == 0).IsFalse()
		})
		g.It("Should have empty description", func() {
			g.Assert(project.Description).Equal("")
		})
		g.It("Should have empty link", func() {
			g.Assert(project.Link).Equal("")
		})
		g.It("Should have empty tags", func() {
			t.Log(project.Tags)
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestCreateProjectWithAllFields(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create project with all fields", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"Portfolio","description":"Description","link":"https://www.nick-omen.com/"}`).
			Post(url)
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 201", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusCreated)
		})
		g.It("Should have title `Portfolio`", func() {
			g.Assert(project.Title).Equal("Portfolio")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(project.ID == 0).IsFalse()
		})
		g.It("Should have description `Description`", func() {
			g.Assert(project.Description).Equal("Description")
		})
		g.It("Should have link `https://www.nick-omen.com/`", func() {
			g.Assert(project.Link).Equal("https://www.nick-omen.com/")
		})
		g.It("Should have empty tags", func() {
			t.Log(project.Tags)
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestUpdateProject(t *testing.T) {
	g := Goblin(t)

	g.Describe("Update project fields", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"New title","description":"","link":"https://www.new-link.com/"}`).
			Put(fmt.Sprintf("%s%d/", url, createdProjectId))
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
		g.It("Should have title `New title`", func() {
			g.Assert(project.Title).Equal("New title")
		})
		g.It("Should have id from url param", func() {
			g.Assert(project.ID).Equal(createdProjectId)
		})
		g.It("Should have empty description", func() {
			g.Assert(project.Description).Equal("")
		})
		g.It("Should have link `https://www.new-link.com/`", func() {
			g.Assert(project.Link).Equal("https://www.new-link.com/")
		})
		g.It("Should have empty tags", func() {
			t.Log(project.Tags)
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestDeleteProject(t *testing.T) {
	g := Goblin(t)

	g.Describe("Delete project", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Delete(fmt.Sprintf("%s%d/", url, createdProjectId))
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
	})

	g.Describe("Get deleted project", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(fmt.Sprintf("%s%d/", url, createdProjectId))
		g.It("Should have status code 404", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusNotFound)
		})
	})
}

func main() {
	db = storage.Connect("test")
	SetDatabase(db)
}