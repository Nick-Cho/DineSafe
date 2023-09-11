import React from 'react';

import Navbar from './components/Navbar'
import Home from './pages/Home'
import RestaurantInfo from './pages/RestaurantInfo';

import { Route, Routes } from 'react-router-dom';

function App() {
  return (
    <div>
      <Navbar/>
      <Routes>
        <Route path="/" element={<Home/>}></Route>
        <Route path="/review/:name" element={<RestaurantInfo/>}></Route>
      </Routes>
      
    </div>
  );
}

export default App;
