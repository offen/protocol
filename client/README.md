# client
Reference client implementing the Offen protocol

## Usage

The library is currently not available on `npm` or any other package registry.
All code is contained in `index.js` in ESM format.
We might consider publishing it as a npm package once the protocol has become stable.

### Constructing a new client

The client constructor expects an absolute URL for the endpoint.
Further options can be passed in a second and optional argument.

```js
const client = new OffenProtocolClient('https://example.com/protocol-endpoint', options)
```

The following options can be passed to a client:

#### `fetch`

A [`fetch`][mdn-fetch] compatible function that is used for calling the endpoint.
Defaults to `window.fetch`.
Pass a wrapped version of `fetch` if you need to customize behavior.

[mdn-fetch]: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API

#### `serializeBody`

A function that is used to serialize the request body for `register`, `submit` and `purge` calls.
Defaults to `JSON.stringify`.
Pass a custom serializer that accepts a single argument and returns a string your server expects non-JSON encodings.
The function can be either synchronous or asynchronous.

#### `contentType`

The value sent in the `Content-Type` headers when a request body is sent.
Defaults to `application/json`.
Pass a different value in case your server expects other types of payloads, also see `serializeBody`.

#### `handleResponse`

An function that is handling the fetch response.
By default it expects a 2xx response and a JSON response that is parsed and returned as a Javascript object.
204 responses are expected to be empty and return `null`.
The function can be either synchronous or asynchronous.

### Methods on the client

The client exposes a method for each action defined in the protocol.
All methods return a Promise.

#### `probe({ params })`

Send a __probe__ request to the endpoint, appending the given `params` to the URL.
This request intentionally omits sending cookies.

#### `register({ params, body })`

Send a __register__ request to the endpoint, appending the given `params` to the URL and sending the given `body`.
This request intentionally omits sending cookies.

#### `submit({ params, body })`

Send a __submit__ request to the endpoint, appending the given `params` to the URL and sending the given `body`.

#### `query({ params })`

Send a __query__ request to the endpoint, appending the given `params` to the URL.

#### `purge({ params, body })`

Send a __purge__ request to the endpoint, appending the given `params` to the URL and sending the given `body`.
