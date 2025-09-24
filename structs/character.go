package main

import "fmt"

var spellMap = map[string]map[string]int{
	"Dragon's breath": {"cost": 50, "damage": 100},
}

type Character struct {
	Name   string
	Health int
	Mana   int
}

func (c *Character) Attack(target *Character) {
	baseDamage := 10
	target.Health = target.Health - baseDamage
	fmt.Printf("%s strikes with a sword and deals %d damage\n", c.Name, baseDamage)
}

func (c *Character) CastSpell(spellName string, target *Character) {
	spellDamage, spellCost := spellMap[spellName]["damage"], spellMap[spellName]["cost"]
	c.Mana = c.Mana - spellCost
	if target.Health-spellDamage < 0 {
		target.Health = 0
	} else {
		target.Health = target.Health - spellDamage
	}

	fmt.Printf("%s casts a spell and deals %d damage\n", c.Name, spellDamage)
}

func (c *Character) String() string {
	return fmt.Sprintf("Name: %s | Health: %d | Mana: %d", c.Name, c.Health, c.Mana)
}

func main() {
	dragon := &Character{
		Name:   "Smaug",
		Health: 100,
		Mana:   100,
	}

	hero := &Character{
		Name:   "Bilbo Beggins",
		Health: 50,
		Mana:   10,
	}

	fmt.Println(hero)
	fmt.Println(dragon)
	fmt.Println("------------------------------------------------")
	hero.Attack(dragon)
	dragon.CastSpell("Dragon's breath", hero)
	fmt.Println("------------------------------------------------")
	fmt.Println(hero)
	fmt.Println(dragon)
}
