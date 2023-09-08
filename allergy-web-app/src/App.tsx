import React, {useEffect} from 'react';
import Navbar from './components/Navbar'
import Home from './pages/Home'

import { Route, Routes } from 'react-router-dom';
import './App.css';

function App() {
  useEffect(() => {
    if ("geolocation" in navigator) {
      // User already enabled geolocation
      navigator.geolocation.getCurrentPosition(function(position) {
        const {latitude, longitude} = position.coords;
        console.log("Latitude is :", latitude);
        console.log("Longitude is :", longitude);
      });
    } else {
      console.log("Geolocation is not enabled on this browser");
    }
  },[])
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
