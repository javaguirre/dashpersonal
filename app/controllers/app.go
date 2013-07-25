package controllers

import (
    "encoding/json"
    "github.com/robfig/revel"
    "github.com/taikoa/dashpersonal/app/models"
    "io/ioutil"
    "log"
)


type App struct { *revel.Controller
}

type Result struct {
    Message string
    Status int
}

func (c App) Index() revel.Result {
    var widgets []models.Widget = models.GetWidgets()
	return c.Render(widgets)
}

func (c App) Update() revel.Result {
    var widget models.Widget
    b, err:= ioutil.ReadAll(c.Request.Body)

    if err != nil {
        log.Fatal("Problem reading the body: " + err.Error())
    }

    if err := json.Unmarshal(b, &widget); err != nil {
        log.Fatal("Problem with Unmarshal: " + err.Error())
    }

    models.SetWidget(widget)

    return c.RenderJson(widget)
}
