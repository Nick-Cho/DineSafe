import React, {useEffect} from 'react'
import {useParams} from 'react-router-dom'
import axios from "axios"
function RestaurantInfo() {
  let params = useParams();
  useEffect(() => {
    
  },[])

  return (
    <div>
      <h1>
      {params.name}        
      </h1>
    </div>
  )
}

export default RestaurantInfo