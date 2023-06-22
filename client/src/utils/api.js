class Api {

  getInitialsCards() {
    return fetch('http://localhost:8080/' + 'cards', {
      method: 'GET',
      headers: {
        'X-Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU2OTYzMjgsIm5iZiI6MTY3MzEwNDMyOH0.4-nZZdIOH7lGOANpkTWt0pc0icTGVpSeKTs3RpzKWvE',
        'Content-Type': 'application/x-www-form-urlencoded',
      }
    })
        .then((response) => {
          if (response.ok) {
            return response.json();
          }
          return Promise.reject(new Error(response.status));
        });
  }

  addCard({ title, description, deadlineAt }) {
    return this._request(this._baseUrl + 'cards', {
      method: 'POST',
      headers: this._header,
      body: { title, description, deadlineAt },
    })
  }

  _request(url, options) {
    return fetch(url, options).then(this._checkResponse);
  }

  _checkResponse(response) {
    if (response.ok) {
      return response.json();
    }
    return Promise.reject(new Error(response.status));
  }
}

const api = new Api({
  // baseUrl: 'http://localhost:8080/',
  // headers: {
  //   'X-Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU2OTYzMjgsIm5iZiI6MTY3MzEwNDMyOH0.4-nZZdIOH7lGOANpkTWt0pc0icTGVpSeKTs3RpzKWvE',
  //   'Content-Type': 'application/x-www-form-urlencoded',
  // }
});

export default api;