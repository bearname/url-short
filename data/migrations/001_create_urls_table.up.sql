CREATE TABLE IF NOT EXISTS urls
(
    id            UUID                      NOT NULL UNIQUE PRIMARY KEY,
    original_url  VARCHAR(2048)             NOT NULL,
    creation_date DATE DEFAULT CURRENT_DATE NOT NULL,
    alias         VARCHAR(255) UNIQUE       NOT NULL
)
