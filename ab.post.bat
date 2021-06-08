@echo off
ab -p data\dump\short.json -T application/json  -c 10 -n 20000 http://localhost:8000/api/v1/urls