# *`SubscriptionEndpoint`*

### *`Paket internet user yang sedang aktif`*
```
GET /api/plans/customers/{NIK}
Content-Type: application/json
Authorization: Bearer <your-access-token>

RESPONSE =>
{
    "status": "OK",
    "message": "Successfully created new network provider.",
    "data": [
        { 
            "id" : 9
            "code_packet_internet" : "PKT01",
            "packet_internet" : "ComboGacor",
            "description" : "Gege toto bang ini kencang sekali"
            "lifetime" : 080757898764,
            "payment_time" : 1785212654,
            "status" : "paid",
            "created_at" : 1452254321
        },
        ....
    
    ]
}
```

### *`Membeli create paket internet user beli (ADMIN YG MENAMBHAKAN)`*
```
nanti payment time otomatis di isi jika satus di isi paid
POST /api/plans/customers
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "NIK" : "",    
    "packet_internet_code" : "",    
    "status" : ""        
}
```

### *`Jike  ingin mengubah packet internet user (ADMIN)`*
```
PUT /api/plans/customers/{id-subs} 
Content-Type: application/json
Authorization: Bearer <your-access-token>
{
    "NIK" : "",    
    "packet_internet_code" : "",    
    "status" : ""  
}
```