package main

import (
	"bufio"
	"flag"
	"fmt"
	ui "github.com/gizak/termui"
	"net"
	"os"
)

var uiStory = ui.NewPar("")
var uiInput = ui.NewPar("")
var uiStatus = ui.NewPar("")
var inputChannel = make(chan string)
var inputData string

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	uiStory.X = 0
	uiStory.Y = 0
	uiStory.Height = 15
	uiStory.TextFgColor = ui.ColorWhite
	uiStory.BorderLabel = "Story"
	uiStory.BorderFg = ui.ColorCyan

	uiInput.X = 0
	uiInput.Y = 15
	uiInput.Height = 3
	uiInput.TextFgColor = ui.ColorWhite
	uiInput.BorderLabel = "Input"
	uiInput.BorderFg = ui.ColorCyan

	uiStatus.Text = "Initializing"
	uiStatus.X = 0
	uiStatus.Y = 18
	uiStatus.Height = 3
	uiStatus.TextFgColor = ui.ColorWhite
	uiStatus.BorderLabel = "Status"
	uiStatus.BorderFg = ui.ColorCyan

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, uiStory)),
		ui.NewRow(
			ui.NewCol(12, 0, uiInput)),
		ui.NewRow(
			ui.NewCol(12, 0, uiStatus)))

	// calculate layout
	ui.Body.Align()

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		exit()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		// handle all other key pressing
		k, ok := e.Data.(ui.EvtKbd)
		if ok {
			if k.KeyStr == "C-8" {
				l := len(inputData)
				if l >= 1 {
					inputData = inputData[0 : l-1]
				}
			} else if k.KeyStr == "<enter>" && len(inputData) > 0 {
				inputChannel <- inputData
				inputData = ""
			} else if k.KeyStr == "<space>" {
				inputData = inputData + " "
			} else if len(k.KeyStr) == 1 {
				inputData = inputData + k.KeyStr
			}
		}

		uiInput.Text = inputData
		draw()
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Render(ui.Body)
	})

	ui.Handle("/update/status", func(e ui.Event) {
		msg := e.Data.(string)
		uiStatus.Text = msg
		draw()
	})

	draw()
	ui.Loop()

	// Conf && Connect
	var server string
	var port int

	flag.StringVar(&server, "server", "127.0.0.1", "Server host")
	flag.IntVar(&port, "port", 6666, "Server port")
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		ui.SendCustomEvt("/update/status", "Error dialing in.")
		exit()
	}

	ui.SendCustomEvt("/update/status", "Connected.")

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			ui.SendCustomEvt("/update/status", "Error reading from Stdin.")
		}

		fmt.Fprintf(conn, input)
	}

	ui.SendCustomEvt("/update/status", "Exiting.")
}

func draw() {
	ui.Render(uiStory, uiInput, uiStatus)
}

func exit() {
	ui.StopLoop()
	os.Exit(-1)
}
