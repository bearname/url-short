CREATE
OR REPLACE FUNCTION getNArrayS(el text[], count int) RETURNS text AS
$$
SELECT string_agg(el[random() * (array_length(el, 1) - 1) + 1], ' ')
FROM generate_series(1, count) g(i) $$
    VOLATILE
    LANGUAGE SQL;

WITH T(ray) AS (
    SELECT (string_to_array(pg_read_file('C:\Users\mikha\go\src\url-short\data\dump\urls.csv')::text, E '\n'))
)
INSERT
INTO urls (id, original_url, creation_date, alias)
SELECT uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
       getNArrayS(T.ray, 1),
       NOW() + (random() * (NOW() + '90 days' - NOW())) + '30 days',
    LEFT (md5(random()::text), 8)
FROM generate_series(1, 10000) s(i), T;

COPY (SELECT alias From urls) To 'C:/Users/mikha/go/src/url-short/data/dump/alias.csv' With CSV DELIMITER ',' HEADER;