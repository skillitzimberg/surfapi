package datautil

import (
	"reflect"
	"testing"
)

func TestGetBouyData(t *testing.T) {
	rawBouyData := GetBouyData()
	var testData []byte

	if reflect.TypeOf(rawBouyData) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, rawBouyData)
	}
}

func TestHandleRawData(t *testing.T) {
	rawBouyData := GetBouyData()
	bouyData := HandleRawData(rawBouyData)
	var testData []string

	if reflect.TypeOf(bouyData) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, bouyData)
	}
}

func TestGetSwellHeight(t *testing.T) {
	rawBouyData := GetBouyData()
	bouyData := HandleRawData(rawBouyData)
	swellHeight := GetSwellHeight(bouyData)
	var testData float64

	if reflect.TypeOf(swellHeight) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, swellHeight)
	}
}
