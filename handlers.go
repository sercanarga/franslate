package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

type Handler struct{}

var internal = &Internal{}
var service = &Service{}

func (h *Handler) InputBoxChanged(t string) {
	if t == "" {
		outputBox.SetText("")
		return
	}

	h.translateHandler()
}

func (h *Handler) translateHandler() {
	s := internal.GetSettingsFile()
	if s.ApiKey != "" {
		if inputBox.Text != "" {
			service.Translate(s.InputLanguage, s.OutputLanguage, s.ApiKey, inputBox.Text)
		}
	} else {
		outputBox.SetText("Please enter your API Key in the settings.")
	}

}

func (h *Handler) NewLanguageSelect(languages []string, defaultLang string, onChange func(string)) *widget.Select {
	selectWidget := widget.NewSelect(languages, onChange)
	selectWidget.SetSelected(defaultLang)

	return selectWidget
}

func (h *Handler) InputLangChange(value string) {
	internal.SyncSettingsFile(&Settings{
		InputLanguage: value,
	})

	h.translateHandler()
}

func (h *Handler) OutputLangChange(value string) {
	internal.SyncSettingsFile(&Settings{
		OutputLanguage: value,
	})

	h.translateHandler()
}

func (h *Handler) SwitchButtonClick(inputLangSelect, outputLangSelect *widget.Select) {
	inputLang := inputLangSelect.Selected
	outputLang := outputLangSelect.Selected
	inputLangSelect.SetSelected(outputLang)
	outputLangSelect.SetSelected(inputLang)

	internal.SyncSettingsFile(&Settings{
		InputLanguage:  outputLang,
		OutputLanguage: inputLang,
	})

	h.translateHandler()
}

func (h *Handler) ClearInputBox() {
	inputBox.SetText("")
}

func (h *Handler) CopyOutputToClipboard() {
	w.Clipboard().SetContent(outputBox.Text)
}

func (h *Handler) SettingsButtonClick() {
	settingsWindow := a.NewWindow(fmt.Sprintf("%s - %s", appName, "Settings"))
	settingsWindow.Resize(fyne.NewSize(600, 200))
	settingsWindow.CenterOnScreen()

	apiKeyEntry := widget.NewEntry()
	apiKeyEntry.SetPlaceHolder("AIzaSyBMqGQu_lWIj6dG__yzxzgN3S9yB1Zhgmo")
	apiKeyEntry.Text = internal.GetSettingsFile().ApiKey

	inputDelayEntry := widget.NewEntry()
	inputDelayEntry.SetPlaceHolder("300")
	inputDelayEntry.Text = internal.GetSettingsFile().InputDelay

	settingsFileEntry := widget.NewEntry()
	settingsFileEntry.Text = internal.GetDataPath() + "/settings.json"

	getApiKeyLink, _ := url.Parse("https://aistudio.google.com/app/apikey")
	contributeLink, _ := url.Parse("https://github.com/sercanarga/franslate")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:     "API Key:",
				Widget:   apiKeyEntry,
				HintText: "Enter your Gemini API Key",
			},
			{
				Text:     "Input Delay(ms):",
				Widget:   inputDelayEntry,
				HintText: "Enter the input delay in milliseconds",
			},
			{
				Text:     "Settings:",
				Widget:   settingsFileEntry,
				HintText: "Just for information (not editable)",
			},
			{
				Widget: container.NewHBox(
					widget.NewHyperlink("Get API Key", getApiKeyLink),
					widget.NewHyperlink("Contribute", contributeLink),
				),
			},
		},
		OnSubmit: func() {
			if apiKeyEntry.Text != "" {
				internal.SyncSettingsFile(&Settings{
					ApiKey: apiKeyEntry.Text,
				})
			}

			if inputDelayEntry.Text != "" {
				internal.SyncSettingsFile(&Settings{
					InputDelay: inputDelayEntry.Text,
				})
			}

			settingsWindow.Close()
		},
		SubmitText: "Save",
	}

	settingsWindow.SetContent(form)
	settingsWindow.Show()
}
