package etatequipe

import "exercices1/stats"

func CompterVivants(equipe [6]stats.Soldat, index int) int {

	if index == 6 {
		return 0
	}

	if equipe[index].Vie > 0 {
		return 1 + CompterVivants(equipe, index+1)
	}

	return CompterVivants(equipe, index+1)
}


// func peutContinuer(equipe [6]Soldat) bool