import React, { useState } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';
import Login from './Login';
import AddCard from './AddCard';
import EditProfile from './EditProfile';
import InfoToolTip from './InfoToolTip';
import * as auth from '../utils/auth.js';

function App() {
    const navigate = useNavigate();
    const [isInfoTooltipPopupOpen, setIsInfoTooltipPopupOpen] = useState(false);

    function closeAllPopups() {
        setIsInfoTooltipPopupOpen(false);
    }

    function handleAuthResult() {
        setIsInfoTooltipPopupOpen(true);
    }

    function handleLogin(login, password) {
        if (!login || !password) {
            return;
        }
        auth.authorize(login, password)
            .then((res) => {
                if (res.token) {
                    localStorage.setItem('jwt', res.token);
                    navigate.push("/");
                } else {
                    handleAuthResult(false);
                }
            })
            .catch((err) => {
                console.log(err);
                handleAuthResult(false);
            });
    }

    return (
        <div className='page'>
            <div className='page__container'>
                <Header />
                <Routes>
                    <Route path='/sign-in' element={<Login onLogin={handleLogin} />} />
                    <Route path='/' element={<Main />} />
                    <Route path='/add-plan' element={<AddCard />} />
                    <Route path='/edit-profile' element={<EditProfile />} />
                </Routes>
                <Footer />
                <InfoToolTip
                    isOpen={isInfoTooltipPopupOpen}
                    onClose={closeAllPopups}
                />
            </div>
        </div>
    );
}

export default App;
