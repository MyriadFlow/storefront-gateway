docker run --name="marketplace" --rm -d -p 5432:5432 \
-e POSTGRES_PASSWORD=lazarus \
-e POSTGRES_USER=lazarus \
-e POSTGRES_DB=lazarus \
postgres -c log_statement=all

cp .env-sample .env