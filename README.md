# protocol
A specification for discoverable exchange of data over a single HTTP endpoint

## Abstract

Paradigms around the collection of usage data on the web are shifting.
While collection and usage of such data itself isn't likely to go away soon, application developers are looking for new tools and approaches they can use to handle these transactions in a fair and transparent manner.

The Offen protocol aims for creating a common set of idioms that can be picked up by implementors that strive for maximum transparency and interoperability.
It specifies a set of operations that can be used to transparently collect usage data over a single HTTP endpoint and allow users to manage and review the data that is associated to them.

## Operations

To describe the collection and handling of data, the Offen protocol defines a set of five operations that should be the base for implementing the exchange of data in an application or library.

### Probe

A client can probe a server in order to acquire additional information about the service or collect prerequisistes that might be needed for registering or submitting data.
This operation is optional for implementations that do not require similar data to be fetched.

### Register

Before submitting any data, a client needs to register with the server.
The response must send a unique identifier in the response that the client can then use to identify itself in subsequent requests.
Trying to register when already having registered must fail.
Clients that want to re-register are expected to purge their data beforehand.

### Submit

When data is being collected, a client submits a payload to the server.

### Query

At any time, a client can query the server for all data that is associated with it.

### Purge

In case a client decides it wants all of the associated data removed from the server, it performs a purge.
Servers can also decide to unregister the client when performing such a purge.

## Mapping operations to HTTP requests

When following the Offen protocol, all client-server communication is expected to be handled by a single endpoint.
Operations map to HTTP verbs and depend on whether the client has registered yet.

### Probe

A client probes a server by performing a `GET` request that does not send a user id.

### Register

A client registers with a server by performing a `POST` request that does not send a user id.
In case additional data is required for registration, the client can send it in the request's body.
The server's response must include a user identifier.

### Submit

A client submits data to the server by performing a `POST` request that includes a user id.
Servers can implement different mechanisms for detecting these identifiers.

### Query

A client queries a server by performing a `GET` request that includes a user id.

### Purge

A client purges its data on the server by performing a `DELETE` request that includes a user id.
In case the client wishes to deregister itself, it has to signal this in the request body.

## Example flow

### 1. Probe

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

### 3. Submit

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

### 4. Query

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

### 5. Purge

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
