# Go DDD Skeleton

# How to run
1. Create database with the same name wrote in GO_DDD_SKELETON_SQL_DB_NAME environment variable.
2. Change the connection parameters needed to connect to the database.
3. Execute dump.sql to insert test data.
4. Initialize environment variables. See env folder
5. Execute next command:
```sh
$  go run policy/cmd/policy-api/*.go 
```

or

1. Create database with the same name wrote in GO_DDD_SKELETON_SQL_DB_NAME environment variable.
2. Change the connection parameters needed to connect to the database.
3. Execute dump.sql to insert test data.
4. Execute next commands:
```sh
$ docker network create golang // same network as the infrastructure
$ docker-compose up 
```