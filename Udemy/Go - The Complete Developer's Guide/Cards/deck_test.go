package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card in deck is expected to be Ace of Clubs, but is %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card in deck is expected to be King of Diamonds, but is %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_desktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected %v cards in deck, but got %v", len(d), len(loadedDeck))
	}

	os.Remove("_decktesting")
}
