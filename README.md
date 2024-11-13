# ponego

# Description

Ponego est un générateur de sites statiques écrit en Go, spécialement conçu pour les utilisateurs d'Obsidian. Comme les anciennes pierres qui préservaient la connaissance à travers les âges, Ponego transforme vos notes personnelles en pages web accessibles au monde entier. Simple, rapide et flexible, il respecte la structure de vos notes tout en leur donnant une nouvelle vie sur le web

# Objectifs :

Apprendre go et faire quelque chose qui pourrait m'etre utile plus tard.

## Parse le markdown pour en faire du html

- Utilisation de goldmark pour parser le markdown

## Créer un site web

- Utilisation de librairie de base de GO pour créer un site web

## Créer un service cloud pour hébérger les sites

- Aucune idée

# Structure des fichiers inspirés par Hugo

structure généré par claude.ai

```
ponego/                    # Racine du projet
├── go.mod                 # module github.com/yunko006/ponego
├── internal/
│   └── ponenotes/
│       └── ponenotes.go   # package ponenotes
└── tests/
    └── ponenotes_test.go  # package ponenotes_test
```

# Inspirations :

- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files
- https://incident.io/blog/safe-by-default
- https://dev.to/chasefleming/building-a-go-static-site-generator-using-elem-go-3fhh
