package controllers

import (
    "encoding/json"
    "github.com/robfig/revel"
    "dashpersonal/app/models"
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
    var result Result
    b, err:= ioutil.ReadAll(c.Request.Body)

    if err != nil {
        log.Fatal("Problem reading the body: " + err.Error())
    }

    log.Printf(string(b))
    err2 := json.Unmarshal(b, &widget)

    if err2 != nil {
        log.Fatal("Problem with Unmarshal: " + err2.Error())
    }
    models.SetWidget(widget)

    return c.RenderJson(result)
}
