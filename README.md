# amartha_code_test

Question no 1 

System Design
![alt text](image.png)

This design is focusing on functional MVP, which is possible to be enhanced by requirement. Some code logic also based on some business assumption

For DB, i used map since we focus on mvp and simplicity.

Endpoints:
1. /loans [POST]
This endpoint function is to create new loans
Example:
```
curl --location 'http://localhost:8080/loans' \
--header 'Content-Type: application/json' \
--data '{
    "principal": 5000000,
    "interest_rate": 10,
    "installment_term": 15
}'
```

2. /loans/{id}/schedule [GET]
This endpoint function is to show loans payment schedule
```
curl --location 'http://localhost:8080/loans/1/schedule'
```

3. /loans/{id}/outstandingbalance [GET]
This endpoint function is to show outstanding balance of a loan
```
curl --location 'http://localhost:8080/loans/1/outstandingbalance'
```

4. /loans/{id}/status [GET]
This endpoint function is to show delinquent status of loan
```
curl --location 'http://localhost:8080/loans/1/status'
```

5. /loans/{id}/repay [POST]
This endpoint function is to do loan payment
```
curl --location 'http://localhost:8080/loans/1/repay' \
--header 'Content-Type: application/json' \
--data '{
    "paid_amount": 1000000
}'
```