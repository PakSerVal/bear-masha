import React from 'react';
import { Route, Routes } from 'react-router-dom';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';
import Login from './Login';
import AddCard from './AddCard';
import EditProfile from './EditProfile';

function App() {
  return (
      <div className='page'>
          <div className='page__container'>
              <Header />
              <Routes>
                  <Route path='/' element={<Main />} />
                  <Route path='/sign-in' element={<Login />} />
                  <Route path='/add-plan' element={<AddCard />} />
                  <Route path='/edit-profile' element={<EditProfile />} />
              </Routes>
              <Footer />
          </div>
      </div>
  );
}

export default App;
