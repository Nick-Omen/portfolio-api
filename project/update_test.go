package project

import (
	"testing"
	"github.com/go-resty/resty"
	. "github.com/franela/goblin"
	"encoding/json"
	"fmt"
)

func TestUpdateProject(t *testing.T) {
	project := &Project{}
	g := Goblin(t)

	res, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"title":"New project"}`).
		Post(testUrl)
	json.Unmarshal(res.Body(), &project)
	createdId := project.ID

	g.Describe("Update project fields", func() {
		res, _ = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"Edited","description":"descr","link":"http://test.com/"}`).
			Put(fmt.Sprintf("%s%d/", testUrl, createdId))

		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &project)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have title `Edited`", func() {
			g.Assert(project.Title).Equal("Edited")
		})
		g.It(fmt.Sprintf("Should equal %d", createdId), func() {
			g.Assert(project.ID).Equal(createdId)
		})
		g.It("Should have description `descr`", func() {
			g.Assert(project.Description).Equal("descr")
		})
		g.It("Should have link `http://test.com/`", func() {
			g.Assert(project.Link).Equal("http://test.com/")
		})
		g.It("Should have empty tags", func() {
			g.Assert(len(project.Tags)).Equal(0)
		})
	})
}
