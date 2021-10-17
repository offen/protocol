# protocol
A specification for the discoverable exchange of data over a single HTTP endpoint

__This document is an early draft and may change at any time__.
Open an issue on this repository in case you have feedbacks or wish to implement the specification.

## Abstract

Paradigms around the collection of usage data on the web are shifting.
While collection and usage of such data itself isn't likely to go away soon, application developers are looking for new tools and approaches they can use to handle these transactions in a fair and transparent manner.

__The Offen protocol aims for establishing a common set of idioms that can be picked up by implementors that strive for transparency and interoperability__.
It specifies a set of operations that can be used to transparently collect usage data over a single HTTP endpoint and allow users to manage and review the data that is associated to them.

Applications that implement the Offen protocol can be audited by researchers and 3rd party tools like browser extensions.

### Motivation

This document aims to collect the learnings and patterns we have found and implemented when building [Offen][offen-repo].
While the subject seems simple enough, we found it contains unexpected subtleties that are easy to miss, which is why we want to formalize our approach and help other developers to do what we are doing in Offen: let users access their data.

[offen-repo]: https://github.com/offen/offen

## Operations

To describe the collection and handling of data, the Offen protocol defines __a set of five operations__ that should be the base for implementing the exchange of data in an application or library.

### Probe

A client can probe a server in order to acquire additional information about the service or collect prerequisistes that might be needed for registering or submitting data.
This operation is optional for implementations that do not require similar data to be fetched.

### Register

Before submitting any data, a client must __register__ with the server.
The response must send a unique identifier in the response that the client can then use to identify itself in subsequent requests.
Trying to register multiple times is expected to fail, clients that want to re-register are expected to purge their data beforehand.

### Submit

When data is being collected, a client __submits__ a payload to the server.

### Query

At any time, a client can __query__ the server for all data that is associated with it.

### Purge

In case a client decides it wants all of the associated data removed from the server, it performs a __purge__.
Servers can also decide to unregister the client when performing such a purge.

## Mapping operations to HTTP requests

When following the Offen protocol, all client-server communication is expected to be handled by a single endpoint.
Operations map to HTTP verbs and depend on whether the client has registered yet.
The client is expected to be a browser, the server can be anything that handles HTTP.

### Probe

A client probes a server by performing a `GET` request that does not send a user id.

#### Request

```http
GET / HTTP/1.1
Accept: application/json
Host: example.com
```

#### Response

```http
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: application/json
Date: Sun, 17 Oct 2021 06:54:02 GMT

// Application specific response body ...
```

### Register

A client registers with a server by performing a `POST` request that does not send a user id.
In case additional data is required for registration, the client can send it in the request's body.
The server's response must include a user identifier.
In case a user identifier is sent with the request, the server must reject the request with an appropriate status code.

#### Request

```http
POST / HTTP/1.1
Accept: application/json
Host: example.com
```

#### Response

```http
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: application/json
Date: Sun, 17 Oct 2021 06:54:02 GMT
Set-Cookie: user=ea244f5a-f308-416e-9a55-95ff3c8722b0; Path=/; Expires=Thu, 21 Apr 2022 06:54:02 GMT; HttpOnly

// Application specific optional response body ...
```

### Submit

A client submits data to the server by performing a `PUT` request that includes a user id.
Servers can implement different mechanisms for detecting these identifiers.

#### Request

```http
PUT / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=43ce013c-3677-4e93-b0f2-d60b862d481d

// Application specific request body ...
```

#### Response

```http
HTTP/1.1 201 Created
Content-Length: 12
Content-Type: application/json
Date: Sun, 17 Oct 2021 06:54:04 GMT
Set-Cookie: user=ea244f5a-f308-416e-9a55-95ff3c8722b0; Path=/; Expires=Thu, 21 Apr 2022 06:54:04 GMT; HttpOnly

// Application specific response body ...
```


### Query

A client queries a server by performing a `GET` request that includes a user id.

#### Request

```http
GET / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=43ce013c-3677-4e93-b0f2-d60b862d481d

// Application specific request body ...
```

#### Response

```http
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: application/json
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

// Application specific response body ...
```

### Purge

A client purges its data on the server by performing a `DELETE` request that includes a user id.
In case the client wishes to deregister itself, it has to signal this in the request body.

#### Request

```http
DELETE / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com
Cookie: user=43ce013c-3677-4e93-b0f2-d60b862d481d

// Application specific request body ...
```

#### Response

```http
HTTP/1.1 204 No Content
Content-Length: 12
Date: Fri, 15 Oct 15 2021 12:04:12 GMT
Set-Cookie: user=ea244f5a-f308-416e-9a55-95ff3c8722b0; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT; HttpOnly
```
