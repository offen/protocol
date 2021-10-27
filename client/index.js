export default class Client {
  constructor (endpoint, options = {}) {
    this.endpoint = endpoint
    this.fetch = options.fetch || window.fetch
  }

  probe ({ body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method: 'GET',
      body: body && JSON.stringify(body)
    })
      .then(handleFetchResponse)
  }

  register ({ body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method: 'POST',
      body: body && JSON.stringify(body)
    })
      .then(handleFetchResponse)
  }

  submit ({ body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method: 'PUT',
      body: body && JSON.stringify(body),
      credentials: 'include'
    })
      .then(handleFetchResponse)
  }

  query ({ body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method: 'GET',
      body: body && JSON.stringify(body),
      credentials: 'include'
    })
      .then(handleFetchResponse)
  }

  purge ({ body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.searchParams = new window.URLSearchParams(params)
    }
    return this.fetch(url, {
      method: 'DELETE',
      body: body && JSON.stringify(body),
      credentials: 'include'
    })
      .then(handleFetchResponse)
  }
}

function handleFetchResponse (response) {
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
