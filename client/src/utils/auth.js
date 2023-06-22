export const BASE_URL = 'http://localhost:8080/';

export const authorize = (login, password) => {
    return fetch(`${BASE_URL}auth`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        },
        body: "login=" + login + "&" + "password=" + password,
    })
        .then((res) => {
            if (!res.ok) {
                return Promise.reject(`Ошибка: ${res.status}`);
            }
            return res.json();
        });
};