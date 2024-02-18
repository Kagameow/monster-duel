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

func (attack *SimpleAttack) GetCooldown() int {
	return 0
}

func (attack *SimpleAttack) CooldownTick() {}

var Attack = SimpleAttack{
	name:       "Attack",
	multiplier: 1,
}

var Bite = SimpleAttack{
	name:       "Bite",
	multiplier: 1,
}

type AttackWithCooldown struct {
	name            string
	currentCooldown int
	initialCooldown int
	multiplier      int
}

func (attack *AttackWithCooldown) RunAction(attacker *structs.Creature, opponent *structs.Creature, _ *structs.Game) {
	if attack.currentCooldown > 0 {
		message := fmt.Sprintf("%v is on cooldown!\n", attack.name)
		output.PrintAndWriteToLog(message)
		return
	}
	attackPower := utils.RandomRangeInt(attacker.MinDamage, attacker.MaxDamage) * attack.multiplier
	opponent.Hp -= attackPower
	message := fmt.Sprintf(
		"%v used %v against %v, and dealt %v damage!\n",
		attacker.Name, attack.name, opponent.Name, attackPower,
	)
	output.PrintAndWriteToLog(message)
	attack.currentCooldown = attack.initialCooldown
}

func (attack *AttackWithCooldown) GetName() string {
	return attack.name
}

func (attack *AttackWithCooldown) GetCooldown() int {
	return attack.currentCooldown
}

func (attack *AttackWithCooldown) CooldownTick() {
	if attack.currentCooldown > 0 {
		attack.currentCooldown -= 1
	}
}

var StrongAttack = AttackWithCooldown{
	name:            "Strong Attack",
	currentCooldown: 3,
	initialCooldown: 3,
	multiplier:      2,
}
