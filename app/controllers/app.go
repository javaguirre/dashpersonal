package controllers

import (
    "encoding/json"
    "github.com/robfig/revel"
    "github.com/robfig/revel/cache"
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
    var widgets []models.Widget

    if err:= cache.Get("widgets", &widgets); err != nil {
        log.Printf("Cache not working or first time request")
        widgets = models.GetWidgets()
        go cache.Set("widgets", widgets, cache.DEFAULT)
    }

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
    go cache.Delete("widgets")

    return c.RenderJson(widget)
}
