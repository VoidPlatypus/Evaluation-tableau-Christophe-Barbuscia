package main

import "fmt"

type Soldat struct {
	nom     string
	vie     int
	attaque int
}

func main() {

	equipe := [6]Soldat{
		{"Arthas", 1200, 250},
		{"Kael", 850, 320},
		{"Thrall", 1500, 180},
		{"Sylvanas", 700, 400},
		{"Garrosh", 1000, 280},
		{"Jaina", 500, 450},
	}
	afficherEquipe(equipe)

	tank := trouverPlusDeVie(equipe)
	fmt.Printf("Le tank est %s avec %d PV\n", tank.nom, tank.vie)

	dps := trouverPlusDAttaque(equipe)
	fmt.Printf("le plus gros dps de l'équipe est %s avec %d str\n", dps.nom, dps.attaque)

	vieMoyenne := calculerVieMoyenne(equipe)
	fmt.Printf("la vie moyenne des heros est de %f\n", vieMoyenne)

	unstuff := compterFaibles(equipe)
	fmt.Printf("les heros les moins stuff sont au nombre de %d\n", unstuff)


	degats := 0

    fmt.Print("attaque de l'ennemie en Dgt : ")
    fmt.Scan(&degats)  

    attaquerEquipe(&equipe, degats)
    afficherEquipe(equipe)
}

func afficherEquipe(equipe [6]Soldat) {
    for i := 0; i < len(equipe); i++ {
        if equipe[i].vie == 0 {
            fmt.Printf("Nom: %s Vie: KO, fallait pas rester dans les AOE   Attaque: %d\n", equipe[i].nom, equipe[i].attaque)
        } else {
            fmt.Printf("Nom: %s Vie: %d Attaque: %d\n", equipe[i].nom, equipe[i].vie, equipe[i].attaque)
        }
	}
}
func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	PVmax := equipe[0]
	for _, soldat := range equipe {
		if soldat.vie > PVmax.vie {
			PVmax = soldat
		}
	}
	return PVmax
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	Strmax := equipe[0]
	for _, soldat := range equipe {
		if soldat.attaque > Strmax.attaque {
			Strmax = soldat
		}
	}
	return Strmax
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	vietotal := 0

	for _, soldat := range equipe {
		vietotal += soldat.vie
	}
	return float64(vietotal) / float64(len(equipe))
}

func compterFaibles(equipe [6]Soldat) int {
	soldatfaible := 0
	for _, soldat := range equipe {
		if soldat.vie < 800 {
			soldatfaible++
		}
	}
	return soldatfaible
}

// cette fonction me permet de gerer l'attaque
func attaquerEquipe(equipe *[6]Soldat, degats int) {
    for i := range equipe {
        if equipe[i].vie == 0 {
            continue 
        }
        equipe[i].vie -= degats
        if equipe[i].vie < 0 {
            equipe[i].vie = 0 
        }
    }
}
