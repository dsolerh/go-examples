package main

import (
	"fmt"
	"math"
	"sort"
)

// Card represents a playing card
type Card struct {
	Suit  string // "H", "D", "S", "C"
	Value string // "9", "10", "J", "Q", "K", "A"
}

// String returns string representation of card
func (c Card) String() string {
	return c.Suit + c.Value
}

// EuchreRuleBot implements rule-based euchre strategy
type EuchreRuleBot struct {
	Position    int // 0=dealer, 1=left of dealer, 2=partner, 3=right of dealer
	TrumpSuit   string
	Hand        []Card
	CardsPlayed []Card
	TricksWon   [2]int // [team1, team2]
	Score       [2]int
}

// NewEuchreRuleBot creates a new bot instance
func NewEuchreRuleBot(position int) *EuchreRuleBot {
	return &EuchreRuleBot{
		Position:    position,
		Hand:        make([]Card, 0),
		CardsPlayed: make([]Card, 0),
		TricksWon:   [2]int{0, 0},
		Score:       [2]int{0, 0},
	}
}

// EvaluateHandStrength evaluates hand strength for trump calling decisions
// Returns a score from 0-10 indicating hand quality
func (bot *EuchreRuleBot) EvaluateHandStrength(hand []Card, trumpSuit string, upcard *Card) float64 {
	if trumpSuit == "" {
		return 0
	}

	strength := 0.0
	trumpCards := make([]Card, 0)
	offSuitAces := 0

	trumpHierarchy := bot.getTrumpHierarchy(trumpSuit)

	for _, card := range hand {
		if bot.isTrump(card, trumpSuit) {
			trumpCards = append(trumpCards, card)
			// Award points based on trump strength
			if bot.cardInSlice(card, trumpHierarchy[:2]) { // Right/Left bower
				strength += 3
			} else if len(trumpHierarchy) >= 5 && bot.cardInSlice(card, trumpHierarchy[:5]) { // Top 5 trump
				strength += 2
			} else {
				strength += 1
			}
		} else if card.Value == "A" { // Off-suit aces
			offSuitAces++
			strength += 1
		}
	}

	// Bonus for multiple trump
	if len(trumpCards) >= 3 {
		strength += 1
	}
	if len(trumpCards) >= 4 {
		strength += 1
	}

	// Bonus for off-suit aces
	strength += float64(offSuitAces) * 0.5

	// Include upcard if dealer
	if upcard != nil && bot.Position == 0 {
		if bot.isTrump(*upcard, trumpSuit) {
			strength += 1
		}
	}

	return math.Min(strength, 10)
}

// DecideOrderUp decides whether to order up the dealer in first round of bidding
func (bot *EuchreRuleBot) DecideOrderUp(upcard Card, hand []Card, teamScore, opponentScore int) bool {
	trumpSuit := upcard.Suit
	handStrength := bot.EvaluateHandStrength(hand, trumpSuit, &upcard)

	scorePressure := bot.getScorePressure(teamScore, opponentScore)

	// Position-based thresholds
	var threshold float64
	switch bot.Position {
	case 1: // Left of dealer (first to bid)
		threshold = 7.0 - scorePressure
	case 2: // Partner of dealer
		threshold = 6.0 - scorePressure
	case 3: // Right of dealer (last before dealer)
		threshold = 5.5 - scorePressure
	default: // Dealer
		threshold = 4.5 - scorePressure // Dealer gets upcard
	}

	return handStrength >= threshold
}

