package output

import (
	"fmt"

	"github.com/kagameow/monsterduel/logger"
	"github.com/kagameow/monsterduel/structs"
)

func PrintAndWriteToLog(s string) {
	fmt.Print(s)
	logger.WriteStringToLog(s)
}

func DisplayGameTurnInfo(game *structs.Game) {
	playerInfo := fmt.Sprintf("%v hp: %v\n", game.Player.Name, game.Player.Hp)
	PrintAndWriteToLog(playerInfo)
	monsterInfo := fmt.Sprintf("%v hp: %v\n", game.Monster.Name, game.Monster.Hp)
	PrintAndWriteToLog(monsterInfo)
}

func DisplayGameRoundInfo(game *structs.Game) {
	roundInfo := fmt.Sprintf("Round: %v\n", game.Round)
	PrintAndWriteToLog(roundInfo)
}
