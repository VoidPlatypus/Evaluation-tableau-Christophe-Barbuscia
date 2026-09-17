package main

import (
	"exercices1/attaques"
	"exercices1/etatequipe"
	"exercices1/stats"
	"fmt"
)

func main() {
	equipe := []stats.Soldat{
		{Nom: "Arthas", Vie: 1200, Attaque: 250},
		{Nom: "Kael", Vie: 850, Attaque: 320},
		{Nom: "Thrall", Vie: 1500, Attaque: 180},
		{Nom: "Sylvanas", Vie: 700, Attaque: 400},
		{Nom: "Garrosh", Vie: 1000, Attaque: 280},
		{Nom: "Jaina", Vie: 500, Attaque: 450},
	}

	tank := stats.TrouverPlusDeVie(equipe)
	fmt.Printf("Le tank est %s avec %d PV\n", tank.Nom, tank.Vie)

	dps := stats.TrouverPlusDAttaque(equipe)
	fmt.Printf("le plus gros dps de l'équipe est %s avec %d str\n", dps.Nom, dps.Attaque)

	vieMoyenne := stats.CalculerVieMoyenne(equipe)
	fmt.Printf("la vie moyenne des heros est de %f\n", vieMoyenne)

	unstuff := stats.CompterFaibles(equipe)
	fmt.Printf("les heros les moins stuff sont au nombre de %d\n", unstuff)

	afficherEquipe(equipe)

	degats := 0
	NombreAttaque := 0
	vivants := 0
	Morts := 0

	fmt.Println("combien d'attaques fait l'énnemi?")
	fmt.Scan(&NombreAttaque)
	for i := 0; i < NombreAttaque; i++ {
		attaques.AttaquerEquipe(&equipe, degats)

		afficherEquipe(equipe)

		vivants = etatequipe.CompterVivants(equipe, 0)
		Morts = 6 - vivants
		fmt.Printf("%d personnages vivants\n", vivants)

		if vivants == 0 {
			fmt.Println("Team wype, go reset!")
			break
		}
	}
	if etatequipe.PeutContinuer(equipe) == true {
		fmt.Println("=== FIN DE LA BATAILLE ===\n", "Nombre de Heros vivant :", vivants, "\n", "Nombre de Heros mort ::", Morts, "\n", "l'équipe peut continuer le combat")
	}
	if etatequipe.PeutContinuer(equipe) == false {
		fmt.Println("=== FIN DE LA BATAILLE ===\n", "Nombre de Heros vivant :", vivants, "\n", "Nombre de Heros mort :", Morts, "\n", " Raid wype, go reset")
	}

}

func afficherEquipe(equipe []stats.Soldat) {
	for i := 0; i < len(equipe); i++ {
		if equipe[i].Vie == 0 {
			fmt.Printf("Nom: %s Vie: KO, fallait pas rester dans les AOE   Attaque: %d\n", equipe[i].Nom, equipe[i].Attaque)
		} else {
			fmt.Printf("Nom: %s Vie: %d Attaque: %d\n", equipe[i].Nom, equipe[i].Vie, equipe[i].Attaque)
		}
	}
}
