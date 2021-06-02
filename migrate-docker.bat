@echo off
docker run -v %1/data/migrations:/migrations  --network host migrate/migrate  -path=/migrations/ -database postgresql://url-short:1234@localhost:5432/url-short?sslmode=disable %2 %3