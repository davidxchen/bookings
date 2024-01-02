# Bookings and reservations

This is the repository for my bookings and reservations project

- Built in Go ver. 1.21.5
- Uses the [chi router](https://github.com/go-chi/chi/v5)
- Uses the Alex Edwards' [SCS session management](https://github.com/alexedwards/scs/v2)
- Uses [nosurf](https://github.com/justinas/nosurf)


## PostgreSQL in Docker (Sample arm64 image)
```shell
docker run -d \
--name some-postgres \
-e POSTGRES_PASSWORD=mysecretpassword \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v /custom/mount:/var/lib/postgresql/data \
arm64v8/postgres
```

## [Soda](https://gobuffalo.io/documentation/database/pop/) for database migrations

## Unit test and Coverage
### for Windows (Powershell)
```shell
go test --coverprofile c.out
go tool cover -html c.out
```
### macOS or Linux distros
```shell
go test --coverprofile c.out && go tool cover -html c.out
```