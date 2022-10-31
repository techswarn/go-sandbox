package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but received %v", d[0])
	}

	if d[len(d) - 1] != "Four of clubs" {
		t.Errorf("Expected Four of Clubs but received %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck :=newDeck()
	deck.saveToFile("_deckTesting")

	newDeckFromFile := newDeckFromFile("_deckTesting")
	if len(newDeckFromFile) != 16 {
		t.Errorf("Expected 20 Cards, but received %v", len(newDeckFromFile))
	}
	os.Remove("_decktesting")
}