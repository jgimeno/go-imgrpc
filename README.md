How to launch the microservice?

```
    go run server/main.go
```

And from the client we can save an image:
```
go run client/main.go save filepath/file.png
```

After it it returns an id.

We can retrieve the image in the file format we want,
```
go run client/main.go get 6fe9c852-1190-11e8-9375-7200063ff500 png
```

It only supports jpg and png for now.

Running the tests:
```
go test ./...
```
