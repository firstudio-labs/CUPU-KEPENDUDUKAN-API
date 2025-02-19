
# *`AdminVillage`*


### *`ViewAll Citizens data by NIK`*
```
GET /api/citizens-data/{NIK}
Content-Type: application/json
Authorization: Bearer <your-access-token>

RESPONSE =>
{
    "NIK"         : "1234567890123456",
    "full_name"   : "John Doe",
    "province"    : "Jawa Barat",
    "district"    : "Bandung",
    "sub_district": "Kecamatan A",
    "village"     : "Desa B",
    "roles"       : "technician",
}

```

### *`ViewAll Citizens data with pagination per page 10 data`*
```
GET /api/citizens-data?page={page-number}
Content-Type: application/json
Authorization: Bearer <your-access-token>

RESPONSE =>
{
  "status": "success",
  "message": "Successfully retrieved citizens data.",
  "data": {
    "page": 1,
    "total_pages": 5,
    "total_records": 50,
    "citizens": [
      {
        "nik": "1234567890123456",
        "full_name": "John Doe",
        "village": "Village A"
      },
      {
        "nik": "2345678901234567",
        "full_name": "Jane Smith",
        "village": "Village B"
      },
      ...
    ]
  }
}

```

### *`add Citizens data`*

```
POST /api/citizens-data
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "NIK"         : "1234567890123456",
    "full_name"   : "John Doe",
    "province"    : "Jawa Barat",
    "district"    : "Bandung",
    "sub_district": "Kecamatan A",
    "village"     : "Desa B",
    "roles"       : "technician",
    "username"    : "johndoe",
    "password"    : "password123",
}

RESPONSE =>
{
    "status"  : "OK",
    "message" : "success added new citizen data",
}
```

### *`Update Citizens With By NIK`*
```
PATCH /api/citizens-data/{NIK}
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "full_name"   : "John Doe",
    "province"    : "Jawa Barat",
    "district"    : "Bandung",
    "sub_district": "Kecamatan A",
    "village"     : "Desa B",
    "roles"       : "technician",
    "username"    : "johndoe",
    "password"    : "password123",
}

RESPONSE =>
{
    "status"  : "OK",
    "message" : "success updated citizen data",
}

```

### *`Delete Citizens With By NIK (Actualy not real delete we do softdelete)`* 
```
DELETE /api/citizens-data/{NIK}
Content-Type: application/json
Authorization: Bearer <your-access-token>

RESPONSE =>
{
    "status"  : "OK",
    "message" : "delete citizen data success",
}

```


