package tests

import (
    "github.com/robfig/revel"
    "dashpersonal/app/models"
    "encoding/json"
    "log"
    "bytes"
)

type AppTest struct {
	revel.TestSuite
}

func (t AppTest) Before() {
	println("Set up")
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t AppTest) After() {
	println("Tear down")
}

func (t AppTest) TestAddWidget() {
    widget := models.Widget{"X", "1", "1", "1", "1", "Hello world"}
    var widget_created models.Widget

    content, err := json.Marshal(widget);

    if err != nil {
        log.Fatal("Error")
    }

    t.Post("/", "application/json", bytes.NewBuffer(content))
    t.AssertOk()
    t.AssertContentType("application/json")
    err_created := json.Unmarshal(t.ResponseBody, &widget_created)

    if err_created != nil {
        log.Fatal("Error Response")
    }

    t.Assert(widget_created.Content == widget.Content)
}
