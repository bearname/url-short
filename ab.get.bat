@echo off
ab -c 10 -n 20000 http://localhost:8000/api/v1/urls/%1