import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { useCookies } from "react-cookie";
import { useDispatch, useSelector } from "react-redux";
import env from "react-dotenv";
import axios from "axios"

import { getUserLat, getUserLong } from "../redux/reducers/appReducer";
import { AppDispatch } from "../redux/store";
function RestaurantInfo() {
  const [cookies, setCookies] = useCookies(["latitude", "longitude"])
  const [address, setAddress] = useState<string>("");
  const [rating, setRating] = useState<number>(0);
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const params = useParams();
  const dispatch: AppDispatch = useDispatch();
  const { latitude, longitude } = useSelector((state: any) => {
    return {
      latitude: getUserLat(state),
      longitude: getUserLong(state)
    };
  });
  useEffect(() => {
    if (
      cookies?.latitude === "" ||
      cookies?.longitude === "" ||
      cookies?.latitude === undefined ||
      cookies?.longitude === undefined ||
      latitude === "" ||
      longitude === ""
    ) {
      if ("geolocation" in navigator) {
        // User already enabled geolocation
        navigator.geolocation.getCurrentPosition(function (position) {
          const { latitude, longitude } = position.coords;
          setCookies("latitude", latitude, { path: "/" });
          setCookies("longitude", longitude, { path: "/" });
          dispatch({
            type: "app/setUserLat",
            payload: cookies.latitude
          });
          dispatch({
            type: "app/setUserLong",
            payload: cookies.longitude
          })
        });

      } else {
        console.log("Geolocation is not enabled on this browser");
      }
    }
    // eslint-disable-next-line
  }, [latitude, longitude, cookies.latitude, cookies.longitude]);

  useEffect(() => {
    const getRestaurantInfo = async () => {
      try {
        const response = await axios.get(`${env.API_URL}/searchRestaurant?search=${params.name}&longitude=${cookies.longitude}&latitude=${cookies.latitude}`);
        if (response.status === 202) {
          setAddress(response.data.candidates[0].formatted_address);
          setRating(response.data.candidates[0].rating);
          setIsOpen(response.data.candidates[0].opening_hours.open_now);
        }
        // console.log("search restaurant return: ", response);
      } catch (err: any) {
        console.log("searchRestaurant endpoint error: ", err.response.data);
      }
    }
    getRestaurantInfo();
  }, [])

  return (
    <div className="grid grid-cols-16 mt-16">
      <div className="col-start-4 col-span-10">
        <h1 className="text-black font-uber font-bold text-4xl">
          {params.name}
        </h1>
        <div className="flex font-uber font-medium text-lg">
          <h1 className="text-black">
            Rating:
          </h1>
          <span>&nbsp;</span>
          <h1 className="text-blue">
            {rating}
          </h1>
          <h1 className={`${isOpen ? "text-green" : "text-red"} ml-5`}> {isOpen ? "Open Now!" : "Closed"} </h1>
        </div>
        <h1 className="font-uber font-medium text-sm">
          {address}
        </h1>
      </div>
    </div>
  )
}

export default RestaurantInfo