package euchre

import (
	"errors"
	"math/rand/v2"
	"slices"
)

type Suit byte

const (
	NoSuit Suit = iota
	ClubsSuit
	SpadesSuit
	HeartsSuit
	DiamondsSuit
)

type Rank byte

const (
	Seven Rank = iota + 7 // 7
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
}

var allCards []Card

func init() {
	var allSuits = [...]Suit{ClubsSuit, SpadesSuit, HeartsSuit, DiamondsSuit}
	var allCardValues = [...]Rank{
		Ace,
		King,
		Queen,
		Jack,
		Ten,
		Nine,
	}
	allCards = make([]Card, 0, len(allSuits)*len(allCardValues))

	// for _, value := range allCardValues[:] {
	// 	for _, suit := range allSuits[:] {
	// 		allCards = append(allCards, Card{value: value, suit: suit})
	// 	}
	// }
}

type Player string

type EuchreGame struct {
	currentPlayer Player
	leadingSuit   Suit
	playerCards   map[Player][]Card
	kittyCards    []Card
}

func (g *EuchreGame) InitPlayers() {
	g.playerCards = make(map[Player][]Card, 4)
}

var (
	ErrNotPlayerTurn   = errors.New("not player turn")
	ErrUnknownPlayer   = errors.New("unknown player")
	ErrNotOwnedCard    = errors.New("card not owned")
	ErrInvalidCardSuit = errors.New("invalid card suit")
)

const NumberOfPlayers = 4

// splits the cards per payer and add the remainder to the kitty
func (g *EuchreGame) DealCards() {
	perm := rand.Perm(len(allCards))
	i := 0
	for player := range g.playerCards {
		for range 5 {
			g.playerCards[player] = append(g.playerCards[player], allCards[perm[i]])
			i++
		}
	}
	for range 4 {
		g.kittyCards = append(g.kittyCards, allCards[perm[i]])
		i++
	}
}

// need a way to validate the cards played by the players
func (g *EuchreGame) ValidateCardPlayed(card Card, player Player) error {
	if player != g.currentPlayer {
		return ErrNotPlayerTurn
	}
	cards, exist := g.playerCards[player]
	if !exist {
		return ErrUnknownPlayer
	}
	if !slices.Contains(cards, card) {
		return ErrNotOwnedCard
	}
	if g.leadingSuit == NoSuit {
		g.leadingSuit = card.Suit
	} else if g.leadingSuit != card.Suit && CardsContainSuit(cards, g.leadingSuit) {
		return ErrInvalidCardSuit
	}
	return nil
}

func CardsContainSuit(cards []Card, suit Suit) bool {
	return slices.ContainsFunc(cards, func(e Card) bool { return e.Suit == suit })
}

func RemoveCard(cards []Card, card Card) []Card {
	return slices.DeleteFunc(cards, func(e Card) bool { return e == card })
}

// 1- deal cards
// 2- send cards to each connected player
// 		- send also the first card in the kitty
// 			cards: [{cardValue: (J, Diamond), used: false}]
// 3- bid trump phase
// 		- ask each player in order (including bots) if want to "Pick it up".
// 				send to client -> {opcode: x, player: Z, timer: Y}
// 				reply from client -> {pass:boolean, trump:byte, goAlone: boolean}
// 				server sends -> {opcode:x, player: Dealer}
// 				client sends -> {opcode:x, card: ()}
// 			- if a trump is chosen move to `4- game phase`
// 			- else star 'second round'
// 		- (second round) ask each player in order to name a trump suit (cannot be the one turned down)
// 			send to client -> {opcode: x, player: Z, timer: Y}
// 			reply from client -> {pass:boolean, trump:byte, goAlone: boolean}
// 		- (third phase) “Stick the Dealer” ??
// 4- game phase (hand)
// 		- the 'maker' needs to choose if "going alone"
// 			- if yes then remove it's partner from the players in that game
// 		- play the tricks
// 			- each player in order plays a card, and at the end of the turn the highest played card is the winner,
// 			  a trick is awarded to the team that play that card
// 				server sends -> {opcode: x, player: Z, timer: Y, usableCards: [(),()]}
// 				client sends -> {opcode: x1, card: (A, Diamond)}
// 				re-broadcast -> {opcode: x1, card: (A, Diamond)} (to everyone)
// 		- after all the cards have been played the game is over, and the teams receive the score.
// 			- if a team reached X (normally 10) points the match is over
// 			- else another game starts
// 				server sends -> {opcode: x, player: Z, cards: []}
