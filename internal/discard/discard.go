package discard

import (
	"unoSimulator/internal/deck"
)

type Discard []deck.Card

func CreateDiscard()Discard{
	var d Discard
	return d
}

func (d Discard) ValidPlay(c deck.Card) bool{
	topCard:=d[len(d)-1]
	if(topCard.Color==c.Color || topCard.Number==c.Number || c.Color=="Any" || (topCard.Power==c.Power && topCard.Power!="")){
		return true
	}
	return false
}

func (d *Discard) Discard(d2 *deck.Deck, n int){
	*d = append(*d, (*d2)[n])
	*d2 = append((*d2)[:n], (*d2)[n+1:]...)
}

func (d Discard) TopCard() deck.Card {
	return d[len(d)-1]
}