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
curl -X GET "http://localhost:8080/card/find_all" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY"
```

### Get ID

```sh
curl -X GET "http://localhost:8080/card/find_by_id?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY"
```

### Create
```sh
curl -X POST "http://localhost:8080/card/create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDAxNDYxfQ.W5Roja93nDVM693jQ7UAzmV1OBLI4HsPzNhgcTJzBgs" \
-d '{
  "user_id": 1,
  "card_type": "credit",
  "expire_date": "2025-12-31T00:00:00Z",
  "cvv": "123",
  "card_provider": "bni"
}'
```

### Update

```sh
curl -X PUT "http://localhost:8080/card/update?=id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY" \
-d '{
  "card_id": 1,
  "user_id": 1,
  "card_type": "credit",
  "expire_date": "2026-06-30T00:00:00Z",
  "cvv": "456",
  "card_provider": "bca"
}'
```

### Delete

```sh
curl -X DELETE "http://localhost:8080/card/delete?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY"
```
-----------------------------------

## Card Receiver

### Get All

```sh
curl -X GET "http://localhost:8080/card/find_all" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4"
```

### Get ID

```sh
curl -X GET "http://localhost:8080/card/find_by_id?id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4"
```

### Create
```sh
curl -X POST "http://localhost:8080/card/create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4" \
-d '{
  "user_id": 2,
  "card_type": "credit",
  "expire_date": "2025-12-31T00:00:00Z",
  "cvv": "123",
  "card_provider": "bni"
}'
```

### Update

```sh
curl -X PUT "http://localhost:8080/card/update?=id=2" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4" \
-d '{
  "card_id": 2,
  "user_id": 2,
  "card_type": "debit",
  "expire_date": "2026-06-30T00:00:00Z",
  "cvv": "456",
  "card_provider": "bca"
}'
```

### Delete

```sh
curl -X DELETE "http://localhost:8080/card/delete?id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4"
```

----------------------------------

# Merchant

## Get All

```sh
curl -X GET "http://localhost:8080/merchant/find_all" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY
"
```

## Get ID

```sh
curl -X GET "http://localhost:8080/merchant/find_by_id?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY
"
```

## Get Name

```sh
curl -X GET "http://localhost:8080/merchant/find_by_name?name=MerchantName" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY"
```

## Create

```sh
curl -X POST "http://localhost:8080/merchant/create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY" \
-d '{
  "name": "Ticket Dota",
  "user_id": 1
}'

```

## Update

```sh
curl -X PUT "http://localhost:8080/merchant/update?=id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY" \
-d '{
  "merchant_id": 1,
  "name": "Updated Merchant",
  "user_id": 123,
  "status": "active"
}'
```

## Delete

```sh
curl -X DELETE "http://localhost:8080/merchant/delete?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDMwMTg5fQ.CrT6yFE93SjCc74y3AGZCEmRNwLK19je-zPwwYuQPgY"
```


---------------------------------

# Transaction

## Get All

```sh
curl -X GET "http://localhost:8080/transaction/find_all" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDI2NzgzfQ.LGNCC1GF8iP_EGp5-YBL6Uif5VNcQ81an6lRy6X3S_Y"
```

## Get ID

```sh
curl -X GET "http://localhost:8080/transaction/find_by_id?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4"
```

## Create

```sh
curl -X POST "http://localhost:8080/transaction/create" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4" \
-d '{
  "card_number": "4460909111027133",
  "amount": 50000,
  "payment_method": "bni",
  "transaction_time": "2024-12-10T15:04:05Z"
}'

