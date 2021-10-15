# protocol
A specification for discoverable exchange of data over a single HTTP endpoint

__This document is an early draft and may change at any time__.
Open an issue on this repository in case you have feedbacks or wish to implement the specification.

## Abstract

Paradigms around the collection of usage data on the web are shifting.
While collection and usage of such data itself isn't likely to go away soon, application developers are looking for new tools and approaches they can use to handle these transactions in a fair and transparent manner.

The Offen protocol aims for creating a common set of idioms that can be picked up by implementors that strive for maximum transparency and interoperability.
It specifies a set of operations that can be used to transparently collect usage data over a single HTTP endpoint and allow users to manage and review the data that is associated to them.

### Motivation

This document aims to collect the learnings and patterns we have found and implemented when building [Offen][offen-repo].
While the subject seems simple enough, we found it contains unexpected subtleties that are easy to miss, which is why we want to formalize our approach and help other developers to do what we are doing in Offen: let users access their data.

[offen-repo]: https://github.com/offen/offen

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
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

// Application specific response body ...
```

### Register

A client registers with a server by performing a `POST` request that does not send a user id.
In case additional data is required for registration, the client can send it in the request's body.
The server's response must include a user identifier.
In case a user identifier is sent with the request, the server must reject it.

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
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

// Application specific optional response body ...
```

#### Receiving the user identifier

In case the client is a browser, the most probable way of sending the user identifier is using the `Set-Cookie` header.
In case the client is another service other means of submitting the identifier can be chosen.

### Submit

A client submits data to the server by performing a `PUT` request that includes a user id.
Servers can implement different mechanisms for detecting these identifiers.

#### Request

```http
PUT / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com

// Application specific request body ...
```

#### Response

```http
HTTP/1.1 201 Created
Content-Length: 12
Content-Type: application/json
Date: Fri, 15 Oct 15 2021 12:04:12 GMT

// Application specific response body ...
```

#### Sending the user identifier

In case the client is a browser, the most probable way of sending the user identifier is using the `Cookie` header.
In case the client is another service other means of submitting the identifier can be chosen.

### Query

A client queries a server by performing a `GET` request that includes a user id.

#### Request

```http
GET / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com

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

#### Sending the user identifier

In case the client is a browser, the most probable way of sending the user identifier is using the `Cookie` header.
In case the client is another service other means of submitting the identifier can be chosen.

### Purge

A client purges its data on the server by performing a `DELETE` request that includes a user id.
In case the client wishes to deregister itself, it has to signal this in the request body.

#### Request

```http
DELETE / HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: example.com

// Application specific request body ...
```

#### Response

```http
HTTP/1.1 204 No Content
Content-Length: 12
Date: Fri, 15 Oct 15 2021 12:04:12 GMT
```

#### Sending the user identifier

In case the client is a browser, the most probable way of sending the user identifier is using the `Cookie` header.
In case the client is another service other means of submitting the identifier can be chosen.