// DecideTrumpSuit chooses trump suit in second round of bidding
func (bot *EuchreRuleBot) DecideTrumpSuit(hand []Card, teamScore, opponentScore int, passedSuits []string) *string {
	suits := []string{"H", "D", "S", "C"}
	var bestSuit *string
	bestStrength := 0.0

	scorePressure := bot.getScorePressure(teamScore, opponentScore)
	threshold := 6.0 - scorePressure

	for _, suit := range suits {
		if !bot.stringInSlice(suit, passedSuits) {
			strength := bot.EvaluateHandStrength(hand, suit, nil)
			if strength > bestStrength {
				bestStrength = strength
				bestSuit = &suit
			}
		}
	}

	// Stick the dealer rule - dealer must call if everyone passes
	if bot.Position == 0 && bestSuit == nil {
		for _, suit := range suits {
			if !bot.stringInSlice(suit, passedSuits) {
				strength := bot.EvaluateHandStrength(hand, suit, nil)
				if strength > bestStrength {
					bestStrength = strength
					bestSuit = &suit
				}
			}
		}
	}

	if bestStrength >= threshold {
		return bestSuit
	}
	return nil
}

// DecideGoAlone decides whether to go alone (play without partner)
func (bot *EuchreRuleBot) DecideGoAlone(hand []Card, trumpSuit string, teamScore, opponentScore int) bool {
	handStrength := bot.EvaluateHandStrength(hand, trumpSuit, nil)
	sureTricks := bot.countSureTricks(hand, trumpSuit)

	// Score-based decisions
	if teamScore >= 7 { // Close to winning, be conservative
		return sureTricks >= 4
	} else if opponentScore >= 8 { // Desperate situation
		return handStrength >= 8 || sureTricks >= 3
	} else { // Normal play
		return sureTricks >= 4 || handStrength >= 9
	}
}

// ChooseDiscard chooses which card to discard when picking up as dealer
func (bot *EuchreRuleBot) ChooseDiscard(hand []Card, upcard Card, trumpSuit string) Card {
	newHand := make([]Card, len(hand))
	copy(newHand, hand)
	newHand = append(newHand, upcard)

	// Never discard trump
	nonTrump := make([]Card, 0)
	for _, card := range newHand {
		if !bot.isTrump(card, trumpSuit) {
			nonTrump = append(nonTrump, card)
		}
	}

	if len(nonTrump) == 0 {
		// All trump - discard lowest
		trumpHierarchy := bot.getTrumpHierarchy(trumpSuit)
		for i := len(trumpHierarchy) - 1; i >= 0; i-- {
			if bot.cardInSlice(trumpHierarchy[i], newHand) {
				return trumpHierarchy[i]
			}
		}
	}

	// Discard lowest non-trump, preferring cards that don't help
	type cardPriority struct {
		priority int
		card     Card
	}

	discardPriority := make([]cardPriority, 0)

	for _, card := range nonTrump {
		priority := 0

		// Low cards are better to discard
		switch card.Value {
		case "9", "10":
			priority += 3
		case "J":
			priority += 2
		case "Q", "K":
			priority += 1
		}

		// Singleton suits are worse to discard (lose flexibility)
		suitCount := 0
		for _, c := range newHand {
			if c.Suit == card.Suit && c != card {
				suitCount++
			}
		}
		if suitCount == 0 {
			priority -= 2
		}

		discardPriority = append(discardPriority, cardPriority{priority, card})
	}

	if len(discardPriority) > 0 {
		// Sort by priority (highest first)
		sort.Slice(discardPriority, func(i, j int) bool {
			return discardPriority[i].priority > discardPriority[j].priority
		})
		return discardPriority[0].card
	}

	return nonTrump[0]
}

// ChooseCardToPlay chooses which card to play during trick-taking phase
func (bot *EuchreRuleBot) ChooseCardToPlay(hand []Card, trickCards []Card, trumpSuit string, leaderPos int) Card {
	legalCards := bot.getLegalCards(hand, trickCards, trumpSuit)

	if len(trickCards) == 0 { // Leading the trick
		return bot.chooseLeadCard(legalCards, trumpSuit, hand)
	}

	ledSuit := bot.getEffectiveSuit(trickCards[0], trumpSuit)

	if bot.canFollowSuit(legalCards, ledSuit, trumpSuit) {
		return bot.playFollowingSuit(legalCards, trickCards, trumpSuit, ledSuit)
	} else {
		return bot.playOffSuit(legalCards, trickCards, trumpSuit)
	}
}

