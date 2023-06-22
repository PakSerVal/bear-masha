import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import logo from '../images/logo.svg';

function Header() {
    const navigate = useNavigate();
    function signOut(){
        localStorage.removeItem('jwt');
        navigate('/sign-in');
    }

    return (
        <header className='header page__header section'>
            <img className='header__logo' alt='Логотип' src={logo}/>
            <NavLink onClick={signOut} to='/sign-in' className='header__link link'>Выйти</NavLink>
        </header>
    );
}

export default Header;