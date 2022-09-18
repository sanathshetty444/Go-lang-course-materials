package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := createDeck()
	if len(d) != 16 {
		t.Errorf("Expected length is 16 got %v", len(d))
	}
}

func TestSaveAndCreateDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := createDeck()
	d.saveToFile("_decktesting")
	d = createDeckFromFile("_deckTesting")
	if len(d) != 16 {
		t.Errorf("Expected length is 16 got %v", len(d))
	}

	os.Remove("_decktesting")
}
