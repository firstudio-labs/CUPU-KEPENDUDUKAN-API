##### *`GET ALL CITIZEN`*
```
GET /api/citizens?page={page}
Content-Type: application/json
X-API-Key: <API_KEY>

RESPONSE ==>
{
    "status": "OK",
    "message": "successfully get 10/page",
    "data": {
        "pagination": {
            "current_page": 1,
            "total_page": 249,
            "total_items": 2487,
            "items_per_page": 10,
            "next_page": 2,
            "prev_page": 1
        },
        "citizens": [
            {
                "id": 764,
                "nik": 3323092110770003,
                "kk": 3323092402072308,
                "full_name": "SURATNO",
                "gender": "Laki-Laki",
                "birth_date": "22/10/1977",
                "age": 43,
                "birth_place": "TEMANGGUNG",
                "address": "MENTO BAWANG",
                "province_id": 13,
                "district_id": 210,
                "sub_district_id": 3048,
                "village_id": 38269,
                "rt": "001",
                "rw": "001",
                "postal_code": 0,
                "citizen_status": "WNI",
                "birth_certificate": "Ada",
                "birth_certificate_no": "8118/DIS/2000",
                "blood_type": "A",
                "religion": "Islam",
                "marital_status": "Kawin Tercatat",
                "marital_certificate": "Ada",
                "marital_certificate_no": "261/20/XI/2006",
                "marriage_date": "",
                "divorce_certificate": "Tidak Ada",
                "divorce_certificate_no": "",
                "divorce_certificate_date": "",
                "family_status": "KEPALA KELUARGA",
                "mental_disorders": "Tidak Ada",
                "disabilities": "",
                "education_status": "SLTP/SMP/Sederajat",
                "job_type_id": 6,
                "nik_mother": "",
                "mother": "DJUMARIYAH",
                "nik_father": "",
                "father": "KOMARI",
                "coordinate": ""
            },
            ....
    }               
}
```


##### *`GET CITIZEN BY NIK`*
```
GET /api/citizens/{nik}
Content-Type: application/json
X-API-Key: <API_KEY>

RESPONSE ==>
{
    "status": "OK",
    "message": "successfully get",
    "data": {
        "id": 765,
        "nik": 3323124211860002,
        "kk": 3323092402072308,
        "full_name": "IIS MUNARTI",
        "gender": "Perempuan",
        "birth_date": "02/11/1986",
        "age": 34,
        "birth_place": "TEMANGGUNG",
        "address": "MENTO BAWANG",
        "province_id": 13,
        "district_id": 210,
        "sub_district_id": 3048,
        "village_id": 38269,
        "rt": "001",
        "rw": "001",
        "postal_code": 0,
        "citizen_status": "WNI",
        "birth_certificate": "Ada",
        "birth_certificate_no": "3323-LT-03102014-0102",
        "blood_type": "O",
        "religion": "Islam",
        "marital_status": "Kawin Tercatat",
        "marital_certificate": "Tidak Ada",
        "marital_certificate_no": "",
        "marriage_date": "",
        "divorce_certificate": "Tidak Ada",
        "divorce_certificate_no": "",
        "divorce_certificate_date": "",
        "family_status": "ISTRI",
        "mental_disorders": "Tidak Ada",
        "disabilities": "",
        "education_status": "SLTP/SMP/Sederajat",
        "job_type_id": 3,
        "nik_mother": "",
        "mother": "SURATI",
        "nik_father": "",
        "father": "MUPARDI",
        "coordinate": ""
    }
}

{
    "status": "ERROR",
    "message": "Nik is not suitable"
}

```

##### *`CREATE CITIZEN`*
```
POST /api/citizens
Content-Type: application/json
X-API-Key: <API_KEY>

{
  "nik": 1234567890123456,
  "kk": 9876543210987654,
  "full_name": "John Doe",
  "gender": 1,
  "birth_date": "1990-01-01",
  "age": 35,
  "birth_place": "Jakarta",
  "address": "Jl. Merdeka No. 10, RT 01 RW 02",
  "province_id": 1,
  "district_id": 2,
  "sub_district_id": 3,
  "village_id": 4,
  "rt": "01",
  "rw": "02",
  "postal_code": 12345,
  "citizen_status": 1,
  "birth_certificate": 1,
  "birth_certificate_no": "BC1234567890",
  "blood_type": 1,
  "religion": 1,
  "marital_status": 1,
  "marital_certificate": 1,
  "marital_certificate_no": "MC1234567890",
  "marriage_date": "2015-06-15",
  "divorce_certificate": 2,
  "divorce_certificate_no": "DC1234567890",
  "divorce_certificate_date": "2020-12-31",
  "family_status": 1,
  "mental_disorders": 2,
  "disabilities": 1,
  "education_status": 6,
  "job_type_id": 5,
  "nik_mother": "9876543210123456",
  "mother": "Jane Doe",
  "nik_father": "8765432109876543",
  "father": "Richard Roe",
  "coordinate": "-6.200000, 106.816666"
}

RESPONSE ==>

{
  "status": "CREATED",
  "message": "successfully created new Citizen"
}

{
  "status": "ERROR",
  "message": "citizen with this NIK already exists"
}
```

##### *`UPDATE CITIZEN`*
```
PUT /api/citizens
Content-Type: application/json
X-API-Key: <API_KEY>

{
  "kk": 9876543210987654,
  "full_name": "John Doe",
  "gender": 1,
  "birth_date": "1990-01-01",
  "age": 35,
  "birth_place": "Jakarta",
  "address": "Jl. Merdeka No. 10, RT 01 RW 02",
  "province_id": 1,
  "district_id": 2,
  "sub_district_id": 3,
  "village_id": 4,
  "rt": "01",
  "rw": "02",
  "postal_code": 12345,
  "citizen_status": 1,
  "birth_certificate": 1,
  "birth_certificate_no": "BC1234567890",
  "blood_type": 1,
  "religion": 1,
  "marital_status": 1,
  "marital_certificate": 1,
  "marital_certificate_no": "MC1234567890",
  "marriage_date": "2015-06-15",
  "divorce_certificate": 2,
  "divorce_certificate_no": "DC1234567890",
  "divorce_certificate_date": "2020-12-31",
  "family_status": 1,
  "mental_disorders": 2,
  "disabilities": 1,
  "education_status": 6,
  "job_type_id": 5,
  "nik_mother": "9876543210123456",
  "mother": "Jane Doe",
  "nik_father": "8765432109876543",
  "father": "Richard Roe",
  "coordinate": "-6.200000, 106.816666"
}

RESPONSE ==>

{
  "status": "OK",
  "message": "successfully updated Citizen"
}

```

##### *`DELETE CITIZEN`*
```
DELETE /api/citizens/{nik}
Content-Type: application/json
X-API-Key: <API_KEY>

RESPONSE ==>
{
  "status": "OK",
  "message": "DELETE Citizen successfully"
}

{
  "status": "ERROR",
  "message": "citizen with NIK 1234567890123456 not found"
}
```