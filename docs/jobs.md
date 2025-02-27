### *`Login`*
```
POST /api/login
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
    "status" : "OK",
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
