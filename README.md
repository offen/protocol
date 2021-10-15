# protocol
A specification for transparent and discoverable exchange of data over a single HTTP endpoint

## Example flow

### 1. Acquire prerequisites (optional)

Request:

```
GET / HTTP/1.1
Accept: application/json
Host: example.com
```

Response:

```
HTTP/1.1 201 Created
Content-Length: 12
Content-Type: application/json
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

{
  "response": {}
}
```

### 2. Register

Request:

```
POST / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com

{
  "payload": {}
}

```

Response:

```
HTTP/1.1 204 No Content
Set-Cookie: user=a7715269-1d77-4162-b1ee-fc3a050d7998; Path=/; Expires=Tue, 19 Apr 2022 10:05:11 GMT; HttpOnly
Date: Fri, 15 Oct 15 2021 12:04:12 GMT
```

### 3. Submit data

Request:

```
POST / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=a7715269-1d77-4162-b1ee-fc3a050d7998

{
  "payload": {}
}

```

Response:

```
HTTP/1.1 201 Created
Content-Type: application/json
Set-Cookie: user=a7715269-1d77-4162-b1ee-fc3a050d7998; Path=/; Expires=Tue, 19 Apr 2022 10:05:11 GMT; HttpOnly
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

{ "ack": true }
```

### 4. Query data

Request:

```
GET / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=a7715269-1d77-4162-b1ee-fc3a050d7998
```

Response:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

{ "data": {} }
```

### 5. Delete data

Request:

```
DELETE / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=a7715269-1d77-4162-b1ee-fc3a050d7998
```

Response:

```
HTTP/1.1 204 No Content
Date: Fri, 15 Oct 15 2021 12:04:12 GMT
```
