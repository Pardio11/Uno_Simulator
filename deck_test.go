package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 108 {
		t.Errorf("Length of new deck is not the expected.\nExpected: %v Actual length: %v",108,len(d))
	}
}

func TestSaveDeckAndLoadDeck(t *testing.T) {
	os.Remove("_unoDeckTesting")
	d := newDeck()
	d.saveToFile("_unoDeckTesting")
	loadDeck := newDeckFromFile("_unoDeckTesting")
	passed := true
	for i := range d {
		if d[i]!=loadDeck[i]{
			passed = false
		}
	}
	
	if len(loadDeck) != len(d) {
		passed = false
	}

	if !passed {
		t.Errorf("The loaded deck is not the same as created")
	}
	os.Remove("_unoDeckTesting")
}