```

## Update

```sh
curl -X PUT "http://localhost:8080/transaction/update/?id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4" \
-d '{
  "transaction_id": 1,
  "card_number": "1234-5678-9876-5432",
  "amount": 600000,
  "payment_method": "Debit Card",
  "transaction_time": "2024-12-11T15:04:05Z"
}'
```

## Delete

```sh
curl -X DELETE "http://localhost:8080/transaction/delete?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzM0MDMwMjM0fQ.GtDcZNcsgM-GoD6qhZgpux92Q0tTcsP4bE08L3Yr8u4"
```


-----------------------------------
# Saldo


## Saldo Sender

### Get all Saldo

```sh
curl -X GET http://localhost:8080/saldo/find_all \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDI3ODYyfQ.nSrm-y9YIgFNmsDXYHkEGypQYnKXdA1l_LwCMk_2rAc"
```

### Get Specific Saldo by ID

```sh
curl -X GET "http://localhost:8080/saldo/find_by_id?id=1" \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```

### Create Saldo

```sh
curl -X POST http://localhost:8080/saldo/create \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDAxNDYxfQ.W5Roja93nDVM693jQ7UAzmV1OBLI4HsPzNhgcTJzBgs" \
-H "Content-Type: application/json" \
-d '{
  "card_number": "4460909111027133",
  "total_balance": 500000
}'
```

### Update Saldo

```sh
curl -X PUT "http://localhost:8080/saldo/update?id=1" \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw" \
-H "Content-Type: application/json" \
-d '{
  "saldo_id": 1,
  "card_number": 1,
  "total_balance": 100000,
}'
```

### Delete Saldo

```sh
curl -X DELETE "http://localhost:8080/saldo/delete?id=1" \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjM2NDAsImlhdCI6MTczMjA2MDA0MH0.A61IWywfRTetrqXTy9oBXGGdr5DBss-aU-1-SW46ZCw"
```
-------------------



## Saldo Receiver

### Get all Saldo

```sh
curl -X GET http://localhost:8080/saldo/find_all \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDI3ODYyfQ.nSrm-y9YIgFNmsDXYHkEGypQYnKXdA1l_LwCMk_2rAc"
```

### Get Specific Saldo by ID

```sh
curl -X GET "http://localhost:8080/saldo/find_by_id?id=1" \
-H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s"
```

### Create Saldo

```sh
curl -X POST http://localhost:8080/saldo/create \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDAxNjYzfQ.YEEX3i8h2KbVlF2Dz1dB9BkyxX716et15Tr2WQSuLK8" \
-H "Content-Type: application/json" \
-d '{
  "card_number": "4173630085552615",
  "total_balance": 50000
}'
```

### Update Saldo

```sh
curl -X PUT "http://localhost:8080/saldo/update?id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM0MDAxNjYzfQ.YEEX3i8h2KbVlF2Dz1dB9BkyxX716et15Tr2WQSuLK8" \
-H "Content-Type: application/json" \
-d '{
  "saldo_id": 2,
  "card_number": "4173630085552615",
  "total_balance": 100000
}'
```

### Delete Saldo

```sh
curl -X DELETE "http://localhost:8080/saldo/delete?id=1" \
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
curl -X GET "http://localhost:8080/topup/find_by_id?id=1" \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json"
```


### Create a New Topup

```sh
curl -X POST http://localhost:8080/topup/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json" \
     -d '{
           "card_number": 1,
           "topup_no": "TX123456",
           "topup_amount": 100000,
           "topup_method": "alfamart"
         }'
```

### Update an Existing Topup

```sh
curl -X PUT "http://localhost:8080/topup/update?id=1" \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNjgwNDEsImlhdCI6MTczMjA2NDQ0MX0.sZyj7nmY6RQOyQ0etO76AXTpj5r1MwZQDfnMpVuByo0" \
     -H "Content-Type: application/json" \
     -d '{
           "card_number": 1,
           "topup_id": 1,
           "topup_amount": 150000,
           "topup_method": "indomart"
         }'
```

### Delete a Topup

```sh
curl -X DELETE "http://localhost:8080/topup/delete?id=1" \
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
curl -X GET "http://localhost:8080/topup/find_by_id?id=2" \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json"
```

### Create a New Topup

```sh
curl -X POST http://localhost:8080/topup/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json" \
     -d '{
           "card_number": 2,
           "topup_no": "TX123456",
           "topup_amount": 100000,
           "topup_method": "alfamart"
         }'
```

### Update an Existing Topup

```sh
curl -X PUT "http://localhost:8080/topup/update?id=2" \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNjM2NzcsImlhdCI6MTczMjA2MDA3N30.k017TiFhBpsdLCvSKos10eMT4yd8ieuD_m-qMkfZV3s" \
     -H "Content-Type: application/json" \
     -d '{
           "card_number": 2,
           "topup_id": 2,
           "topup_amount": 150000,
           "topup_method": "indomart"
         }'
```

### Delete a Topup

```sh
curl -X DELETE "http://localhost:8080/topup/delete?id=2" \
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
curl -X GET "http://localhost:8080/transfer/find_by_id?id=1" \
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
curl -X PUT "http://localhost:8080/transfer/update?id=1" \
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
curl -X DELETE "http://localhost:8080/transfer/delete?id=1" \
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

### Create a New Withdraw

```sh
curl -X POST http://localhost:8080/withdraw/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzIwNzI2MTksImlhdCI6MTczMjA2OTAxOX0.we0y1YH05TQ-g46C2Q_v9-rkuQkrwA_H1DghHoSRlHY" \
     -H "Content-Type: application/json" \
     -d '{
             "card_number": 1,
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
         "card_number": 1,
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

### Create a New Withdraw

```sh
curl -X POST http://localhost:8080/withdraw/create \
     -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MzIwNzMxNzMsImlhdCI6MTczMjA2OTU3M30.fSWyOTrtvAUxZtIs3JXe0GZxL-xbOzy0r5bE-TI3ZzY" \
     -H "Content-Type: application/json" \
     -d '{
         "card_number": 1,
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
         "card_number": 2,
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
