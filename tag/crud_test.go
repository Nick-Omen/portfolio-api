package tag

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
var createdTagId int
var url = "http://localhost:9002/tag/"

func TestCreateTagWithTitleOnly(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create tag with title only", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"New tag"}`).
			Post(url)
		tag := &Tag{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &tag)
			createdTagId = tag.ID
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 201", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusCreated)
		})
		g.It("Should have title `New tag`", func() {
			g.Assert(tag.Title).Equal("New tag")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(tag.ID == 0).IsFalse()
		})
		g.It("Should have empty link", func() {
			g.Assert(tag.Link).Equal("")
		})
	})
}

func TestGetTag(t *testing.T) {
	g := Goblin(t)

	g.Describe("Get tag", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(fmt.Sprintf("%s%d/", url, createdTagId))
		tag := &Tag{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &tag)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
		g.It("Should have title `New tag`", func() {
			g.Assert(tag.Title).Equal("New tag")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(tag.ID == 0).IsFalse()
		})
		g.It("Should have empty link", func() {
			g.Assert(tag.Link).Equal("")
		})
	})
}

func TestCreateTagWithAllFields(t *testing.T) {
	g := Goblin(t)

	g.Describe("Create tag with all fields", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"Tag","link":"https://tag.com/"}`).
			Post(url)
		tag := &Tag{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &tag)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 201", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusCreated)
		})
		g.It("Should have title `Tag`", func() {
			g.Assert(tag.Title).Equal("Tag")
		})
		g.It("Should have id not equal 0", func() {
			g.Assert(tag.ID == 0).IsFalse()
		})
		g.It("Should have link `https://tag.com/`", func() {
			g.Assert(tag.Link).Equal("https://tag.com/")
		})
	})
}

func TestUpdateTag(t *testing.T) {
	g := Goblin(t)

	g.Describe("Update tag fields", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"title":"New tag","link":"https://new-tag.com/"}`).
			Put(fmt.Sprintf("%s%d/", url, createdTagId))
		tag := &Tag{}
		g.It("Should have no errors when parse body", func() {
			err := json.Unmarshal(res.Body(), &tag)
			g.Assert(err).Equal(nil)
		})
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
		g.It("Should have title `New tag`", func() {
			g.Assert(tag.Title).Equal("New tag")
		})
		g.It("Should have id from url param", func() {
			g.Assert(tag.ID).Equal(createdTagId)
		})
		g.It("Should have link `https://new-tag.com/`", func() {
			g.Assert(tag.Link).Equal("https://new-tag.com/")
		})
	})
}

func TestDeleteTag(t *testing.T) {
	g := Goblin(t)

	g.Describe("Delete tag", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Delete(fmt.Sprintf("%s%d/", url, createdTagId))
		g.It("Should have status code 200", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusOK)
		})
	})

	g.Describe("Get deleted tag", func() {
		res, _ := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(fmt.Sprintf("%s%d/", url, createdTagId))
		g.It("Should have status code 404", func() {
			g.Assert(res.StatusCode()).Equal(http.StatusNotFound)
		})
	})
}

func main() {
	db = storage.Connect("test")
	SetDatabase(db)
}