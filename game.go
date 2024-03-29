package main

import (
	"fmt"
	"github.com/kagameow/monsterduel/actions"
	"github.com/kagameow/monsterduel/input"
	"github.com/kagameow/monsterduel/output"
	"github.com/kagameow/monsterduel/structs"
)

func newGame() *structs.Game {
	return &structs.Game{
		Player: structs.Creature{
			Name:      "Player",
			Hp:        100,
			MaxHp:     100,
			MinDamage: 7,
			MaxDamage: 10,
			Actions: map[string]structs.Action{
				"1": &actions.Attack,
				"2": &actions.Heal,
				"3": &actions.StrongAttack,
			},
		},
		Monster: structs.Creature{
			Name:      "Monster",
			Hp:        100,
			MaxHp:     100,
			MinDamage: 8,
			MaxDamage: 12,
			Actions: map[string]structs.Action{
				"1": &actions.Bite,
			},
		},
	}
}

func handleCooldowns(game *structs.Game) {
	for _, action := range game.Player.Actions {
		action.CooldownTick()
	}
	for _, action := range game.Monster.Actions {
		action.CooldownTick()
	}
}

func startRound(game *structs.Game) {
	game.Round += 1
	handleCooldowns(game)
	output.DisplayGameRoundInfo(game)
}

func checkForWinner(game *structs.Game) (isGameOver bool) {
	if game.Player.Hp <= 0 {
		output.PrintAndWriteToLog(fmt.Sprintf("\n%v won!", game.Monster.Name))
		return true
	}
	if game.Monster.Hp <= 0 {
		output.PrintAndWriteToLog(fmt.Sprintf("\n%v won!", game.Player.Name))
		return true
	}
	return
}

func handlePlayerTurn(game *structs.Game) (isGameOver bool) {
	var allowedKeys []string
	for key := range game.Player.Actions {
		if game.Player.Actions[key].GetCooldown() > 0 {
			continue
		}
		allowedKeys = append(allowedKeys, key)
		fmt.Printf("%v: %v\n", key, game.Player.Actions[key].GetName())
	}
	actionKey := input.HandlePlayerInput(allowedKeys)
	game.Player.Actions[actionKey].RunAction(&game.Player, &game.Monster, game)
	output.DisplayGameTurnInfo(game)
	return checkForWinner(game)
}

func handleAITurn(game *structs.Game) (isGameOver bool) {
	game.Monster.Actions["1"].RunAction(&game.Monster, &game.Player, game)
	output.DisplayGameTurnInfo(game)
	return checkForWinner(game)
}

func main() {
	game := newGame()
	output.PrintAndWriteToLog("Welcome to Monster Duel!\n")
	for {
		startRound(game)
		game.IsGameOver = handlePlayerTurn(game)
		if game.IsGameOver {
			break
		}
		game.IsGameOver = handleAITurn(game)
		if game.IsGameOver {
			break
		}
	}
}
