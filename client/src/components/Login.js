import React, { useState } from 'react';

function Login(props) {
    const [data, setData] = useState({
        email: '',
        password: '',
    });

    function handleChange(e) {
        const { name, value } = e.target;
        setData({
            ...data,
            [name]: value,
        });
    }

    function handleSubmit(e) {
        e.preventDefault();
        props.onLogin(data.email, data.password);
    }

    return (
        <div className='form-page'>
            <h2 className='form-page__heading'>Вход</h2>
            <form onSubmit={handleSubmit} className='form-page__form'>
                <div className='form-page__input-wrapper'>
                    <input
                        id='loginEmail'
                        type='text'
                        name='email'
                        className='form-page__input'
                        placeholder='Email'
                        required
                        aria-label='loginEmail'
                        onChange={handleChange}
                    />
                    <input
                        id='loginPassword'
                        type='password'
                        name='password'
                        className='form-page__input'
                        placeholder='Пароль'
                        required
                        aria-label='loginPassword'
                        onChange={handleChange}
                    />
                </div>
                <button type='submit' className='submit-button'>Войти</button>
            </form>
        </div>
    );
}

export default Login;

