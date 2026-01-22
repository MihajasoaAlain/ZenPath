# ZenPath

ZenPath est un outil en **Go** qui organise automatiquement vos fichiers  
(images, documents, vidéos, musique…) dans des dossiers dédiés.  
Rapide, simple et concurrent grâce aux **goroutines**.

## Prérequis

- Go 1.20+ (module mode activé)

## Installation / Build

Depuis la racine du dépôt :

```bash
cd ZenPath
go build -o zenpath ./cmd/zenpath
```

## Utilisation

```bash
./zenpath <dossier>
```

Exemple :

```bash
./zenpath /home/user/Downloads
```

Le programme parcourt le dossier fourni et déplace les fichiers dans des
sous-dossiers par catégorie (Images, Documents, Videos, Musique, etc.).

## Catégories (extensions)

Les catégories sont définies dans `ZenPath/internal/config/config.go`.
Tu peux modifier cette map pour ajouter/supprimer des extensions ou renommer
les dossiers.

## Développement

- Le scan et le déplacement se font en parallèle via des workers.
- Les erreurs de scan sont affichées en console.
- Le déplacement utilise `os.Rename` (déplacement rapide sur le même disque).

Astuce : teste d'abord dans un dossier de test pour éviter de déplacer des
fichiers importants.
