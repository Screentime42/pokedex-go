package commands

import (
	"pokedex-go/internal/api"
	"fmt"
	"strings"
)

type BattleStats struct {
    HP          int
    Attack      int
    Defense     int
    SpecialAtk  int
    SpecialDef  int
    Speed       int
}


func commandInspect(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: inspect <pokemon-name>")
		return
	}
	name := strings.TrimSpace(args[0])
	key := strings.ToLower(name)


	insp, err := api.FetchPokemonResult(name)
	if err != nil {
		fmt.Println("Failed to fetch Pokemon:", err)
		return
	}

	if _, exists := Pokedex[key]; exists {
		fmt.Printf("Name: %s\n", insp.Name)
		fmt.Printf("Height: %v\n",insp.Height)
		fmt.Printf("Weight: %v\n", insp.Weight)

		stats := ExtractStats(insp.Stats)
		fmt.Println("Stats:")
		fmt.Printf("	-hp: %d\n", stats.HP)
		fmt.Printf("	-attack: %d\n", stats.Attack)
		fmt.Printf("	-defense: %d\n", stats.Defense)
		fmt.Printf("	-special-attack: %d\n", stats.SpecialAtk)
		fmt.Printf("	-special-defense: %d\n", stats.SpecialDef)
		fmt.Printf("	-speed: %d\n", stats.Speed)

		fmt.Println("Types:")
		for _, t := range insp.Types {
			fmt.Print("	-" + t.Type.Name + "\n")
		}
	} else {
		fmt.Printf("You havn't caught a %s ... yet!\n", name)
	}
}

func ExtractStats(wrappers []api.StatWrapper) BattleStats {
	stats := BattleStats{}
	for _, s := range wrappers {
		switch s.Stat.Name {
		case "hp":
			stats.HP = s.BaseStat
		case "attack":
			stats.Attack = s.BaseStat
		case "defense":
			stats.Defense = s.BaseStat
		case "special-attack":
			stats.SpecialAtk = s.BaseStat
		case "special-defense":
			stats.SpecialDef = s.BaseStat
		case "speed":
			stats.Speed = s.BaseStat
		}
	}
	return stats
}