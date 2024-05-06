package game

func (g *Game) playCard(index int) {
	card := g.NextPlayer().Deck[index]
	deckPlayer := &(g.NextPlayer().Deck)
	g.DiscardDeck.Discard(deckPlayer, index)
	switch {
	case card.Power == "+4":
		g.cardCount += 2
		g.changeTurn()
	case card.Power == "+2":
		g.cardCount += 2
		g.changeTurn()

	case card.Power == "Skip":
		g.changeTurn()

	case card.Power == "Reverse":
		g.direction = !g.direction
	}
	g.changeTurn()

}

func (g *Game) changeTurn() {
	if !g.direction {
		g.Turn = (g.Turn + 1) % len(g.players)
	} else {
		if g.Turn-1 < 0 {
			g.Turn = len(g.players) - 1
		} else {
			g.Turn = (g.Turn - 1) % len(g.players)
		}
	}
}