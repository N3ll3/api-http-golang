# Test Technique Backend Golang
Le but de ce test est de réaliser une simple API HTTP gérant une base de données
Postgres contenant des artistes Spotify ainsi que leurs tracks.
Pour ce faire une DB créée avec ce schema t'es mis à disposition:
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

Tu as comme droit `SELECT`,`INSERT`,`UPDATE`,`DELETE` sur ces deux tables.

# Specs API

## Installation et configuration



## Routes

L'api devra implementer ces 3 routes:

* `POST /artist`
A pour but d'inscrire un nouvel artist en DB. Elle accepte comme body JSON:
``` json
{ "id": "6olE6TJLqED3rqDCT0FyPh", "name": "Nirvana" }
```
- Middleware pour API-KEY
- Action 
  - struct Artist
  - controles
    - length Name
    - format id ?
- Response

- ArtistRepository => INSERT en base

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

Les routes en `POST` doivent répondre avec un status 201 en cas de succès, celle en `GET` un 200.

Les cas d'erreurs doivent être gérés de façon standard, e.g. 400 pour un mauvais format de requête, 404 pour les ressources inexistantes, 422 en cas de conflit, 500 autres.

Pas besoin de retourner un body juste le statut suffira en cas d'erreur.

Toutes les routes devront être protégées par un middleware checkant l'api key présent dans le request header Api-Key .

La valeur de cette api key doit être hashée en sha256 et ce hash présent dans la variable d'env API_KEY_SHA256 .

# Bonus
Créer une route
`POST /artist/url`

Qui fait pareil que POST /artist/ mais qui prend en body
``` json 
{ 
  "spotify_url": "https://open.spotify.com/artist 6olE6TJLqED3rqDCT0FyPh?si=3fe863c0438a4593"
 } 
```

et va fetch le name directement depuis l'url.
NB: Pour rappel, les ids Spotify sont présent dans le path des urls en dernière position
quand on share une url.
Dans tout les cas je ne je vérifierai pas l'exactitude de l'id mais seulement son bon
format.
