##### *`GET ALL JOBS`*
```
GET /api/jobs
Content-Type: application/json
X-API-Key: <API_KEY>

{
    "status": "OK",
    "message": "successfully getting all jobs",
    "data": [
        {
            "id": 1,
            "code": "1234",
            "name": "KULI"
        },
        {
            "id": 2,
            "code": "PK1",
            "name": "PEKERJAAN"
        },
        .....       
}        
```

##### *`CREATE ALL JOBS`*
```
POST /api/jobs
Content-Type: application/json
X-API-Key: <API_KEY>

{
    "kode" : "OE3",
    "name" : "BUDAK NIPPON"  
}

RESPONSE => 

{
    "status" : "CREATED",
    "message : "successfully create data jobs"
}

{
    "status": "ERROR",
    "message": "job with code PK143432 already exist"
}
        
```

##### *`UPDATE JOBS BY ID`*
```
PUT /api/jobs/{id}
Content-Type: application/json
X-API-Key: <API_KEY>

{
    "kode" : "OE3"
    "name" : "BUDAK NIPPON"  
}      

RESPONSE => 

{
    "status" : "OK",
    "message : "successfully update jobs id "
}

{
    "status": "ERROR",
    "message": "job dengan code 1234S sudah ada"
} 
```

##### *`DELETE JOBS BY ID`*
```
DELETE /api/jobs/{id}
Content-Type: application/json
X-API-Key: <API_KEY>

{
    "kode" : "OE3"
    "name" : "BUDAK NIPPON"  
}    

==> RESPONSE

{
    "status": "OK",
    "message": "deleted jobs id 33"
}

{
    "status": "ERROR",
    "message": "job dengan ID 33 tidak ditemukan"
}    
```

