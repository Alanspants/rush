package main

import "fmt"

type Game struct {
	Money  map[string]int
	Cards  map[*Card]int
	Player []*Player
}

type Card struct {
	MoneyRequire map[string]int // [color]number
	Color        string         // card-color
}

type Player struct {
	MoneyOwn  map[string]int // [color]number
	CardOwn   []*Card        // ownCard
	CardColor map[string]int // [color]number
}

func (game Game) moneyAvailable(player Player, card Card, moneyAmount int) bool {
	cardMoneyAmount := 0
	for color, num := range card.MoneyRequire {
		cardMoneyAmount += game.Money[color] * (num - player.CardColor[color])
	}
	// card所需要的money总数 > moneyAmount
	if cardMoneyAmount > moneyAmount {
		return false
	}

	// player拥有的卡片数量 < card需求的卡片数量
	for color, num := range card.MoneyRequire {
		if player.MoneyOwn[color] < num-player.CardColor[color] {
			return false
		}
	}

	return true
}

func (game Game) purchase(player Player, card *Card) {
	for color, num := range card.MoneyRequire {
		player.MoneyOwn[color] -= num - player.CardColor[card.Color]
	}

	player.CardOwn = append(player.CardOwn, card)
	player.CardColor[card.Color] += 1
	game.Cards[card] -= 1
}

func main() {
	card0 := &Card{
		MoneyRequire: map[string]int{
			"black": 5,
			"blue":  2,
			"green": 1,
			"white": 1,
		},
		Color: "blue",
	}

	card1 := &Card{
		MoneyRequire: map[string]int{
			"black": 3,
			"blue":  2,
			"green": 1,
			"white": 1,
		},
		Color: "green",
	}

	card2 := &Card{
		MoneyRequire: map[string]int{
			"black": 1,
			"blue":  1,
			"green": 1,
			"white": 1,
		},
		Color: "white",
	}

	player0 := &Player{
		MoneyOwn: map[string]int{
			"black": 6,
			"blue":  2,
			"green": 2,
			"white": 1,
		},
	}

	game := &Game{
		Money: map[string]int{
			"black": 2,
			"blue":  2,
			"green": 1,
			"white": 1,
		},
		Cards: map[*Card]int{
			card0: 1,
			card1: 1,
			card2: 2,
		},
		Player: []*Player{
			player0,
		},
	}

	fmt.Println(game.Cards)

	fmt.Println(game.moneyAvailable(*player0, *card0, 17))
	game.purchase(*player0, card0)
	fmt.Println(*player0)
	fmt.Println(game.Player[0])
	fmt.Println(game.Cards)
}
