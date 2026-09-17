# ⚔️ Simulateur de Bataille - Équipe RPG

Un petit simulateur de combat en Go où une équipe de héros affronte des vagues d'attaques ennemies jusqu'à leur survie ou leur anéantissement total.

## 📋 Description

Ce projet simule une bataille entre une équipe de 6 héros et un ennemi qui inflige des dégâts de zone (AOE) à chaque tour. Le programme calcule les statistiques de l'équipe, gère les attaques reçues, et affiche l'état du combat en temps réel jusqu'à la fin de la bataille.

## 🗂️ Structure du projet
exercices1/
├── main.go                    # Point d'entrée du programme
├── stats/
│   └── stats.go                # Structure Soldat + calculs de statistiques
├── etatequipe/
│   └── etat.go                  # Vérification de l'état de l'équipe (vivants/morts)
├── attaques/
│   └── attaqueEquipe.go        # Logique des attaques ennemies
└── go.mod

## 🚀 Fonctionnalités

- **Statistiques de l'équipe** :
  - Recherche du héros avec le plus de PV (tank)
  - Recherche du héros avec la plus grosse attaque (DPS)
  - Calcul de la vie moyenne de l'équipe
  - Comptage des héros "faibles" (moins de 800 PV)

- **Système de combat** :
  - L'utilisateur choisit le nombre d'attaques ennemies
  - Chaque attaque inflige des dégâts à tous les héros vivants
  - Les PV ne peuvent pas descendre en dessous de 0
  - Affichage de l'état de l'équipe après chaque attaque

- **Gestion de fin de combat** :
  - Détection du "Team Wipe" (tous les héros morts)
  - Vérification si l'équipe peut continuer le combat
  - Résumé final avec le nombre de survivants et de morts

## 🛠️ Prérequis

- [Go](https://go.dev/dl/) version 1.18 ou supérieure

## ▶️ Installation et exécution

1. Clonez le repository :
```bash
git clone <url-du-repo>
cd exercices1

Lancez le programme :

go run main.go

Suivez les instructions à l'écran :
Entrez le nombre d'attaques que l'ennemi va effectuer
Pour chaque attaque, entrez la valeur des dégâts infligés



📊 Exemple d'exécution
Le tank est Thrall avec 1500 PV
le plus gros dps de l'équipe est Jaina avec 450 str
la vie moyenne des heros est de 958.333333
les heros les moins stuff sont au nombre de 2

Nom: Arthas Vie: 1200 Attaque: 250
Nom: Kael Vie: 850 Attaque: 320
Nom: Thrall Vie: 1500 Attaque: 180
Nom: Sylvanas Vie: 700 Attaque: 400
Nom: Garrosh Vie: 1000 Attaque: 280
Nom: Jaina Vie: 500 Attaque: 450

combien d'attaques fait l'énnemi?
2
attaque de l'ennemie en Dgt : 300
...
6 personnages vivants

=== FIN DE LA BATAILLE ===
Nombre de Heros vivant : 6
Nombre de Heros mort : 0
l'équipe peut continuer le combat
🧩 Détail des packages
stats
Contient la structure Soldat (Nom, Vie, Attaque) ainsi que les fonctions d'analyse :

TrouverPlusDeVie(equipe []Soldat) Soldat
TrouverPlusDAttaque(equipe []Soldat) Soldat
CalculerVieMoyenne(equipe []Soldat) float64
CompterFaibles(equipe []Soldat) int

etatequipe
Gère l'état global de l'équipe :

CompterVivants(equipe []Soldat, index int) int — comptage récursif des survivants
PeutContinuer(equipe []Soldat) bool — détermine si le combat peut continuer

attaques
Applique les dégâts à l'équipe :

AttaquerEquipe(ptequipe *[]Soldat, degats int) — inflige des dégâts de zone à tous les héros vivants

🐛 Améliorations possibles

 Corriger la variable degats initialisée dans main.go mais écrasée par fmt.Scan dans AttaquerEquipe (le paramètre est inutile actuellement)
 Ajouter une riposte de l'équipe contre l'ennemi
 Ajouter des tests unitaires pour chaque package
 Gérer les erreurs de saisie utilisateur (fmt.Scan)
 Ajouter une interface plus visuelle (couleurs dans le terminal)

📄 Licence
Projet réalisé dans un cadre pédagogique — libre d'utilisation.# Evaluation-tableau-Christophe-Barbuscia