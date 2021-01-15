package main

import (
	"math/rand"
	"time"
)

func main() {

	game := &Game{
		Cards:          make([]int, 14),
		Players:        make([]*Player, 5),
		PendingPlayers: make([]*Player, 0),
		Decks:          8,
		TurnTimeout:    10,
		State:          Idle,
		playerCount:    0,

		HouseVisible: false,
		HouseScore:   0,
		HouseCards:   make([]int, 0),
	}

	print(game.Players[0] == nil)

	game.GenerateNextCard()

}

//phase 0 Betting

//phase1 Dealing

//phase 2 Turns

//phase 3 House tries to beat everyone

//House hits until he has 17 or more

//Process winnings or losses

//State type for enum
type State int

const (
	//Idle waiting for players to resume game
	Idle = iota
	//Betting Players are placing their bets for upcoming round
	Betting
	//InitialDealing House is dealing initial pair of cards
	InitialDealing
	//Round game is processing player's turns
	Round
	//House tries to reach 17 or more to beat players
	House
	//Results final result is calculated, bets are paid or taken
	Results
	//Restarting game is restarting to begin next round
	Restarting
)

//Game blackjack
type Game struct {
	Cards          []int
	Players        []*Player
	Decks          int
	HouseVisible   bool
	HouseScore     int
	HouseCards     []int
	PendingPlayers []*Player
	TurnTimeout    int
	State          State
	currentTurn    *Player
	playerCount    int
}

//PlayerJoin adds a player when he joins the game
func (g *Game) PlayerJoin(playerID string) {

}

//Player represents player in game
type Player struct {
	Position  int
	ID        string
	Cards     []int
	Score     int
	BetAmount float32
}

//UpdateScore updates the players total based on cards
func (p *Player) UpdateScore() {

	index := 0

	for index < len(p.Cards) {
		p.Score += p.Cards[index]
		index++
	}
}

//RandomNumber gens random num
func RandomNumber() int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	return rand.Intn(12) + 1

}

//CountCards remaining cards
func (g *Game) CountCards() int {

	index := 1
	cardsLeft := 0

	for index <= 13 {
		cardsLeft += g.Cards[index]
		index++
	}

	return cardsLeft
}

//GenerateNextCard generate a random number
func (g *Game) GenerateNextCard() int {

	if g.CountCards() == 0 {
		g.RestockDeck()
	}

	num := RandomNumber()

	if g.Cards[num] == 0 {
		g.GenerateNextCard()
	}

	return num

}

//NextPlayer pass turn to next player
func (g *Game) NextPlayer() {

	playerIndex := g.currentTurn.Position + 1

	if playerIndex == len(g.Players) {
		playerIndex = 0
	}
	g.currentTurn = g.Players[playerIndex]

}

//Stand dont receive any cards
func (g *Game) Stand(playetID string) {
	g.NextPlayer()
}

//DealCards deal cards at start of game
func (g *Game) DealCards() {
	g.DealOne()
	g.DealOneHouse()
	g.DealOne()
	g.DealOneHouse()
}

//DealOneHouse deal one to the house
func (g *Game) DealOneHouse() {

	num := g.GenerateNextCard()

	if len(g.HouseCards) < 3 {
		g.HouseVisible = false
	} else {
		g.HouseVisible = true
	}

	g.HouseCards = append(g.HouseCards, num)

	score := 0

	index := 0

	for index < len(g.HouseCards) {
		score += g.HouseCards[index]
		index++
	}

	g.HouseScore = score

}

//DealOne House deals a card to player
func (g *Game) DealOne() {

	index := 0

	for index < len(g.Players) {

		player := g.Players[index]

		if player == nil {
			continue
		}

		num := g.GenerateNextCard()

		//add cards to player's deck
		player.Cards = append(player.Cards, num)
		player.UpdateScore()
		//update total

	}

}

//Hit give another card to the brave gentleman
func (g *Game) Hit(playerID string) {

	//if not in playing state do nothing
	if g.State != Round {
		return
	}

	//if not your turn do nothing
	if g.currentTurn.ID != playerID {
		return
	}

	num := g.GenerateNextCard()

	//add cards to player's deck
	g.currentTurn.Cards = append(g.currentTurn.Cards, num)

	//update total
	g.currentTurn.UpdateScore()

}

//RestockDeck restocks cards
func (g *Game) RestockDeck() {

	index := 1

	//init total cards
	for index <= 13 {

		g.Cards[index] = 4 * g.Decks
		index++

	}
}
