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
ponego/
├── commands/ # Commandes CLI (similaire à Hugo)
│ ├── build.go # Commande 'ponego build'
│ ├── serve.go # Commande 'ponego serve'
│ ├── new.go # Commande 'ponego new'
│ └── root.go # Configuration cobra de base
├── config/ # Gestion de la configuration
│ ├── config.go # Configuration de base
│ └── default.go # Valeurs par défaut
├── parser/ # Parsing Markdown/Obsidian
│ ├── markdown.go # Parser Markdown standard
│ ├── obsidian.go # Extensions Obsidian spécifiques
│ └── frontmatter.go # Gestion des métadonnées
├── output/ # Génération des sorties
│ ├── html/ # Génération HTML
│ ├── assets/ # Gestion des assets
│ └── sitemap.go # Génération sitemap
├── hugolib/ # (dans notre cas 'ponegolib')
│ ├── site.go # Logique principale du site
│ ├── page.go # Structure des pages
│ └── content.go # Gestion du contenu
├── common/ # Utilitaires communs
│ ├── paths.go # Gestion des chemins
│ └── urls.go # Manipulation URLs
├── deps/ # Gestion des dépendances
│ └── deps.go # Injection de dépendances
├── helpers/ # Fonctions helpers
├── resources/ # Gestion des ressources
│ ├── page/ # Templates de pages
│ └── themes/ # Gestion des thèmes
├── cmd/ponego/
│ └── main.go # Point d'entrée
├── tpl/ # Templates
├── layouts/ # Layouts par défaut
└── assets/ # Assets par défaut
```

# Inspirations :

- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files
- https://incident.io/blog/safe-by-default
- https://dev.to/chasefleming/building-a-go-static-site-generator-using-elem-go-3fhh
