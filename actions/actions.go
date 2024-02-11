package actions

import (
	"fmt"
	"github.com/kagameow/monsterduel/output"
	"github.com/kagameow/monsterduel/structs"
	"github.com/kagameow/monsterduel/utils"
)

type SimpleAttack struct {
	name       string
	multiplier int
}

func (attack *SimpleAttack) RunAction(attacker *structs.Creature, opponent *structs.Creature, _ *structs.Game) {
	attackPower := utils.RandomRangeInt(attacker.MinDamage, attacker.MaxDamage) * attack.multiplier
	opponent.Hp -= attackPower
	message := fmt.Sprintf(
		"%v used %v against %v, and dealt %v damage!\n",
		attacker.Name, attack.name, opponent.Name, attackPower,
	)
	output.PrintAndWriteToLog(message)
}

func (attack *SimpleAttack) GetName() string {
	return attack.name
}

var Attack = SimpleAttack{
	name:       "Attack",
	multiplier: 1,
}

type AttackWithCooldown struct {
	name            string
	currentCooldown int
	initialCooldown int
	multiplier      int
}
