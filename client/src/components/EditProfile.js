import React from 'react';

function EditProfile() {
    return (
        <div className='form-page'>
            <h2 className='form-page__heading'>Редактировать профиль</h2>
            <form className='form-page__form'>
                <div className='form-page__input-wrapper'>
                    <input
                        id='loginEmail'
                        type='email'
                        name='email'
                        className='form-page__input'
                        placeholder='Email'
                        required
                        aria-label='loginEmail'
                    />
                    <input
                        id='loginPassword'
                        type='password'
                        name='password'
                        className='form-page__input'
                        placeholder='Пароль'
                        required
                        aria-label='loginPassword'
                    />
                </div>
                <button type='submit' className='form-page__submit-button'>Войти</button>
            </form>
        </div>
    );
}

export default EditProfile;