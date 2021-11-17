package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var (
	appWindow fyne.Window
)

func main() {
	myApp := app.New()
	appWindow = myApp.NewWindow("Picture Frame")

	image := canvas.NewImageFromFile(getImage())
	image.FillMode = canvas.ImageFillContain

	bottom := widget.NewButton("Quit", func() {
		appWindow.Close()
	})
	middle := image

	content := container.NewBorder(nil, bottom, nil, nil, middle)
	appWindow.SetContent(content)

	myApp.Settings().SetTheme(myTheme{})
	appWindow.SetFullScreen(true)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			image.File = getImage()
			image.Refresh()
		}
	}()

	appWindow.ShowAndRun()
}

func getImage() string {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(files))
	name := files[index].Name()

	return fmt.Sprintf("images/%s", name)
}
