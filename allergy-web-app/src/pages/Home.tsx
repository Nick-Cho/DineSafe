import React, { useState, useEffect } from 'react'
import axios from "axios"
import env from "react-dotenv";
import { useCookies } from "react-cookie";
import { useDispatch, useSelector } from "react-redux";

import { getUserLat, getUserLong } from "../redux/reducers/appReducer";
import { AppDispatch } from "../redux/store";
import homeBanner from '../assets/images/home_banner.jpg'
function Home() {
    const [search, setSearch] = useState("");
    const [searchResults, setSearchResults] = useState<string[]>([]);
    const [cookies, setCookies] = useCookies(["latitude", "longitude"]);

    const { latitude, longitude } = useSelector((state: any) => {
        return {
            latitude: getUserLat(state),
            longitude: getUserLong(state)
        };
    });
    const dispatch: AppDispatch = useDispatch();
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
    }, []);

    const handleSearch = async (e: any) => {
        setSearch(e.target.value);
        // console.log("Longitude, latitude: ", longitude, latitude);
        console.log("search request: ", search);

        try {
            const response = await axios.get(`${env.API_URL}/searchRestaurant?search=${search}&longitude=${longitude}&latitude=${latitude}`);
            console.log("search restaurant response: ", response);
            if (response.status === 202) {

            }
        } catch (err: any) {
            console.log("Login failed: ", err);
        }

    }

    return (
        <div className="grid grid-cols-16">
            <div className="bg-white h-screen flex justify-center items-center col-span-8 mx-36 -mt-10">
                <div className="block">
                    <h1 className="font-uber text-btn-gray font-bold text-5xl leading-snug">
                        Allowing you to enjoy your food worry free
                    </h1>

                    <h1 className="font-uber text-btn-gray text-md mt-3">
                        Search a restaurant, check for allergies, and go.
                    </h1>

                    <form>
                        <input
                            type="text"
                            className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                            placeholder="Where do you want to eat?"
                            onChange={(e) => { handleSearch(e) }}
                        />
                        <div className="absolute bg-white rounded-lg ">

                        </div>
                        <button
                            className="bg-black cursor-pointer text-white font-medium rounded-lg mt-4 py-3 px-10"
                        >
                            Search
                        </button>
                    </form>
                </div>
            </div>
            <div className="h-screen w-full flex justify-center items-center col-span-8">
                <img src={homeBanner} alt="home-banner" className="w-full" />
            </div>
        </div>
    )
}

export default Home