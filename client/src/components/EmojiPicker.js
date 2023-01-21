import React, { useState } from 'react';
import emoji from '../utils/emoji';

function EmojiPicker() {
    return (
        <div className='emoji-picker'>
            <div className='emoji-picker__input'>
                <p className='emoji-picker__placeholder'>Выберете эмоджи</p>
            </div>
            <div className='emoji-picker__container'>
                {
                    emoji.map((img) => <img className='emoji-picker__item' src={img.src} alt={img.name}/>)
                }
            </div>
        </div>
    );
}

export default EmojiPicker;