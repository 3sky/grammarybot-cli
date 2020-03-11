package main

import (
	"strings"
	"testing"
)

func TestCheckText(t *testing.T) {
	LANGAUGE := "en-US"
	URL := "http://api.grammarbot.io/v2/check"
	botToken := "XYZ"
	text := "I can't remember how to go their"

	err := CheckText(LANGAUGE, URL, botToken, text)
	if err != nil {
		t.Errorf("Error with CheckText funtion")
	}
}

func TestLoadFile(t *testing.T) {

	PATH := "go.mod"
	str, err := LoadFile(PATH)
	if err != nil {
		t.Errorf("Error with TestLoadFile funtion")
	}
	if !(strings.Contains(str, "github.com/3sky/grammarybot-cli")) {
		t.Errorf("Error with TestLoadFile, string is wrong")
	}
}
