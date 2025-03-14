api run in port 3000
untuk import database nya di ./dump/backup.sql

# DOCUMENTATION API DI
```
./docs/
```

# REQUIREMENT
- Go version: 1.22.6 or more
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

if err use this

go build -o myapp -buildvcs=false ./cmd
```
```bash
ls -l myapp
```
run binary if want see log
```bash
./myapp
``` 
RUN IN BACKGORUND 
```bash
sudo nohup ./myapp &
```

PASTIKAN KILL KALO ADA YG MASIH AKTIF DI PORT 3000
```
lsof -i :3000
```
```
netstat -tulnp | grep :3000
```
```
sudo netstat -tulnp | grep :3000
```
```
sudo kill <PID>
```

JIKA DI DI APP ADA NGINX TAMBHAKAN INI BIAR HEADER TIDAK DI HAPUS AMAN DARI CORS,
```
 
location / {
      add_header 'Access-Control-Allow-Origin' '*' always;
      add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS' always;
      add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Api-Key' always;
      add_header 'Access-Control-Allow-Credentials' 'true' always;
      add_header 'Access-Control-Expose-Headers' 'Content-Length,Content-Range' always;
      
      # Penting: ini akan merespon OPTIONS langsung tanpa proxy ke aplikasi
      if ($request_method = 'OPTIONS') {
          add_header 'Access-Control-Allow-Origin' '*';
          add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
          add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Api-Key';
          add_header 'Access-Control-Max-Age' 1728000;
          add_header 'Content-Type' 'text/plain charset=UTF-8';
          add_header 'Content-Length' 0;
          return 204;
      }
  
      # Sisanya proxy ke aplikasi Anda
      proxy_hide_header Access-Control-Allow-Origin;
      proxy_hide_header Access-Control-Allow-Methods;
      proxy_hide_header Access-Control-Allow-Headers;
      proxy_pass http://localhost:3000;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
}

```