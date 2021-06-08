@echo off
curl -sK -v http://localhost:8000/debug/pprof/profile/heap > %1.out
go tool pprof %1.out