# Working with dependencies
* Database/Resources


## How to inject all dependencies into Handler/Router ?
```
func SayHi(c echo.Context) (err error) {
	res := HelloResponse{Message: "Hello world"}
	return c.JSON(http.StatusOK, res)
}
```

## Solution 1 :: Use global variable !!

### Solution 2 :: Send to SayHi(database)
* Add model.go
* Update server.go
* Update handler.go
* Update handler_test.go

```
```

Testing ...

### Solution 3 :: XXX
```
```

Testing ...


Next => Step 4 :: TODO