# Develop REST API with Echo framework
* Create API
* [Grachful shutdown](https://echo.labstack.com/docs/cookbook/graceful-shutdown)
* API testing with postman and newman

## Step to run
```
$go mod tidy
$go run main.go
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

List of APIs
* http://localhost:1323/hello

## API Testing with Postman and [Newman](https://www.npmjs.com/package/newman)
* Required [NodeJS](https://nodejs.org/en)

Install newman
```
$npm i -g newman
```

Run Postman collection with newman
```
$newman run go-api.postman_collection.json
```

Next => Step 2 :: Refactor and testing api