package main

import (
	"exercices1/stats"
	"fmt"
)

func main() {
	equipe := [6]stats.Soldat{
        {Nom: "Arthas", Vie: 1200, Attaque: 250},
        {Nom: "Kael", Vie: 850, Attaque: 320},
        {Nom: "Thrall", Vie: 1500, Attaque: 180},
        {Nom: "Sylvanas", Vie: 700, Attaque: 400},
        {Nom: "Garrosh", Vie: 1000, Attaque: 280},
        {Nom: "Jaina", Vie: 500, Attaque: 450},
    }

	tank := stats.TrouverPlusDeVie(equipe)
	fmt.Printf("Le tank est %s avec %d PV\n", tank, tank)

	dps := stats.TrouverPlusDAttaque(equipe)
	fmt.Printf("le plus gros dps de l'équipe est %s avec %d str\n", dps.Nom, dps.Attaque)

	vieMoyenne := stats.CalculerVieMoyenne(equipe)
	fmt.Printf("la vie moyenne des heros est de %f\n", vieMoyenne)

	unstuff := stats.CompterFaibles(equipe)
	fmt.Printf("les heros les moins stuff sont au nombre de %d\n", unstuff)
	
	afficherEquipe(equipe)


	degats := 0
	NombreAttaque := 0

	fmt.Println("combien d'attaques fait l'énnemi?")
	fmt.Scan(&NombreAttaque)
	for i := 0; i < NombreAttaque; i++ {
		attaquerEquipe(&equipe, degats)
		afficherEquipe(equipe)

	}

}

func afficherEquipe(equipe [6]stats.Soldat) {
	for i := 0; i < len(equipe); i++ {
		if equipe[i].Vie == 0 {
			fmt.Printf("Nom: %s Vie: KO, fallait pas rester dans les AOE   Attaque: %d\n", equipe[i].Nom, equipe[i].Attaque)
		} else {
			fmt.Printf("Nom: %s Vie: %d Attaque: %d\n", equipe[i].Nom, equipe[i].Vie, equipe[i].Attaque)
		}
	}
}
