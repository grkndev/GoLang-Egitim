// go:build pointer_2

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Players []*Player
	isOver  bool
}
type Player struct {
	Name string
	HP   int
}

func NewGame() *Game {
	return &Game{Players: make([]*Player, 0), isOver: false}
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, HP: 100}
}

func Heal(p *Player, amount int) {
	p.HP += amount
	if p.HP > 100 {
		p.HP = 100
	}
	fmt.Printf("%s, %d canlılığını %d ile iyileştirdi \n", p.Name, p.HP, amount)
}

func Attack(g *Game, p *Player, target *Player, damage int) {
	if p == target {
		return
	}
	target.HP -= damage
	fmt.Printf("%s, %s'e %d hasar vurdu \n", p.Name, target.Name, damage)
	if target.HP < 0 {
		target.HP = 0
		g.isOver = true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	player1 := NewPlayer("Player 1")
	player2 := NewPlayer("Player 2")

	game.Players = append(game.Players, player1, player2)

	for !game.isOver {
		if rand.Intn(2) == 0 {
			if rand.Intn(2) == 0 {
				fmt.Println("1. oyuncu saldırıyor")
				Attack(game, player1, player2, rand.Intn(50)+1)
			} else {
				fmt.Println("2. oyuncu saldırıyor")
				Attack(game, player2, player1, rand.Intn(50)+1)
			}
		} else {
			if rand.Intn(2) == 0 {
				fmt.Println("1. oyuncu iyileşiyor")
				Heal(player1, rand.Intn(50)+1)
			} else {
				fmt.Println("2. oyuncu iyileşiyor")
				Heal(player2, rand.Intn(50)+1)
			}
		}
		fmt.Println("Player 1'nin HP: ", player1.HP)
		fmt.Println("Player 2'nin HP: ", player2.HP)
		fmt.Println("Game Over: ", game.isOver)
		fmt.Println("--------------------------------")
	}

	winner := player1
	if player2.HP > player1.HP {
		winner = player2
	}
	fmt.Printf("Oyun sona erdi. Kazanan: %s", winner.Name)
}
