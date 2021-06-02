# Docs

## Run with docker
### Requirements: docker

Run
```bat
docker composoe up
docker run -v {pathtoproject}/data/migrations:/migrations  --network host migrate/migrate  -path=/migrations/ -database postgresql://url-short:1234@localhost:5432/url-short?sslmode=disable {up | down} {number}
```

## Run local
### Requirements
golang 1.16, postgres:11

database postgres
Set environment variable SERVE_REST_ADDRESS,
DATABASE_ADDRESS,
DATABASE_NAME,
DATABASE_USER , 
DATABASE_PASSWORD

Run
```bat
make go-build
```