package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

var (
	appName   = "FranslateAI"
	appDesc   = "AI Based Free Translate App"
	languages = []string{"English", "French", "German", "Spanish", "Turkish"}
	appWindow fyne.Window
	inputBox  *widget.Entry
	outputBox *widget.Entry

	handler = &Handler{}
	helper  = &Helper{}
)

func main() {
	// Set the scale of the app to 1.2
	os.Setenv("FYNE_SCALE", "1.2")

	app := app.New()
	appWindow = app.NewWindow(fmt.Sprintf("%s - %s", appName, appDesc))
	appWindow.Resize(fyne.NewSize(650, 400))

	appWindow.SetCloseIntercept(func() {
		appWindow.Hide()
	})

	if desk, ok := app.(desktop.App); ok {
		m := fyne.NewMenu(appName,
			fyne.NewMenuItem("Translate Text", func() {
				appWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	inputBox = widget.NewMultiLineEntry()
	outputBox = widget.NewMultiLineEntry()

	appWindow.SetContent(createUI())

	appWindow.Canvas().Focus(inputBox)

	appWindow.ShowAndRun()
}

func createUI() fyne.CanvasObject {
	inputLangSelect := handler.NewLanguageSelect(languages, "English", handler.InputLangChange)
	outputLangSelect := handler.NewLanguageSelect(languages, "Turkish", handler.OutputLangChange)

	switchButton := widget.NewButtonWithIcon("Switch", theme.ViewRefreshIcon(), func() {
		handler.SwitchButtonClick(inputLangSelect, outputLangSelect)
	})

	header := helper.CreateHeader(inputLangSelect, outputLangSelect, switchButton)

	inputBox.SetPlaceHolder("Enter Text to Translate")
	inputBox.Wrapping = fyne.TextWrapWord

	clearButton := helper.NewClearButton(handler.ClearInputBox)
	inputContainer := helper.CreateInputContainer(clearButton)

	outputBox.Disable()
	copyButton := helper.NewCopyButton(handler.CopyOutputToClipboard)
	outputContainer := helper.CreateOutputContainer(copyButton)

	return container.NewBorder(header, nil, nil, nil,
		container.NewGridWithColumns(2, inputContainer, outputContainer),
	)
}
