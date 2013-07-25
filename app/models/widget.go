package models

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

const Filedb = "/home/javaguirre/Proyectos/go/src/dashpersonal/settings.cfg"

type Widget struct {
    Id  string
    Row string
    Col string
    Sizex string
    Sizey string
    Content string
}

func ReadFile() []byte {
    content, err := ioutil.ReadFile(Filedb)

    if err != nil {
        log.Fatal("Problem with the file!: " + err.Error())
    }

    return content
}

func GetWidgets() []Widget {
    var widgets []Widget
    content := ReadFile()

    if err := json.Unmarshal([]byte(content), &widgets); err != nil {
        log.Fatal("Problem with the Unmarshal: " + err.Error())
    }

    return widgets
}

func SaveWidgets(widgets []Widget) {
    content, err := json.Marshal(widgets)

    if err != nil {
        log.Fatal("Problem with the Marshal: " + err.Error())
    }

    write_err := ioutil.WriteFile(Filedb, []byte(content), 644)

    if write_err != nil {
        log.Fatal("Problem writing the file: " + err.Error())
    }
}

func SetWidget(widget Widget) {
    widgets := GetWidgets()
    exists := false

    for i:= 0; i < len(widgets); i++ {
        if widget.Id == widgets[i].Id {
            widgets[i] = widget
            exists = true
        }
    }

    if !exists {
        widgets = append(widgets, widget)
    }

    SaveWidgets(widgets)
}
