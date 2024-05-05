package discard

import "unoSimulator/internal/deck"

type Discard []deck.Card

func CreateDiscard()Discard{
	var d Discard
	return d
}