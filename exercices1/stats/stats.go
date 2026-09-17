package stats

type Soldat struct {
	Nom     string
	Vie     int
	Attaque int
}

func TrouverPlusDeVie(equipe [6]Soldat) Soldat {
	PVmax := equipe[0]
	for _, soldat := range equipe {
		if soldat.Vie > PVmax.Vie {
			PVmax = soldat
		}
	}
	return PVmax
}

func TrouverPlusDAttaque(equipe [6]Soldat) Soldat {
	Strmax := equipe[0]
	for _, soldat := range equipe {
		if soldat.Attaque > Strmax.Attaque {
			Strmax = soldat
		}
	}
	return Strmax
}

func CalculerVieMoyenne(equipe [6]Soldat) float64 {
	vietotal := 0

	for _, soldat := range equipe {
		vietotal += soldat.Vie
	}
	return float64(vietotal) / float64(len(equipe))
}

func CompterFaibles(equipe [6]Soldat) int {
	soldatfaible := 0
	for _, soldat := range equipe {
		if soldat.Vie < 800 {
			soldatfaible++
		}
	}
	return soldatfaible
}
