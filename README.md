# ecomm
---

Config database (Example):
    
    export DB_HOST="127.0.0.1"
    export DB_PORT="5432"
    export DB_DBNAME="ecomm"
    export DB_USER="postgres"
    export DB_PASSWORD="postgres"
    export DB_SSL="disable"


### Migration

Run program with sub arguments, if len(Args) > 1 run migration and create tables.

    go run main.go create-db

### WIP

Domains:

- [ ] Product
- [ ] Client
- [ ] Store
- [ ] Payment

### Client

Register use same URL with POST method.

    curl --location --request POST 'http://localhost:8080/api/v1/client' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "username": "teste",
            "password": "teste1",
            "email": "teste1@gmail.com"
        }'
