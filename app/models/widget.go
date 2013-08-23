package models

import (
    "runtime"
    "path"
    "encoding/json"
    "io/ioutil"
    "log"
)


type Widget struct {
    Id  string
    Row string
    Col string
    Sizex string
    Sizey string
    Content string
}

func getFileDb() string {
    _, filename, _, _ := runtime.Caller(1)
    filedb := path.Join(path.Dir(filename), "../../settings.cfg")

    return filedb
}

func ReadFile() []byte {
    content, err := ioutil.ReadFile(getFileDb())

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

    write_err := ioutil.WriteFile(getFileDb(), []byte(content), 644)

    if write_err != nil {
        log.Fatal("Problem writing the file: " + err.Error())
    }
}

func SetWidget(widget Widget) {
    widgets := GetWidgets()

    if widget.Id == "X" {
        widget.Id = string(len(widgets) + 1)
        widgets = append(widgets, widget)
    } else {
        for i:= 0; i < len(widgets); i++ {
            if widget.Id == widgets[i].Id {
                widgets[i].Content = widget.Content
            }
        }
    }


    SaveWidgets(widgets)
}
