package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Helper struct{}

func (h *Helper) CreateHeader(inputLangSelect, outputLangSelect *widget.Select, switchButton *widget.Button) fyne.CanvasObject {
	return container.NewBorder(nil, nil,
		container.NewHBox(inputLangSelect),
		container.NewHBox(outputLangSelect),
		container.NewHBox(layout.NewSpacer(), switchButton, layout.NewSpacer()),
	)
}

func (h *Helper) NewClearButton(onClick func()) *widget.Button {
	clearButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), onClick)
	clearButton.Importance = widget.LowImportance
	return clearButton
}

func (h *Helper) CreateInputContainer(clearButton *widget.Button) fyne.CanvasObject {
	return container.NewStack(
		inputBox,
		container.NewStack(
			container.NewVBox(
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), container.NewPadded(clearButton)),
				layout.NewSpacer(),
			),
		),
	)
}

func (h *Helper) NewCopyButton(onClick func()) *widget.Button {
	copyButton := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), onClick)
	copyButton.Importance = widget.LowImportance
	return copyButton
}

func (h *Helper) CreateOutputContainer(copyButton *widget.Button, settingsButton *widget.Button) fyne.CanvasObject {
	// @todo: buttons stay on the text. should be fix.
	return container.NewStack(
		outputBox,
		container.NewStack(
			container.NewVBox(
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), container.NewPadded(settingsButton), layout.NewSpacer(), container.NewPadded(copyButton)),
			),
		),
	)
}