// Helper methods

func (bot *EuchreRuleBot) chooseLeadCard(hand []Card, trumpSuit string, fullHand []Card) Card {
	trumpCards := make([]Card, 0)
	for _, card := range hand {
		if bot.isTrump(card, trumpSuit) {
			trumpCards = append(trumpCards, card)
		}
	}

	// Lead trump if strong in trump
	if len(trumpCards) >= 3 {
		trumpHierarchy := bot.getTrumpHierarchy(trumpSuit)
		for _, card := range trumpHierarchy {
			if bot.cardInSlice(card, trumpCards) {
				return card
			}
		}
	}

	// Lead aces in off-suits
	for _, card := range hand {
		if card.Value == "A" && !bot.isTrump(card, trumpSuit) {
			return card
		}
	}

	// Lead lowest off-suit card
	offSuit := make([]Card, 0)
	for _, card := range hand {
		if !bot.isTrump(card, trumpSuit) {
			offSuit = append(offSuit, card)
		}
	}

	if len(offSuit) > 0 {
		sort.Slice(offSuit, func(i, j int) bool {
			return bot.getCardValue(offSuit[i]) < bot.getCardValue(offSuit[j])
		})
		return offSuit[0]
	}

	// Only trump left - lead lowest
	if len(trumpCards) > 0 {
		trumpHierarchy := bot.getTrumpHierarchy(trumpSuit)
		for i := len(trumpHierarchy) - 1; i >= 0; i-- {
			if bot.cardInSlice(trumpHierarchy[i], trumpCards) {
				return trumpHierarchy[i]
			}
		}
	}

	return hand[0] // Fallback
}

func (bot *EuchreRuleBot) playFollowingSuit(legalCards []Card, trickCards []Card, trumpSuit string, ledSuit string) Card {
	currentWinner := bot.getTrickWinner(trickCards, trumpSuit)
	winningCard := trickCards[currentWinner]

	// Try to win if possible
	higherCards := make([]Card, 0)
	for _, card := range legalCards {
		if bot.cardBeats(card, winningCard, trumpSuit, ledSuit) {
			higherCards = append(higherCards, card)
		}
	}

	if len(higherCards) > 0 {
		// Play lowest card that wins
		sort.Slice(higherCards, func(i, j int) bool {
			return bot.getTrumpRank(higherCards[i], trumpSuit) < bot.getTrumpRank(higherCards[j], trumpSuit)
		})
		return higherCards[0]
	} else {
		// Can't win - play lowest
		sort.Slice(legalCards, func(i, j int) bool {
			return bot.getTrumpRank(legalCards[i], trumpSuit) < bot.getTrumpRank(legalCards[j], trumpSuit)
		})
		return legalCards[0]
	}
}

func (bot *EuchreRuleBot) playOffSuit(legalCards []Card, trickCards []Card, trumpSuit string) Card {
	currentWinner := bot.getTrickWinner(trickCards, trumpSuit)
	winningCard := trickCards[currentWinner]

	// Check if partner is winning
	partnerWinning := (currentWinner % 2) == (bot.Position % 2)

	if partnerWinning {
		// Partner winning - discard lowest
		nonTrump := make([]Card, 0)
		for _, card := range legalCards {
			if !bot.isTrump(card, trumpSuit) {
				nonTrump = append(nonTrump, card)
			}
		}
		if len(nonTrump) > 0 {
			sort.Slice(nonTrump, func(i, j int) bool {
				return bot.getCardValue(nonTrump[i]) < bot.getCardValue(nonTrump[j])
			})
			return nonTrump[0]
		}
	}

	// Try to trump if beneficial
	trumpCards := make([]Card, 0)
	for _, card := range legalCards {
		if bot.isTrump(card, trumpSuit) {
			trumpCards = append(trumpCards, card)
		}
	}

	if len(trumpCards) > 0 && !bot.isTrump(winningCard, trumpSuit) {
		// Play lowest trump to win
		sort.Slice(trumpCards, func(i, j int) bool {
			return bot.getTrumpRank(trumpCards[i], trumpSuit) < bot.getTrumpRank(trumpCards[j], trumpSuit)
		})
		return trumpCards[0]
	}

	// Discard lowest non-trump
	nonTrump := make([]Card, 0)
	for _, card := range legalCards {
		if !bot.isTrump(card, trumpSuit) {
			nonTrump = append(nonTrump, card)
		}
	}

	if len(nonTrump) > 0 {
		sort.Slice(nonTrump, func(i, j int) bool {
			return bot.getCardValue(nonTrump[i]) < bot.getCardValue(nonTrump[j])
		})
		return nonTrump[0]
	}

	// Only trump available - play lowest
	sort.Slice(legalCards, func(i, j int) bool {
		return bot.getTrumpRank(legalCards[i], trumpSuit) < bot.getTrumpRank(legalCards[j], trumpSuit)
	})
	return legalCards[0]
}

