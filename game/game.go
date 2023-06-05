package game

import (
	"fmt"
	"strings"
)

type Player interface {
	Move()
	Position() int
	Name() string
}

type Game struct {
	field_length int
	players      []Player
}

func NewGame(field_length int) *Game {
	return &Game{
		field_length: field_length,
		players:      make([]Player, 0),
	}
}

func (g *Game) Join(player Player) *Game {
	g.players = append(g.players, player)
	return g
}

func (g *Game) MovePlayers() {
	for _, player := range g.players {
		player.Move()
	}
}

func (g *Game) CheckWinner() Player {
	for _, player := range g.players {
		if player.Position() > g.field_length {
			return player
		}
	}
	return nil
}

func (g *Game) Print() {
	fmt.Println("|" + strings.Repeat("-", g.field_length+6) + "|")
	for _, player := range g.players {
		pos := player.Position()
		name := player.Name()
		repeat_count := g.field_length - pos + 1
		if repeat_count < 0 {
			repeat_count = 0
		}
		row := "|" + strings.Repeat(" ", pos) + name + strings.Repeat(" ", repeat_count) + "|"
		fmt.Println(row)
	}
	fmt.Println("|" + strings.Repeat("-", g.field_length+6) + "|")
}
