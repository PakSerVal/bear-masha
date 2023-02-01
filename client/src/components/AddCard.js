import React from 'react';
import EmojiPicker from './EmojiPicker';

function AddCard() {
    return (
        <div className='form-page'>
            <h2 className='form-page__heading'>Новый план</h2>
            <EmojiPicker />
            <form className='form-page__form'>
                <div className='form-page__input-wrapper'>
                    <input
                        id='plan-title'
                        type='text'
                        name='plan-title'
                        className='form-page__input'
                        placeholder='Введите ваш план'
                        required
                        aria-label='plan-title'
                    />
                    <input
                        id='plan-deadline'
                        type='date'
                        name='plan-deadline'
                        className='form-page__input form-page__input_type_deadline'
                        required
                        aria-label='plan-deadline'
                    />
                </div>
                <button type='submit' className='submit-button'>Добавить</button>
            </form>
        </div>
    );
}

export default AddCard;