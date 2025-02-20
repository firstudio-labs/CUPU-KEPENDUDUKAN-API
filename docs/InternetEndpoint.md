# *`InternetEndpoint`*

### *`MENDAPATKAN  INFORMASI SUMBER JARINGAN DARI DAERAH BY CODE`*

```
GET /api/network/{id}
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "status": "OK",
    "message": "success get data network provide",
    "data": { 
        "id" : "1"
        "provider_name" : "XL"
        "source" : "GSM"
    }
}
```

### *`MENDAPATKAN SEMUA BY PAGE per page 10`*

```
GET /api/network?page={number}
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "status": "OK",
    "message": "success get data network provide",
    "data": { 
        "page": 1,
        "total_pages": 5,
        "total_records": 50,
        providers : [
            {
                "id" : "1"
                "provider_name" : "XL"
                "source" : "GSM"
            },
            {
                "id" : "2"
                "provider_name" : "PAK UDIN"
                "source" : "Space X"
            }
            ..
        ]
    }
}
```

### *`CreateNew`*
```
POST /api/network
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "provider_name": "Indosat",
    "source": "GSM"
}

RESPONSE =>
{
    "status": "OK",
    "message": "Successfully created new network provider.",
    "data": {
        "id": "3",
        "provider_name": "Indosat",
        "source": "GSM"
    }
}
```


### *`Update By id`*
```
PUT /api/network/{id}
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "provider_name": "Indosat Ooredoo",
    "source": "GSM"
}
```

### *`Delete By id`*
```
DELETE /api/network/{id}
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "provider_name": "Indosat Ooredoo",
    "source": "GSM"
}
```

### *`Get Paket internet by kode`*
```
POST /api/plans 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{


}
```


### *`GetAll paket by page, by provider juga`*
```
GET /api/plans?provider={name} 
GET /api/plans?page={id}
GET /api/plans 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{


}
```


### *`Create Paket internet by kode`*
```
POST /api/plans 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{


}
```

### *`Update Paket internet by kode`*
```
PUT /api/plans/{kode} 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{


}
```

### *`Delete Paket internet by kode`*
```
POST /api/plans/{kode} 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{


}
```