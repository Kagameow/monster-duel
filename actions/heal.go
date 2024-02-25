package actions

import (
	"fmt"
	"github.com/kagameow/monsterduel/output"
	"github.com/kagameow/monsterduel/structs"
)

type SimpleHeal struct {
	name       string
	healAmount int
}

func (heal *SimpleHeal) RunAction(healer *structs.Creature, _ *structs.Creature, _ *structs.Game) {
	healer.Hp += heal.healAmount
	if healer.Hp > healer.MaxHp {
		healer.Hp = healer.MaxHp
	}
	message := fmt.Sprintf("%v healed for %v hp!\n", healer.Name, heal.healAmount)
	output.PrintAndWriteToLog(message)
}

func (heal *SimpleHeal) GetName() string { return heal.name }

func (heal *SimpleHeal) GetCooldown() int { return 0 }

func (heal *SimpleHeal) CooldownTick() {}

var Heal = SimpleHeal{
	name:       "Heal wound",
	healAmount: 15,
}
