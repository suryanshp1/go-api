# go-api
A simple go api


## Commands


#### Initiate project

create a folder go-api , then goto that folder after that run following command

```bash
go mod init goapi
```

#### Install external packages

Add external package import in a file and run following command

```bash
go mod tidy
```
After that check go.mod file for dependencies

#### Command to run api

```bash
go run .\cmd\api\main.go
```

#### use CURL/Postman/Insomnia to test API

```bash
curl --request GET \
  --url 'http://localhost:8000/account/coins/?username=alex' \
  --header 'Authorization: 1234' \
  --header 'User-Agent: insomnia/10.3.1'
```