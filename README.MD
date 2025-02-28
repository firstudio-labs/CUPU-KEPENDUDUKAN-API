# RUN IN PORT 3000
# DOCUMENTATION API DI
```
./doc/
```

# REQUIREMENT
- Go version: 1.22.6
- .env (pastikan setting env itu ada env example)

## HOW TO RUN IN DOCKER
```bash
git clone https://github.com/firstudio-labs/KEPENDUDUKAN-API
```
```bash
cd KEPENDUDUKAN-API
````
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
git clone https://github.com/firstudio-labs/KEPENDUDUKAN-API
```
```bash
cd KEPENDUDUKAN-API
````
```bash
make run
```

## HOW TO RUN IN VM
```bash
go version
```
```bash
git clone https://github.com/firstudio-labs/KEPENDUDUKAN-API
```
```bash
cd KEPENDUDUKAN-API
```
download all dependencies
```bash
go mod tidy
```
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

