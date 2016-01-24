package threewordgame

import (
	ui "github.com/gizak/termui"
	"time"
)

var uiStory = ui.NewPar("")
var uiInput = ui.NewPar("")
var uiStatus = ui.NewPar("")
var inputChannel = make(chan string)
var inputData string

func init() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	uiStory.X = 0
	uiStory.Y = 0
	uiStory.Height = 15
	uiStory.Width = 80
	uiStory.TextFgColor = ui.ColorWhite
	uiStory.BorderLabel = "Story"
	uiStory.BorderFg = ui.ColorCyan

	uiInput.X = 0
	uiInput.Y = 15
	uiInput.Height = 3
	uiInput.Width = 80
	uiInput.TextFgColor = ui.ColorWhite
	uiInput.BorderLabel = "Input"
	uiInput.BorderFg = ui.ColorCyan

	uiStatus.Text = "Initializing"
	uiStatus.X = 0
	uiStatus.Y = 18
	uiStatus.Height = 3
	uiStatus.Width = 80
	uiStatus.TextFgColor = ui.ColorWhite
	uiStatus.BorderLabel = "Status"
	uiStatus.BorderFg = ui.ColorCyan

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		Exit()
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

	draw()
	ui.Loop()
}

func Test() {
	uiStatus.Text = "Initializing."
	draw()
	time.Sleep(time.Second)

	uiStatus.Text = "Initializing.."
	draw()
	time.Sleep(time.Second)

	uiStatus.Text = "Initializing..."
	draw()
	time.Sleep(time.Second)

	uiStatus.Text = "Initializing...."
	draw()
}

func draw() {
	ui.Render(uiStory, uiInput, uiStatus)
}

// Exit stops the ui loop
func Exit() {
	ui.StopLoop()
}

// ClearStory clears the story
func ClearStory() {
	uiStory.Text = ""
	draw()
}

// AddStoryPart adds a part to the story
func AddStoryPart(part string) {
	uiStory.Text = uiStory.Text + " " + part
	draw()
}

// GetInputChannel gets the inputChannel
func GetInputChannel() chan string {
	return inputChannel
}

// SetStatus sets the status
func SetStatus(status string) {
	uiStatus.Text = status
	draw()
}
