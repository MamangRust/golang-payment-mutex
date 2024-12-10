# Auth

## User Sender

### Register User

```sh
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@example.com",
  "password": "password123",
  "confirm_password": "password123"
}'
```

### Login User

```sh
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "john.doe@example.com",
  "password": "password123"
}'
```




## User Receiver


### Register User

```sh
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{
  "firstname": "Jane",
  "lastname": "Doe",
  "email": "jane.doe@example.com",
  "password": "password123",
  "confirm_password": "password123"
}'
```

### Login User

```sh
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
    "email": "jane.doe@example.com",
  "password": "password123"
}'
```

-----------------

# Card

## Card Sender

### Get All

```sh

```



### Get ID

```sh
```

### Create
```sh
```

### Update

```sh

```

### Delete

```sh

```
-----------------------------------

## Card Receiver

### Get All

```sh
```

### Get ID

```sh
```

### Create

```sh
```

### Update

```sh
```

### Delete

```sh
```

----------------------------------

# Merchant

## Get All

## Get ID

## Create

## Update

## Delete


---------------------------------

# Transaction

## Get All

## Get ID

## Create

## Update

## Delete


-----------------------------------
# Saldo


## Saldo Sender

### Get all Saldo

```sh
curl -X GET http://localhost:8080/saldo/find_all \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjI4ODIsImlhdCI6MTczMjA1OTI4Mn0.oXDLZmtHc7vjbauWW9eBqd5s8sIutK6o3gkQxYBV1jc"
```

### Get Specific Saldo by ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_id/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```

### Get Saldos for All Users by User ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_users_id \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```

### Get Saldo for a Specific User by ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_user_id \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```

### Create Saldo

```sh
curl -X POST http://localhost:8080/saldo/create \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw" \
-H "Content-Type: application/json" \
-d '{
  "user_id": 1,
  "total_balance": 50000
}'
```

### Update Saldo

```sh
curl -X PUT http://localhost:8080/saldo/update/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw" \
-H "Content-Type: application/json" \
-d '{
  "saldo_id": 1,
  "user_id": 1,
  "total_balance": 100000,
  "withdraw_amount": 50000,
  "withdraw_time": null
}'
```

### Delete Saldo

```sh
curl -X DELETE http://localhost:8080/saldo/delete/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```
-------------------



## Saldo Receiver

### Get all Saldo

```sh
curl -X GET http://localhost:8080/saldo/find_all \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

### Get Specific Saldo by ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_id/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

### Get Saldos for All Users by User ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_users_id \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

### Get Saldo for a Specific User by ID

```sh
curl -X GET http://localhost:8080/saldo/find_by_user_id \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

### Create Saldo

```sh
curl -X POST http://localhost:8080/saldo/create \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
-H "Content-Type: application/json" \
-d '{
  "user_id": 2,
  "total_balance": 50000
}'
```

### Update Saldo

```sh
curl -X PUT http://localhost:8080/saldo/update/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
-H "Content-Type: application/json" \
-d '{
  "saldo_id": 2,
  "user_id": 2,
  "total_balance": 100000,
  "withdraw_amount": 50000,
  "withdraw_time": null
}'
```

### Delete Saldo

```sh
curl -X DELETE http://localhost:8080/saldo/delete/1 \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

-----------------------------------


# Topup

## Topup Sender

### Get All Topups

```sh
curl -X GET http://localhost:8080/api/topup/find_all \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```

### Get Topup by ID

```sh
curl -X GET http://localhost:8080/topup/find_by_id/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```

### Get Topup by User ID

```sh
curl -X GET http://localhost:8080/topup/find_by_user_id \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```

### Get All Topups for a Specific User

```sh
curl -X GET http://localhost:8080/topup/find_by_users_id \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```

### Create a New Topup

```sh
curl -X POST http://localhost:8080/topup/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json" \
     -d '{
           "user_id": 1,
           "topup_no": "TX123456",
           "topup_amount": 100000,
           "topup_method": "alfamart"
         }'
```

### Update an Existing Topup

```sh
curl -X PUT http://localhost:8080/topup/update/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json" \
     -d '{
           "user_id": 1,
           "topup_id": 1,
           "topup_amount": 150000,
           "topup_method": "indomart"
         }'
```

### Delete a Topup

```sh
curl -X DELETE http://localhost:8080/topup/delete/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```



## Topup Receiver

### Get All Topups

```sh
curl -X GET http://localhost:8080/topup/find_all \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```

### Get Topup by ID

```sh
curl -X GET http://localhost:8080/topup/find_by_id/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```

### Get Topup by User ID

```sh
curl -X GET http://localhost:8080/topup/find_by_user_id \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```

### Get All Topups for a Specific User

```sh
curl -X GET http://localhost:8080/api/topup/find_by_users_id \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```

### Create a New Topup

