import React from 'react';
import { NavLink } from 'react-router-dom';
import logo from '../images/logo.svg';

function Header() {
    return (
        <header className='header page__header section'>
            <img className='header__logo' alt='Логотип' src={logo}/>
            <div className='header__link-wrapper'>
                <NavLink to='/edit-profile' className='header__link link'>Изменить профиль</NavLink>
                <NavLink to='/sign-in' className='header__link link'>Выйти</NavLink>
            </div>
        </header>
    );
}

export default Header;