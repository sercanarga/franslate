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
	"runtime"
	"time"
)

var (
	a         fyne.App
	w         fyne.Window
	inputBox  *widget.Entry
	outputBox *widget.Entry
	handler   = &Handler{}
	helper    = &Helper{}

	appName    = "FranslateAI"
	appDesc    = "AI Based Free Translate App"
	languages  = []string{"English", "French", "German", "Spanish", "Turkish"}
	inputDelay = 300 * time.Millisecond
)

func main() {
	// Set the scale of the app to 1.1
	os.Setenv("FYNE_SCALE", "1.1")

	a = app.New()
	w = a.NewWindow(fmt.Sprintf("%s - %s", appName, appDesc))
	w.Resize(fyne.NewSize(650, 400))

	// https://github.com/fyne-io/fyne/issues/3197
	if runtime.GOOS != "darwin" {
		w.SetCloseIntercept(func() {
			w.Hide()
		})

		if desk, ok := a.(desktop.App); ok {
			m := fyne.NewMenu(appName,
				fyne.NewMenuItem("Translate Text", func() {
					w.Show()
				}))
			desk.SetSystemTrayMenu(m)
		}
	}

	inputBox = widget.NewMultiLineEntry()
	outputBox = widget.NewMultiLineEntry()

	w.SetContent(createUI())

	w.Canvas().Focus(inputBox)

	w.ShowAndRun()

}

func createUI() fyne.CanvasObject {
	currentSettings := internal.GetSettingsFile()
	inputLangSelect := handler.NewLanguageSelect(languages, currentSettings.InputLanguage, handler.InputLangChange)
	outputLangSelect := handler.NewLanguageSelect(languages, currentSettings.OutputLanguage, handler.OutputLangChange)

	switchButton := widget.NewButtonWithIcon("", theme.NewThemedResource(fyne.NewStaticResource("switch-lang.svg", []byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="none" width="1" height="1"><path d="M0 224c0 17.7 14.3 32 32 32s32-14.3 32-32c0-53 43-96 96-96H320v32c0 12.9 7.8 24.6 19.8 29.6s25.7 2.2 34.9-6.9l64-64c12.5-12.5 12.5-32.8 0-45.3l-64-64c-9.2-9.2-22.9-11.9-34.9-6.9S320 19.1 320 32V64H160C71.6 64 0 135.6 0 224zm512 64c0-17.7-14.3-32-32-32s-32 14.3-32 32c0 53-43 96-96 96H192V352c0-12.9-7.8-24.6-19.8-29.6s-25.7-2.2-34.9 6.9l-64 64c-12.5 12.5-12.5 32.8 0 45.3l64 64c9.2 9.2 22.9 11.9 34.9 6.9s19.8-16.6 19.8-29.6V448H352c88.4 0 160-71.6 160-160z"/></svg>`))), func() {
		handler.SwitchButtonClick(inputLangSelect, outputLangSelect)
	})

	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), handler.SettingsButtonClick)
	settingsButton.Importance = widget.LowImportance

	header := helper.CreateHeader(inputLangSelect, outputLangSelect, switchButton)

	var inputBoxDelay *time.Timer
	inputBox.SetPlaceHolder("Enter Text to Translate")
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.OnChanged = func(t string) {
		if inputBoxDelay != nil {
			inputBoxDelay.Stop()
		}
		inputBoxDelay = time.AfterFunc(inputDelay, func() {
			handler.InputBoxChanged(t)
		})
	}

	clearButton := helper.NewClearButton(handler.ClearInputBox)
	inputContainer := helper.CreateInputContainer(clearButton)

	outputBox.Wrapping = fyne.TextWrapWord
	copyButton := helper.NewCopyButton(handler.CopyOutputToClipboard)
	outputContainer := helper.CreateOutputContainer(copyButton, settingsButton)

	return container.NewBorder(header, nil, nil, nil,
		container.NewGridWithColumns(2, inputContainer, outputContainer),
	)
}
