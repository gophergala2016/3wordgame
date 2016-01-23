package threewordgame

import ui "github.com/gizak/termui"

var ui_story *ui.Par
var ui_input *ui.Par
var ui_status *ui.Par
var input_channel = make(chan string)

func SetupFrontend() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	input_data := ""

	ui_story = ui.NewPar("")
	ui_story.X = 0
	ui_story.Y = 0
	ui_story.Height = 15
	ui_story.Width = 80
	ui_story.TextFgColor = ui.ColorWhite
	ui_story.BorderLabel = "Story"
	ui_story.BorderFg = ui.ColorCyan

	ui_input = ui.NewPar("")
	ui_input.X = 0
	ui_input.Y = 15
	ui_input.Height = 3
	ui_input.Width = 80
	ui_input.TextFgColor = ui.ColorWhite
	ui_input.BorderLabel = "Input"
	ui_input.BorderFg = ui.ColorCyan

	ui_status = ui.NewPar("Not Connected")
	ui_status.X = 0
	ui_status.Y = 18
	ui_status.Height = 3
	ui_status.Width = 80
	ui_status.TextFgColor = ui.ColorWhite
	ui_status.BorderLabel = "Status"
	ui_status.BorderFg = ui.ColorCyan

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		// handle all other key pressing
		k, ok := e.Data.(ui.EvtKbd)
		if ok {
			if k.KeyStr == "C-8" {
				l := len(input_data)
				if l >= 1 {
					input_data = input_data[0 : l-1]
				}
			} else if k.KeyStr == "<enter>" && len(input_data) > 0 {
				input_channel <- input_data
				input_data = ""
			} else if len(k.KeyStr) == 1 {
				input_data = input_data + k.KeyStr
			}
		}

		ui_input.Text = input_data
		draw()
	})

	draw()
	ui.Loop()
}

func draw() {
	ui.Render(ui_story, ui_input, ui_status)
}

func ClearStory() {
	ui_story.Text = ""
	draw()
}

func AddStoryPart(part string) {
	ui_story.Text = ui_story.Text + " " + part
	draw()
}

func GetInputChannel() chan string {
	return input_channel
}

func SetStatus(status string) {
	ui_status.Text = status
}
