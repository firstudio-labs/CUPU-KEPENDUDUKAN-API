### *`Login`*
```
POST /api/v1/login
Content-Type: application/json
REQUEST =>

{
    "username":"korium",
    "password":"admin123",
}
```

```
SUCCESS RESPONSE 
{
    "status" : "success",
    "message": "login success",
    "data"   : {
          "token_type"  : "Bearer",
          "expires_in"  : ""
    }
}

BAD REQUEST USER
{
    "status"  : "error",
    "message" : "invalid username and password",
}

INTERNAL SERVER ERROR
{
    "status"  : "error",
    "message" : "internal server error, please try agin later."
}  
```


### *`Register`*
```
POST /api/v1/login
Content-Type: application/json
REQUEST =>
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
  "created_at"  : 1677415078,
  "updated_at"  : 1677415078,
  "deleted_at"  : null
}


```
