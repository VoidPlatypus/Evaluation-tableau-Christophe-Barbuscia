package stats

type Soldat struct {
	Nom     string
	Vie     int
	Attaque int
}

func TrouverPlusDeVie(equipe []Soldat) Soldat {
	PVmax := equipe[0]
	for _, Soldat := range equipe {
		if Soldat.Vie > PVmax.Vie {
			PVmax = Soldat
		}
	}
	return PVmax
}

func TrouverPlusDAttaque(equipe []Soldat) Soldat {
	Strmax := equipe[0]
	for _, Soldat := range equipe {
		if Soldat.Attaque > Strmax.Attaque {
			Strmax = Soldat
		}
	}
	return Strmax
}

func CalculerVieMoyenne(equipe []Soldat) float64 {
	vietotal := 0

	for _, Soldat := range equipe {
		vietotal += Soldat.Vie
	}
	return float64(vietotal) / float64(len(equipe))
}

func CompterFaibles(equipe []Soldat) int {
	Soldatfaible := 0
	for _, Soldat := range equipe {
		if Soldat.Vie < 800 {
			Soldatfaible++
		}
	}
	return Soldatfaible
}
