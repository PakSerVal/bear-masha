import React, { useState, useEffect } from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import ProtectedRoute from './ProtectedRoute';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';
import Login from './Login';
import AddCard from './AddCard';
import InfoToolTip from './InfoToolTip';
import * as auth from '../utils/auth.js';

function App() {
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isInfoTooltipPopupOpen, setIsInfoTooltipPopupOpen] = useState(false);

    useEffect(() => {
        if (isLoggedIn) {
            // api.getInitialsCards()
        }

    })

    function closeAllPopups() {
        setIsInfoTooltipPopupOpen(false);
    }

    function handleAuthResult(isOpen) {
        setIsInfoTooltipPopupOpen(isOpen);
    }

    function handleLogin(login, password) {
        if (!login || !password) {
            return;
        }
        auth.authorize(login, password)
            .then((res) => {
                if (res.token) {
                    localStorage.setItem('jwt', res.token);
                    navigate('/');
                    setIsLoggedIn(true);
                } else {
                    handleAuthResult(true);
                }
            })
            .catch((err) => {
                console.log(err);
                handleAuthResult(true);
            });
    }

    return (
        <div className='page'>
        <div className='page__container'>
                <Header />
                <Routes>
                    <Route path='/sign-in' element={<Login onLogin={handleLogin} />} />
                    <Route
                        path='/'
                        element={
                            <Main />
                            // <ProtectedRoute isLoggedIn={isLoggedIn}>
                            //     <Main />
                            // </ProtectedRoute>
                        }
                    />
                    <Route
                        path="/add-plan"
                        element={
                            <AddCard />
                            // <ProtectedRoute isLoggedIn={isLoggedIn}>
                            //     <AddCard />
                            // </ProtectedRoute>
                        }
                    />
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
