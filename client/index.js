export default class Client {
  constructor (endpoint, options = {}) {
    this.endpoint = endpoint
    this.fetch = options.fetch || window.fetch.bind(window)
    this.serializeBody = options.serializeBody || JSON.stringify
    this.contentType = options.contentType || 'application/json'
    this.handleResponse = options.handleResponse || handleResponse
  }

  _request (method, credentials, { body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.search = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method,
      credentials,
      body: typeof body === 'undefined'
        ? undefined
        : this.serializeBody(body),
      headers: {
        'Content-Type': this.contentType
      }
    })
      .then(this.handleResponse)
  }

  probe ({ params } = {}) {
    return this._request('GET', 'omit', { params })
  }

  register ({ body, params } = {}) {
    return this._request('POST', 'omit', { body, params })
  }

  submit ({ body, params } = {}) {
    return this._request('PUT', 'include', { body, params })
  }

  query ({ params } = {}) {
    return this._request('GET', 'include', { params })
  }

  purge ({ body, params } = {}) {
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
        const err = new Error(errorBody.error)
        err.status = response.status
        throw err
      })
  }
  if (response.status !== 204) {
    return response.json()
  }
  return null
}