func (bot *EuchreRuleBot) getTrumpHierarchy(trumpSuit string) []Card {
	switch trumpSuit {
	case "H":
		return []Card{
			{"H", "J"}, {"D", "J"}, {"H", "A"}, {"H", "K"},
			{"H", "Q"}, {"H", "10"}, {"H", "9"},
		}
	case "D":
		return []Card{
			{"D", "J"}, {"H", "J"}, {"D", "A"}, {"D", "K"},
			{"D", "Q"}, {"D", "10"}, {"D", "9"},
		}
	case "S":
		return []Card{
			{"S", "J"}, {"C", "J"}, {"S", "A"}, {"S", "K"},
			{"S", "Q"}, {"S", "10"}, {"S", "9"},
		}
	case "C":
		return []Card{
			{"C", "J"}, {"S", "J"}, {"C", "A"}, {"C", "K"},
			{"C", "Q"}, {"C", "10"}, {"C", "9"},
		}
	}
	return []Card{}
}

func (bot *EuchreRuleBot) isTrump(card Card, trumpSuit string) bool {
	if card.Suit == trumpSuit {
		return true
	}
	// Left bower check
	if (trumpSuit == "H" || trumpSuit == "D") && card.Suit == map[string]string{"H": "D", "D": "H"}[trumpSuit] && card.Value == "J" {
		return true
	}
	if (trumpSuit == "S" || trumpSuit == "C") && card.Suit == map[string]string{"S": "C", "C": "S"}[trumpSuit] && card.Value == "J" {
		return true
	}
	return false
}

func (bot *EuchreRuleBot) getTrumpRank(card Card, trumpSuit string) int {
	hierarchy := bot.getTrumpHierarchy(trumpSuit)
	for i, c := range hierarchy {
		if c == card {
			return i
		}
	}
	return 100 // Not trump
}

func (bot *EuchreRuleBot) getCardValue(card Card) int {
	values := map[string]int{
		"9": 1, "10": 2, "J": 3, "Q": 4, "K": 5, "A": 6,
	}
	if val, ok := values[card.Value]; ok {
		return val
	}
	return 0
}

func (bot *EuchreRuleBot) getScorePressure(teamScore, opponentScore int) float64 {
	if opponentScore >= 8 {
		return 2.0 // Desperate - take more risks
	} else if teamScore >= 8 {
		return -1.0 // Conservative - close to winning
	} else if opponentScore >= 6 {
		return 1.0 // Some pressure
	}
	return 0.0
}

func (bot *EuchreRuleBot) countSureTricks(hand []Card, trumpSuit string) int {
	hierarchy := bot.getTrumpHierarchy(trumpSuit)
	sureTricks := 0

	// Count top trump cards
	for i := 0; i < 3 && i < len(hierarchy); i++ {
		if bot.cardInSlice(hierarchy[i], hand) {
			sureTricks++
		}
	}

	// Count off-suit aces
	for _, card := range hand {
		if card.Value == "A" && !bot.isTrump(card, trumpSuit) {
			sureTricks++ // Simplified - treating aces as sure tricks
		}
	}

	return sureTricks
}

