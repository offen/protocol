# server

Package server provides primitives to implement an HTTP server that adheres
to the Offen protocol. A default http.Handler is returned and can be used
in any setup.

## Functions

### func [NewHandler](/handler.go#L17)

`func NewHandler(opts ...Option) http.Handler`

NewHandler returns an http.Handler that implements the Offen protocol. Additional
behavior can be specified by passing an arbitrary number of options.

## Types

### type [Adapter](/handler.go#L95)

`type Adapter func(*http.Request, string) (HTTPResult, error)`

Adapter is a function that transforms an HTTP request and a user identifier
into a result type. Implementors can use this to add any kind of application
specific logic to the handler.

### type [HTTPResult](/handler.go#L99)

`type HTTPResult struct { ... }`

HTTPResult is returned by an Adapter. The handler then writes body,
status code and cookie headers accordingly.

### type [Option](/handler.go#L90)

`type Option func(*server)`

Option provides a mechanism for passing optional parameters when calling New.

#### func [WithCookieAttributeDomain](/handler.go#L170)

`func WithCookieAttributeDomain(d string) Option`

WithCookieAttributeDomain sets a `Domain` attribute for the issued cookie.

#### func [WithCookieAttributePath](/handler.go#L165)

`func WithCookieAttributePath(p string) Option`

WithCookieAttributePath sets a `Path` attribute for the issued cookie.

#### func [WithCookieAttributeSameSite](/handler.go#L188)

`func WithCookieAttributeSameSite(v http.SameSite) Option`

WithCookieAttributeSameSite defines the value used for the cookies' `SameSite`
attribute.

#### func [WithCookieAttributeSecure](/handler.go#L182)

`func WithCookieAttributeSecure(b bool) Option`

WithCookieAttributeSecure sets the `Secure` attribute used when issuing
cookies.

#### func [WithCookieName](/handler.go#L160)

`func WithCookieName(n string) Option`

WithCookieName overrides the default cookie name of "user" which is used for
storing the assigned user identifier.

#### func [WithCookieTTL](/handler.go#L176)

`func WithCookieTTL(t time.Duration) Option`

WithCookieTTL sets defines the time to live that is used when calculating
a cookie's `Expires` attribute.

#### func [WithProbeAdapter](/handler.go#L133)

`func WithProbeAdapter(a Adapter) Option`

WithProbeAdapter sets an adapter that is used when probing the endpoint.

#### func [WithPurgeAdapter](/handler.go#L154)

`func WithPurgeAdapter(a Adapter) Option`

WithPurgeAdapter sets an adapter that is used when purging data.

#### func [WithQueryAdapter](/handler.go#L149)

`func WithQueryAdapter(a Adapter) Option`

WithQueryAdapter sets an adapter that is used when querying the endpoint.

#### func [WithRegisterAdapter](/handler.go#L139)

`func WithRegisterAdapter(a Adapter) Option`

WithRegisterAdapter sets an adapter that is used when registering against
the endpoint.

#### func [WithSubmitAdapter](/handler.go#L144)

`func WithSubmitAdapter(a Adapter) Option`

WithSubmitAdapter sets an adapter that is used when submitting data.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
