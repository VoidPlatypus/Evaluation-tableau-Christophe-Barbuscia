package attaques

import (
	"exercices1/stats"
	"fmt"
)

// cette fonction me permet de gerer l'attaque
func AttaquerEquipe(equipe *[6]stats.Soldat, degats int) {

	fmt.Print("attaque de l'ennemie en Dgt : ")
	fmt.Scan(&degats)

	for i := range equipe {
		if equipe[i].Vie == 0 {
			continue
		}
		equipe[i].Vie -= degats
		if equipe[i].Vie < 0 {
			equipe[i].Vie = 0
		}
	}
}
