# DOCUMENTATION API DI
```
./doc/
```
# REQUIREMENT
- Go version: 1.22.6
- .env (pastikan setting env)

## HOW TO RUN IN DOCKER
```bash
docker-compose up -d
```
```bash
docker exec -it mysql bash
```
```bash
SOURCE ./docker-entrypoint-initdb.d/backup.sql
```

## HOW TO RUN IN LOCAL
```bash
make run
```

## HOW TO RUN IN VM
build and generate binary
```bash
go build -o myapp ./cmd
```
```bash
ls -l myapp
```
run binary
```bash
./myapp
```
    

