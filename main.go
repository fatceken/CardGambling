package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	genre  string
	value  string
	isUsed bool
}

type Player struct {
	name string
	card *Card
}

func (c Card) IsSame(a Card) bool {
	return c.value == a.value && c.genre == a.genre
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.value, c.genre)
}

func (p Player) String() string {
	return fmt.Sprintf("%ss card is %s", p.name, p.card)
}

func main() {
	cards := make([]Card, 52)
	createCards(cards)
	shuffleCards(cards)

	dealCard := cards[getRandom()%52]
	fmt.Println("Deal card is ", dealCard)

	players := make([]Player, getRandom()%7+2)
	fmt.Println("Player count : ", len(players))

	for i := 0; i < len(players); i++ {
		players[i] = Player{name: fmt.Sprintf("player_%d", i)}
	}

	isFinished := false

	for !isFinished {
		for i := 0; i < len(players); i++ {
			players[i].card, _ = getCard(cards)
			fmt.Println(fmt.Sprintf("player_%ds card is : %s", i, players[i].card))
			if players[i].card.IsSame(dealCard) {
				fmt.Println(fmt.Sprintf("player_%d wins", i))
				isFinished = true
				break
			} else {
				time.Sleep(2 * time.Second)
				continue
			}
		}
	}

}

func createCards(cards []Card) {

	genres := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	specials := []string{"Jack", "Queen", "King", "Ace"}

	for j := 0; j < 52; j++ {
		if j%13 < 9 {
			cards[j] = Card{value: fmt.Sprint(j%13 + 2), genre: genres[j%4]}
		} else {
			cards[j] = Card{value: specials[j%13-9], genre: genres[j%4]}
		}
	}
}

func getCard(cards []Card) (*Card, error) {

	for i := 0; i < len(cards); i++ {
		if !cards[i].isUsed {
			cards[i].isUsed = true
			return &cards[i], nil
		}
	}

	return &Card{}, errors.New("empty")
}

func getRandom() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func shuffleCards(cards []Card) {
	fmt.Println("Shuffling cards...")

	for i := 0; i < 100000; i++ {
		func(card1, card2 *Card) {
			*card1, *card2 = *card2, *card1
		}(&cards[getRandom()%52], &cards[getRandom()%52])
	}
}