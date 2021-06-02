@echo
migrate -database postgres://postgres:postgres@127.0.0.1:5432/urlshort?sslmode=disable -path data/migrations %1 1