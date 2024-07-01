package main

import "fyne.io/fyne/v2/widget"

type Handler struct{}

func (h *Handler) NewLanguageSelect(languages []string, defaultLang string, onChange func(string)) *widget.Select {
	selectWidget := widget.NewSelect(languages, onChange)
	selectWidget.SetSelected(defaultLang)
	return selectWidget
}

func (h *Handler) InputLangChange(value string) {
	// Handle input language change
}

func (h *Handler) OutputLangChange(value string) {
	// Handle output language change
}

func (h *Handler) SwitchButtonClick(inputLangSelect, outputLangSelect *widget.Select) {
	inputLang := inputLangSelect.Selected
	outputLang := outputLangSelect.Selected
	inputLangSelect.SetSelected(outputLang)
	outputLangSelect.SetSelected(inputLang)
}

func (h *Handler) ClearInputBox() {
	inputBox.SetText("")
}

func (h *Handler) CopyOutputToClipboard() {
	appWindow.Clipboard().SetContent(outputBox.Text)
}
