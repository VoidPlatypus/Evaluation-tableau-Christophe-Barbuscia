package attaques

import (
	"exercices1/stats"
	"fmt"
)

// cette fonction me permet de gerer l'attaque
func AttaquerEquipe(ptequipe *[]stats.Soldat, degats int) {

	fmt.Print("attaque de l'ennemie en Dgt : ")
	fmt.Scan(&degats)

	for i := range *ptequipe {
		if (*ptequipe)[i].Vie == 0 {
			continue
		}
		(*ptequipe)[i].Vie -= degats
		if (*ptequipe)[i].Vie < 0 {
			(*ptequipe)[i].Vie = 0
		}
	}
}
