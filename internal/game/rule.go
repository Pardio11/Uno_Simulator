package game

func (g *Game) playCard(index int) {
	card := g.NextPlayer().Deck[index]
	deckPlayer := &(g.NextPlayer().Deck)
	g.DiscardDeck.Discard(deckPlayer, index)
	switch {
	case card.Power == "+4":
		g.CardCount += 4
		g.changeTurn()
	case card.Power == "+2":
		g.CardCount += 2
		g.changeTurn()

	case card.Power == "Skip":
		g.changeTurn()

	case card.Power == "Reverse":
		g.Direction = !g.Direction
	}
	g.changeTurn()

}

func (g *Game) changeTurn() {
	if !g.Direction {
		g.Turn = (g.Turn + 1) % len(g.Players)
	} else {
		if g.Turn-1 < 0 {
			g.Turn = len(g.Players) - 1
		} else {
			g.Turn = (g.Turn - 1) % len(g.Players)
		}
	}
}