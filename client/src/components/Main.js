import React from 'react';
import Card from './Card';
import { NavLink } from 'react-router-dom';
import avatar from '../images/default-avatar.svg';

function Main() {
    return (
        <main className='content page__content'>
            <section className='profile section content__section'>
                <img src={avatar} className='profile__avatar' alt='Аватарка'/>
                <div className='profile__info'>
                    <h1 className='profile__text profile__text_type_name'>Мария, Сергей и Дейзи</h1>
                    <p className='profile__text profile__text_type_status'>Счастливые пёсьи родители</p>
                </div>
            </section>

            <section className='section content__section'>
                <div className='list-header'>
                    <div className='list-header__toggle'>
                        <NavLink to='#' className='list-header__toggle-link list-header__toggle-link_active'>Хотим сделать</NavLink>
                        <label className='list-header__switch'>
                            <input type='checkbox' className='list-header__checkbox' />
                            <span className='list-header__slider'></span>
                        </label>
                        <NavLink to='#' className='list-header__toggle-link'>Сделали</NavLink>
                    </div>
                    <NavLink to='/add-plan' className='list-header__link'>&#43; Новый план</NavLink>
                </div>
                <ul className='cards page__cards'>
                    <Card />
                    <Card />
                </ul>
            </section>
        </main>
    );
}

export default Main;