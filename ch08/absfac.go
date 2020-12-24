package main

import (
	"errors"
	"fmt"
)

// Button interface
type Button interface {
	Paint()
	OnClick()
}

// Label interface
type Label interface {
	Paint()
}

// WinButton is implementation of Button on Windows OS
type WinButton struct{}

// Paint draws element
func (WinButton) Paint() { fmt.Println("win button paint") }

// OnClick acts on click
func (WinButton) OnClick() { fmt.Println("win button click") }

// WinLabel is implementation of Label on Windows OS
type WinLabel struct{}

// Paint is implementation of Label on Windows OS
func (WinLabel) Paint() { fmt.Println("win label paint") }

// MacButton is implementation of Button on Mac OS
type MacButton struct{}

// Paint draws element
func (MacButton) Paint() { fmt.Println("mac button paint") }

// OnClick acts on click
func (MacButton) OnClick() { fmt.Println("mac button click") }

// MacLabel is implementation of Label on Mac OS
type MacLabel struct{}

// Paint is implementation of Label on Mac OS
func (MacLabel) Paint() { fmt.Println("mac label paint") }

// UIFactory is abstract factory
type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

// WinFactory is implementation of UIFactory for Windows OS
type WinFactory struct{}

// CreateButton returns WinButton
func (WinFactory) CreateButton() Button {
	return WinButton{}
}

// CreateLabel returns WinLabel
func (WinFactory) CreateLabel() Label {
	return WinLabel{}
}

// MacFactory is implementation of UIFactory for Mac OS
type MacFactory struct{}

// CreateButton returns MacButton
func (MacFactory) CreateButton() Button {
	return MacButton{}
}

// CreateLabel returns MacLabel
func (MacFactory) CreateLabel() Label {
	return MacLabel{}
}

// CreateFactory returns UIFactory for OS and error
func CreateFactory(os string) (UIFactory, error) {
	switch os {
	case "win":
		return WinFactory{}, nil
	case "mac":
		return MacFactory{}, nil
	}
	return nil, errors.New("CreateFactory: unknown os")
}

// Run draw UI
func Run(f UIFactory) {
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}

func main() {
	winUI, err := CreateFactory("win")
	if err != nil {
		fmt.Println(err)
		return
	}
	Run(winUI)

	macUI, err := CreateFactory("mac")
	if err != nil {
		fmt.Println(err)
		return
	}
	Run(macUI)
}
