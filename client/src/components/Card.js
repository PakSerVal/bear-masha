import React, {useState} from 'react';
import emoji from '../images/emoji_apple.svg';

function Card() {
    const [isCardMenuOpened, setIsCardMenuOpened] = useState(false);
    const [isMenuOpened, setIsMenuOpened] = useState(false);
    const [isMenuButtonHidden, setIsMenuButtonHidden] = useState(false);

    function handleOpenMenuButtonClick() {
        setIsCardMenuOpened(true);
        setIsMenuOpened(true);
        setIsMenuButtonHidden(true);
    }

    function handleCloseMenuButtonClick() {
        setIsCardMenuOpened(false);
        setIsMenuOpened(false);
        setIsMenuButtonHidden(false);
    }

    const isCardMenuOpen = isCardMenuOpened ? 'card_menu-opened' : '';
    const isMenuOpen = isMenuOpened ? 'card__menu_opened' : '';
    const isButtonHidden = isMenuButtonHidden ? 'card__open-button_hidden' : '';

    return (
        <li className={`card ${isCardMenuOpen}`}>
            <div className='card__info'>
                <img className='card__emoji' src={emoji} alt='apple'/>
                <div className='card__description'>
                    <h2 className='card__text card__text_type_title'>Пробежать 7 км</h2>
                    <p className='card__text card__text_type_deadline'>10.10.2022</p>
                    {/*<p className='card__text card__text_type_done'>Выполнено</p>*/}
                </div>
                <button className={`card__open-button button ${isButtonHidden}`} type='button' onClick={handleOpenMenuButtonClick}></button>
            </div>
            <div className={`card__menu ${isMenuOpen}`}>
                <button className='card__close-button button' type='button' onClick={handleCloseMenuButtonClick}></button>
                <button className='card__menu-button link' type='button'>Сделали</button>
                <button className='card__menu-button link' type='button'>Удалить</button>
                <button className='card__menu-button link' type='button'>Редактировать</button>
            </div>
        </li>
    );
}

export default Card;