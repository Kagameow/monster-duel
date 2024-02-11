package structs

type Creature struct {
	Name      string
	Hp        int
	MinDamage int
	MaxDamage int
	Actions   map[string]Action
	IsPlayer  bool
}

type Game struct {
	Player     Creature
	Monster    Creature
	Round      int
	IsGameOver bool
}

type Action interface {
	GetName() string
	RunAction(*Creature, *Creature, *Game)
}
