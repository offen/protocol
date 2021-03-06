/**
* Copyright 2021 - Offen Authors <hioffen@posteo.de>
* SPDX-License-Identifier: Apache-2.0
*/

export default class OffenProtocolClient {
  constructor (endpoint, options = {}) {
    this.endpoint = endpoint
    this.fetch = options.fetch || window.fetch.bind(window)
    this.serializeBody = options.serializeBody || JSON.stringify
    this.contentType = options.contentType || 'application/json'
    this.handleResponse = options.handleResponse || this._handleResponse
  }

  probe ({ params } = {}) {
    return this._request('GET', 'omit', { params })
  }

  register ({ params, body } = {}) {
    return this._request('POST', 'omit', { body, params })
  }

  submit ({ params, body } = {}) {
    return this._request('PUT', 'include', { body, params })
  }

  query ({ params } = {}) {
    return this._request('GET', 'include', { params })
  }

  purge ({ params, body } = {}) {
    return this._request('DELETE', 'include', { body, params })
  }

  _request (method, credentials, { body, params }) {
    const url = new window.URL(this.endpoint)
    if (params) {
      url.search = new window.URLSearchParams(params)
    }

    const serializedBody = typeof body === 'undefined'
      ? undefined
      : this.serializeBody(body)

    return Promise.resolve(serializedBody)
      .then((body) => {
        return this.fetch(url, {
          method,
          credentials,
          body,
          headers: body
            ? { 'Content-Type': this.contentType }
            : undefined
        })
      })
      .then(this.handleResponse)
  }

  _handleResponse (response) {
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
}
