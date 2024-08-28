## Running unit tests

`go test ./...`

## Running the server

1. `docker-compose up --build` or `docker compose up --build`

## Making Requests

1. You can make requests to `localhost:80`
   1. example curl:

```
curl --location 'localhost:80/api/strong_password_steps' \
--header 'Content-Type: application/json' \
--header 'Cookie: redirect_to=%2Fstaff%2Fv1%2Forders' \
--data '{
    "init_password": "aaaBc1"
}'
```

## Getting the access log

1. Seeing access logs in postgresql:
   1. connect to localhost with these parameters
      1. host: `localhost`
      2. port: `5433`
      3. username: `myuser`
      4. password: `mypassword`
      5. database: `agnos`
