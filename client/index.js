export default class Client {
  constructor (endpoint, options = {}) {
    this.endpoint = endpoint
    this.fetch = options.fetch || window.fetch
    this.serializeBody = options.serializeBody || JSON.stringify
    this.contentType = options.contentType || 'application/json'
    this.handleResponse = options.handleResponse || handleResponse
  }

  _request (method, credentials, { body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method,
      body: typeof body !== 'undefined' && this.serializeBody(body),
      credentials,
      headers: {
        'Content-Type': this.contentType
      }
    })
      .then(this.handleResponse)
  }

  probe ({ body, params }) {
    return this._request('GET', 'omit', { body, params })
  }

  register ({ body, params }) {
    return this._request('POST', 'omit', { body, params })
  }

  submit ({ body, params }) {
    return this._request('PUT', 'include', { body, params })
  }

  query ({ body, params }) {
    return this._request('GET', 'include', { body, params })
  }

  purge ({ body, params }) {
    return this._request('DELETE', 'include', { body, params })
  }
}

function handleResponse (response) {
  if (response.status >= 400) {
    return response.clone().json()
      .catch(function () {
        return response.text()
          .then(function (rawBody) {
            return { error: rawBody }
          })
      })
      .then(function (errorBody) {
        var err = new Error(errorBody.error)
        err.status = response.status
        throw err
      })
  }
  if (response.status !== 204) {
    return response.json()
  }
  return null
}
