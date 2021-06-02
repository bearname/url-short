# Docs

## Run with docker

```bat
docker composoe up
docker run -v {pathtoproject}/data/migrations:/migrations  --network host migrate/migrate  -path=/migrations/ -database postgresql://url-short:1234@localhost:5432/url-short?sslmode=disable {up | down} {number}
```

## Run local
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