package project

import (
	"testing"
	"github.com/go-resty/resty"
	"encoding/json"
	"net/http"
	. "github.com/franela/goblin"
	"strings"
	"nick_omen_api/tag"
	"fmt"
)

func TestCreateProjectWithTitleOnly(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create project with title only", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"New project"}`).
			Post(testUrl)
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
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
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestCreateProjectWithAllFields(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create project with all fields", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"Portfolio","description":"Description asdfajkljalskdjf lajsdlkfj alsjdfljka sldkfjla","link":"https://www.nick-omen.com/"}`).
			Post(testUrl)
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
		g.It("Should have description", func() {
			g.Assert(project.Description).Equal("Description asdfajkljalskdjf lajsdlkfj alsjdfljka sldkfjla")
		})
		g.It("Should have link `https://www.nick-omen.com/`", func() {
			g.Assert(project.Link).Equal("https://www.nick-omen.com/")
		})
		g.It("Should have empty tags", func() {
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}

func TestCreateProjectWithEmptyBody(t *testing.T) {
	g := Goblin(t)

	g.Describe("Project creation should fail", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{}`).
			Post(testUrl)

		g.It("Should have status code 400", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusBadRequest)
		})
	})
}

func TestCreateProjectWithTags(t *testing.T) {
	g := Goblin(t)

	tg := &tag.Tag{}
	res, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"title":"new tag"}`).
		Post(strings.Replace(testUrl, "project", "tag", 1))
	json.Unmarshal(res.Body(), &tg)

	g.Describe("Create project with tag", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(fmt.Sprintf(`{"title":"New project with tag","tag_ids":[%d]}`, tg.ID)).
			Post(testUrl)
		project := &Project{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have 1 tag", func() {
			g.Assert(len(project.Tags)).Equal(1)
		})

		g.It(fmt.Sprintf("Should have tag with id %d", tg.ID), func() {
			g.Assert(project.Tags[0].ID).Equal(tg.ID)
		})
	})
}
