# API Golang
Le but de cette APIest de réaliser une simple API HTTP gérant une base de données Postgres contenant des artistes Spotify ainsi que leurs tracks.

# BDD

``` sql
  CREATE DOMAIN "spotify_id" AS character varying(22)
  COLLATE "default"
  CONSTRAINT check_spotify_id_format CHECK (VALUE::text ~* '^[0-9azA-Z]{22}$'::text);
  CREATE TABLE spotify_artist (
  id spotify_id NOT NULL,
  name varchar(255) NOT NULL,
  PRIMARY KEY (id)
  );
  CREATE TABLE spotify_track (
  id spotify_id NOT NULL,
  artist_id spotify_id NOT NULL,
  name varchar(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (artist_id) REFERENCES spotify_artist(id)
)
```

# Specs API

## Configuration

- L'API est configurée avec des variables d'environnement sur la base de [template.env](config/.env.template)

- Structure des dossiers
``` bash

├── app
│   ├── routes.go
│
├── error
│   └── apiError.go
│
├── handler
│   └── artistHandlers.go
│
├── middleware
│   └── apiKeyMiddleware.go
│
├── config
│   └── .env.template
│
├── domain
│   ├── artist.go
│
├── .gitignore
├── ca.pem
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
├── .env
```


## Sécurité
Toutes les routes devront être protégées par un middleware checkant l'api key présent dans le request header Api-Key .

La valeur de cette api key doit être hashée en sha256 et ce hash présent dans la variable d'env API_KEY_SHA256 .

## Routes
L'api implemente 4 routes:

* `POST /artist`
A pour but d'inscrire un nouvel artist en DB. Elle accepte comme body JSON:
``` json
{ "id": "6olE6TJLqED3rqDCT0FyPh", "name": "Nirvana" }
```

* `POST /artist/:id/track`
A pour but d'inscrire une nouvelle track en DB. Elle accepte comme body JSON:
``` json
{ "id": "5muVpPu8Fj9fXfDbbqDdrZ", "name": "Bleach" }
```

* `GET /artists/`

Retournes la liste de tout les artiste en db

Format de réponses:

``` json
[
  {
    "id": "6olE6TJLqED3rqDCT0FyPh",
    "name": "Nirvana",
    "tracks": [
      {
      "id": "5muVpPu8Fj9fXfDbbqDdrZ",
      "name": "Bleach"
      }, 
    ...]
}, ...]
```

* `POST /artist/url`

Qui fait pareil que POST /artist/ mais qui prend en body
``` json 
{ 
  "spotify_url": "https://open.spotify.com/artist 6olE6TJLqED3rqDCT0FyPh?si=3fe863c0438a4593"
 } 
```
et va fetch le name directement depuis l'url.

