1. 
$ go mod init example/web-service-gin

2. 
$ go get .

3. 
$ go run .

4. GET
$ curl http://localhost:8080/albums

5. POST

```shell
$ curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4", "title": "The Modern Second of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```

or

```shell
curl http://localhost:8080/albums \
--include --header \
"Content-Type: application/json" \
--request "POST" --data \
'{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

```shell
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

6. GET again

```shell
curl http://localhost:8080/albums --header "Content-Type: application/json"
--request "GET"
```

or 

```shell
curl http://localhost:8080/albums \
--header \
"Content-Type: application/json" \
--request "GET"
```

7. GET getAlbumByID

```shell
curl http://localhost:8080/albums/2
```