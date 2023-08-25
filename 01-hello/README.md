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

## Performance testing
* [wrk](https://github.com/wg/wrk)
* [go-wrk](https://github.com/tsliwowicz/go-wrk)
```
$wrk -t 5 -c 100 -d 10s http://localhost:1323/hello
```

## Profiling with [pprof](https://pkg.go.dev/net/http/pprof)
Update file `main.go`

```
import _ "net/http/pprof"

...

e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
```

Access to url = http://localhost:1323/debug/pprof

Working with pprof tool
```
// Heap/Memory
$go tool pprof http://localhost:1323/debug/pprof/heap

// CPU
$go tool pprof http://localhost:1323/debug/pprof/profile
```

Pprof UI with [Graphviz](https://graphviz.org/)
```
$brew install graphviz

$go tool pprof --http=:8888 http://localhost:1323/debug/pprof/heap
```
Access to url => http://localhost:8888/ui/


Next => Step 2 :: Refactor and testing api