```sh
curl -X POST http://localhost:8080/topup/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json" \
     -d '{
           "user_id": 2,
           "topup_no": "TX123456",
           "topup_amount": 100000,
           "topup_method": "alfamart"
         }'
```

### Update an Existing Topup

```sh
curl -X PUT http://localhost:8080/topup/update/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json" \
     -d '{
           "user_id": 2,
           "topup_id": 2,
           "topup_amount": 150000,
           "topup_method": "indomart"
         }'
```

### Delete a Topup

```sh
curl -X DELETE http://localhost:8080/topup/delete/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```
-----------------------------



# Transfer

## Transfer Sender

## Get All Transfers
```sh
curl -X GET http://localhost:8080/transfer/find_all \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

## Get Transfer by ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_id \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

## Get Transfer Users by Transfer ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_users \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

## Get Transfer User by Transfer ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_user \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"

```

## Create a New Transfer

```sh
curl -X POST http://localhost:8080/transfer/create \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY" \
    -H "Content-Type: application/json" \
    -d '{
        "transfer_from": 1,
        "transfer_to": 2,
        "transfer_amount": 1000
    }'
```

## Update an Existing Transfer

```sh
curl -X PUT http://localhost:8080/transfer/update/1 \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY" \
    -H "Content-Type: application/json" \
    -d '{
        "transfer_id": 1,
        "transfer_from": 1,
        "transfer_to": 2,
        "transfer_amount": 100000
    }'
```

## Delete a Transfer

```sh
curl -X DELETE http://localhost:8080/transfer/delete/1 \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```


## Transfer Receiver

## Get All Transfers
```sh
curl -X GET http://localhost:8080/transfer/find_all \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

## Get Transfer by ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_id \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

## Get Transfer Users by Transfer ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_users_id \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

## Get Transfer User by Transfer ID

```sh
curl -X GET http://localhost:8080/transfer/find_by_user_id \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"

```

## Create a New Transfer

```sh
curl -X POST http://localhost:8080/transfer/create \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY" \
    -H "Content-Type: application/json" \
    -d '{
        "transfer_from": 2,
        "transfer_to": 1,
        "transfer_amount": 1000
    }'
```

## Update an Existing Transfer

```sh
curl -X PUT http://localhost:8080/transfer/update/2 \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY" \
    -H "Content-Type: application/json" \
    -d '{
        "transfer_id": 2,
        "transfer_from": 2,
        "transfer_to": 1,
        "transfer_amount": 100000
    }'
```

## Delete a Transfer

```sh
curl -X DELETE http://localhost:8080/transfer/delete/2 \
    -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```
----------------------------


# Withdraw

## Withdraw Sender

### Get All Withdraws

```sh
curl -X GET http://localhost:8080/withdraw \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```


### Get Withdraw by ID

```sh
curl -X GET http://localhost:8080/withdraw/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

### Get Withdraw Users

```sh
curl -X GET http://localhost:8080/withdraw/find_by_users \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

### Get Withdraw User

```sh
curl -X GET http://localhost:8080/withdraw/find_by_user \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```

### Create a New Withdraw

```sh
curl -X POST http://localhost:8080/withdraw/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY" \
     -H "Content-Type: application/json" \
     -d '{
             "user_id": 1,
             "withdraw_amount": 100000,
             "withdraw_time": "2024-12-08T10:30:00Z"
         }'
```

### Update Withdraw

```sh
curl -X PUT http://localhost:8080/withdraw/update/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY" \
     -H "Content-Type: application/json" \
     -d '{
         "user_id": 1,
         "withdraw_id": 1,
         "withdraw_amount": 100000,
         "withdraw_time": "2024-12-08T10:30:00Z"
     }'
```

### Delete Withdraw

```sh
curl -X DELETE http://localhost:8080/withdraw/delete/1 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY"
```


## Withdraw Receiver

### Get All Withdraws

```sh
curl -X GET http://localhost:8080/withdraw \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```


### Get Withdraw by ID

```sh
curl -X GET http://localhost:8080/withdraw/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

### Get Withdraw Users

```sh
curl -X GET http://localhost:8080/withdraw/find_by_users \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

## Get Withdraw User

```sh
curl -X GET http://localhost:8080/withdraw/find_by_user \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```

### Create a New Withdraw

```sh
curl -X POST http://localhost:8080/withdraw/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY" \
     -H "Content-Type: application/json" \
     -d '{
         "user_id": 1,
         "withdraw_amount": 100000,
         "withdraw_time": "2024-12-08T10:30:00Z"
     }'
```

### Update Withdraw

```sh
curl -X PUT http://localhost:8080/withdraw/update/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY" \
     -H "Content-Type: application/json" \
     -d '{
         "user_id": 2,
         "withdraw_id": 2,
         "withdraw_amount": 100000,
         "withdraw_time": "2024-12-08T10:30:00Z"
     }'
```

### Delete Withdraw

```sh
curl -X DELETE http://localhost:8080/withdraw/delete/2 \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY"
```
