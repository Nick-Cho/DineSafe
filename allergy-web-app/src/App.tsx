import React, {useEffect} from 'react';

import Navbar from './components/Navbar'
import Home from './pages/Home'

import { Route, Routes } from 'react-router-dom';
import './App.css';

function App() {
  return (
    <div>
      <Navbar />
      
      <Routes>
        <Route path="/" element={<Home/>}></Route>
      </Routes>
      
    </div>
  );
}

export default App;
