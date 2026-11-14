# 🐾 Dog Tracks

A dog training API built in Go for keeping track of dogs and their training sessions.

I'm building this project to practice API design, testing, and keeping HTTP handlers, business logic, and data access separate. The longer term goal is to connect it to a mobile app.

## Current status

Work in progress. The API currently uses seeded, in-memory repositories. No database setup is required, and changes are lost when the server restarts.

Current functionality includes:

- Create, retrieve, update, and soft delete dogs.
- Create, retrieve, update, and soft delete training sessions.
- Retrieve training sessions for a dog.
- Include deleted records when retrieving dogs or sessions.
- Handler and service tests using fake repositories.

## Run locally

Use Go 1.27.2 or newer, as specified in `go.mod`.

```sh
git clone https://github.com/Mortr0n/dog_tracks.git
cd dog_tracks
go run .
```

The server listens at **http://localhost:8000**.

Try these URLs in a browser or Postman:

- Get a seeded dog: http://localhost:8000/api/dog/1/false
- Get that dog's training sessions: http://localhost:8000/api/trainingsessions/1/false

## Run tests

From the project root:

```sh
go test ./...
```

## API routes

| Method | Route | Purpose |
| --- | --- | --- |
| GET | `/api/dog/{id}/{deleted}` | Get a dog |
| POST | `/api/dog` | Create a dog |
| PATCH | `/api/dog/{id}` | Update a dog |
| DELETE | `/api/dog/{id}` | Soft delete a dog |
| GET | `/api/trainingsession/{id}/{deleted}` | Get a training session |
| GET | `/api/trainingsessions/{id}/{deleted}` | Get sessions for a dog |
| POST | `/api/trainingsession` | Create a training session |
| PATCH | `/api/trainingsession/{id}` | Update a training session |
| DELETE | `/api/trainingsession/{id}` | Soft delete a training session |

For the plural `/api/trainingsessions` route, `{id}` is the **dog ID**. For the singular route, it is the **session ID**.

Set `{deleted}` to `false` to exclude deleted records or `true` to include them.

## Project structure

| Folder | Responsibility |
| --- | --- |
| `route/` | Register HTTP routes |
| `handler/` | Read requests and write responses |
| `service/` | Business logic and repository interfaces |
| `repository/` | In-memory fake repositories and seed data |
| `object/` | Dog and training models |
| `helper/` | Shared response, error, and test helpers |

## Planned

- PostgreSQL persistence.
- A mobile client for tracking training sessions.