func (bot *EuchreRuleBot) getLegalCards(hand []Card, trickCards []Card, trumpSuit string) []Card {
	if len(trickCards) == 0 {
		return hand
	}

	ledSuit := bot.getEffectiveSuit(trickCards[0], trumpSuit)

	// Must follow suit if possible
	sameSuit := make([]Card, 0)
	for _, card := range hand {
		if bot.getEffectiveSuit(card, trumpSuit) == ledSuit {
			sameSuit = append(sameSuit, card)
		}
	}

	if len(sameSuit) > 0 {
		return sameSuit
	}
	return hand
}

func (bot *EuchreRuleBot) getEffectiveSuit(card Card, trumpSuit string) string {
	if bot.isTrump(card, trumpSuit) {
		return trumpSuit
	}
	return card.Suit
}

func (bot *EuchreRuleBot) canFollowSuit(legalCards []Card, ledSuit string, trumpSuit string) bool {
	for _, card := range legalCards {
		if bot.getEffectiveSuit(card, trumpSuit) == ledSuit {
			return true
		}
	}
	return false
}

func (bot *EuchreRuleBot) cardBeats(card1, card2 Card, trumpSuit, ledSuit string) bool {
	// Both trump
	if bot.isTrump(card1, trumpSuit) && bot.isTrump(card2, trumpSuit) {
		return bot.getTrumpRank(card1, trumpSuit) < bot.getTrumpRank(card2, trumpSuit)
	}

	// Only card1 is trump
	if bot.isTrump(card1, trumpSuit) {
		return true
	}

	// Only card2 is trump
	if bot.isTrump(card2, trumpSuit) {
		return false
	}

	// Neither trump - must be same suit to beat
	if card1.Suit != card2.Suit {
		return false
	}

	return bot.getCardValue(card1) > bot.getCardValue(card2)
}

func (bot *EuchreRuleBot) getTrickWinner(trickCards []Card, trumpSuit string) int {
	if len(trickCards) == 0 {
		return 0
	}

	ledSuit := bot.getEffectiveSuit(trickCards[0], trumpSuit)
	winnerIdx := 0
	winningCard := trickCards[0]

	for i := 1; i < len(trickCards); i++ {
		if bot.cardBeats(trickCards[i], winningCard, trumpSuit, ledSuit) {
			winnerIdx = i
			winningCard = trickCards[i]
		}
	}

	return winnerIdx
}

// Utility functions
func (bot *EuchreRuleBot) cardInSlice(card Card, slice []Card) bool {
	for _, c := range slice {
		if c == card {
			return true
		}
	}
	return false
}

func (bot *EuchreRuleBot) stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// Example usage
func main() {
	// Create a new bot instance
	bot := NewEuchreRuleBot(1) // Left of dealer

	// Example hand and game state
	hand := []Card{
		{"H", "J"}, {"S", "A"}, {"C", "K"}, {"D", "10"}, {"H", "9"},
	}
	upcard := Card{"H", "A"}
	teamScore := 3
	opponentScore := 5

	// Bidding decisions
	shouldOrder := bot.DecideOrderUp(upcard, hand, teamScore, opponentScore)
	fmt.Printf("Should order up %s: %t\n", upcard, shouldOrder)

	// If trump is hearts
	trumpSuit := "H"
	shouldGoAlone := bot.DecideGoAlone(hand, trumpSuit, teamScore, opponentScore)
	fmt.Printf("Should go alone: %t\n", shouldGoAlone)

	// Card play
	trickCards := []Card{{"S", "K"}} // Opponent led King of Spades
	cardToPlay := bot.ChooseCardToPlay(hand, trickCards, trumpSuit, 0)
	fmt.Printf("Card to play: %s\n", cardToPlay)

	// Hand strength evaluation
	strength := bot.EvaluateHandStrength(hand, trumpSuit, &upcard)
	fmt.Printf("Hand strength: %.1f/10\n", strength)
}
