import React, { useState, useEffect } from 'react'
import axios from "axios"
import env from "react-dotenv";
import { useCookies } from "react-cookie";
import { useDispatch, useSelector } from "react-redux";

import { getUserLat, getUserLong } from "../redux/reducers/appReducer";
import AutoComplete from "../components/AutoComplete";
import { AppDispatch } from "../redux/store";
import homeBanner from '../assets/images/home_banner.jpg'

export interface State {
    search_results: {
        name: string, 
        formatted_address: string, 
        rating: number, 
        opening_hours: {open_now: boolean}
    }[]
}
function Home() {
    const [searchResults, setSearchResults] = useState<State["search_results"]>([]);
    const [cookies, setCookies] = useCookies(["latitude", "longitude"]);
    const [showAutoComplete, setShowAutoComplete] = useState<boolean>(false);
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
        // eslint-disable-next-line
    }, [latitude, longitude , cookies.latitude, cookies.longitude]);

    const handleSearch = async (e: any) => {
        // setSearch(e.target.value);
        // console.log("Longitude, latitude: ", longitude, latitude);
        // console.log("search request: ", search);
        setShowAutoComplete(true);
        var search = e.target.value;
        if (search[search.length-1] === " ") {
            search = search.slice(0, search.length-1);
        } 
        search = search.replace(/\s/g, "%20");
        
        
        try {
            const response = await axios.get(`${env.API_URL}/searchRestaurant?search=${search}&longitude=${longitude}&latitude=${latitude}`);
            // console.log("search request string: ", search);
            // console.log("search restaurant response: ", response);
            
            if (response.status === 202) {
                setSearchResults(response.data.candidates);
                // console.log("search results: ", searchResults);
            }
        } catch (err: any) {
            if (err?.response?.status === 400) {
                setSearchResults([]);
            }
            console.log("searchRestaurant request failed: ", err);
        }
    }

    return (
        <div className="grid grid-cols-16 gap-2">
            <div className="col-span-10 lg:col-span-8">
                <div className="bg-white w-screen md:w-full lg:w-auto lg:flex justify-center items-center mx-10 md:mx-40 lg:mx-36 mt-10 md:mt-36 ">
                    <div className="block">
                        <h1 className="font-uber text-btn-gray font-bold text-4xl md:text-5xl leading-snug">
                            Allowing you to enjoy your food worry free
                        </h1>

                        <h1 className="font-uber text-btn-gray text-md mt-3">
                            Search a restaurant, check for allergies, and go.
                        </h1>

                        <form>
                            <div className = "w-100">
                                <input
                                    type="text"
                                    className="w-96 xl:w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                                    placeholder="Where do you want to eat?"
                                    onChange={(e) => { handleSearch(e) }}
                                    onClick={() => { setShowAutoComplete(!showAutoComplete) }}
                                />
                                {showAutoComplete && <AutoComplete 
                                searchResults={searchResults} 
                                />}
                            </div>
                            <button
                                className="bg-black cursor-pointer text-white font-medium rounded-lg mt-4 py-3 px-10"
                            >
                                Search
                            </button>
                        </form>
                    </div>
                </div>
            </div>
            <div className="col-span-7 lg:col-span-8">
                <div className="lg:mt-36 w-screen lg:w-full lg:flex justify-center items-center">
                    <img src={homeBanner} alt="home-banner" className="w-full" />
                </div>
            </div>
        </div>
    )
}

export default Home