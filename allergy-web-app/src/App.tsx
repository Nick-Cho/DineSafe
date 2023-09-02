import React from 'react';
import Navbar from './components/Navbar'
import { Route, Routes } from 'react-router-dom';
import './App.css';

function App() {
  return (
    <div>
      <Navbar />
      <div>
        <Routes>
          <Route path="/"></Route>
        </Routes>
      </div>
    </div>
  );
}

export default App;
