import React, { useState } from 'react';
import emojis from '../utils/emoji';

function EmojiPicker() {
    const [isEmojiPickerOpen, setIsEmojiPickerOpen] = useState(false);
    const [chosenEmoji, setChosenEmoji] = useState(null);

    function handleInputButtonClick() {
        setIsEmojiPickerOpen(!isEmojiPickerOpen);
    }

    function handleEmojiClick(emoji) {
        setChosenEmoji(emoji);
        setIsEmojiPickerOpen(false);
    }

    const isEmojiContainerOpen = isEmojiPickerOpen ? 'emoji-picker__container_opened' : '';

    return (
        <div className='emoji-picker'>
            <div className={`emoji-picker__input ${isEmojiPickerOpen ? 'emoji-picker__input_chosen' : ''}`} onClick={handleInputButtonClick}>
                {
                    chosenEmoji
                        ? <img className='emoji-picker__item emoji-picker__item_type_input' src={chosenEmoji.src} alt={chosenEmoji.name}/>
                        : <p className='emoji-picker__placeholder'>Выберете эмоджи</p>
                }
                {
                    isEmojiPickerOpen
                        ? <button className='emoji-picker__input-button emoji-picker__input-button_upside button' type='button'></button>
                        : <button className='emoji-picker__input-button button' type='button'></button>
                }
            </div>
            <div className={`emoji-picker__container ${isEmojiContainerOpen}`}>
                {
                    emojis.map((img) => <img className='emoji-picker__item' src={img.src} alt={img.name} onClick={() => handleEmojiClick(img)} />)
                }
            </div>
        </div>
    );
}

export default EmojiPicker;