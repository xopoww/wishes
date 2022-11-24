# wishes

Wishes is a wishlist sharing app created as a pet project. Currently the project consists of the backend REST API server. This project is in WIP status, see [TODO](#todo) for details on planned features.

## Features

- Add wishlists consisting of an arbitrary amount of items - your wishes
- See other's wishlists and mark some items as "taken" by you, which means that you're going to grant that wish and others should choose something else
- List owner cannot see who marked which items on their list
- Three list access options: private (for drafts), public or available via access token provided by list owner (which can be embedded in sharing link in a web app)

## Some implementation details

The server's OpenAPI 2.0 spec can be found [here](/api/wishes.yml), the server code is generated via [go-swagger](https://github.com/go-swagger/go-swagger). The code is loosely organized into three layers, following the clean architecture approach:
 - Entity ([models package](/internal/models))
 - Usecases ([service package](/internal/service/))
 - Adapters ([sqlite package](/internal/repository/sqlite/) with a repository implementation and [handlers package](/internal/controllers/handlers/) defining HTTP handlers for go-swagger).

 The lists model supports optimistic versioning: each list has a revision id, which is included in GET response. All modifications to list's items increment revision id. These requests, as well as "taking" or "untaking" an item, require revision id to be passed with the request. If request's revision id is lower than the one on the server, a 409 HTTP code is returned meaning that the client should make another GET request before submitting the initial change.

 Apart from unit tests written in Go, this project also has e2e tests assesing REST API as a whole. These tests are written in python using pytest and can be found [here](/test/api_test/).

 ## Building and running

The server can be built via `make build`. The resulting application needs some environment variables to run:
 - `WISHES_JWT_SECRET` - base64-encoded secret used to sign user API tokens
 - `WISHES_LIST_SECRET` - base64-encoded secret used to sign list access tokens
 - `WISHES_DBS` - path to sqlite3 database file
 - `PORT` (optional) - port for the server to listen to

 To create a new, empty database with required migrations applied, you must install [migrate utility](https://github.com/golang-migrate/migrate/) and make sure it is awailable in PATH as `migrate` and run
 ```bash
 bash scripts/db/create.sh <path to dst file>
 ```

 Alternatively, you can apply these migrations yourself using this (or any other) utility. The migration files can be found [here](/internal/repository/sqlite/migrations/).


 To run e2e tests, you must navigate to `/test/api_test` directory, install python dependencies from `requiremets.txt` and run 
 ```bash
 WISHES_HOST=<wishes server base path> pytest .
 ```
 `WISHES_HOST` environment variable is a full base path to the API server. E.g. if the server is run locally on port 65000, `WISHES_HOST` must be set to `http://localhost:65000/api` (notice the `/api` part).

 All of this can also be achieved via Docker compose: create `./dev/dev.envrc` file with `WISHES_JWT_SECRET` and `WISHES_LIST_SECRET` variables and run
 ```
 docker compose --profile test build
 docker compose up -d
 docker compose run api_test .
 ```

 ## TODO

 - switch from SQLite to another DBMS
 - create web SPA for the app and integrate it with API
 - ~~add OAuth2 authentication~~
 - implement SSE to notify front-end about changes to opened lists and other events