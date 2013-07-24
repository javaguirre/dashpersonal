package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    type Widget struct {
        Row int
        Col int
        Sizex int
        Sizey int
        Content string
    }
    widget := Widget{1, 1, 1, 1, "my content"}
    widgets := []Widget{widget}
	return c.Render(widgets)
}
