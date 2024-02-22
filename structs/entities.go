package structs

type Creature struct {
	Name      string
	Hp        int
	MaxHp     int
	MinDamage int
	MaxDamage int
	Actions   map[string]Action
}

type Game struct {
	Player     Creature
	Monster    Creature
	Round      int
	IsGameOver bool
}

type Action interface {
	GetName() string
	GetCooldown() int
	RunAction(*Creature, *Creature, *Game)
	CooldownTick()
}
