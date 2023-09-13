import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { useCookies } from "react-cookie";
import { useDispatch, useSelector } from "react-redux";
import env from "react-dotenv";
import axios from "axios"

import { getUserLat, getUserLong } from "../redux/reducers/appReducer";
import { AppDispatch } from "../redux/store";
import AddReviewForm from "../components/Forms/AddReviewForm"
function RestaurantInfo() {
  const [cookies, setCookies] = useCookies(["latitude", "longitude"])
  //Variables for restaurant info
  const [address, setAddress] = useState<string>("");
  const [rating, setRating] = useState<number>(0);
  const [isOpen, setIsOpen] = useState<boolean>(false); // Tracking if the restaurant is currently open
  const params = useParams();

  const [reviewsLoaded, setReviewsLoaded] = useState<boolean>(false); // Tracks if the reviews have been loaded from the database
  const [reviews, setReviews] = useState<string[]>([]);
  const [addReview, setAddReview] = useState<boolean>(false); // Tracks if the user wants to bring up the add review ui

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
      if (env?.API_URL) {
        try {
          const searchResponse = await axios.get(`${env.API_URL}/searchRestaurant?search=${params.name}&longitude=${cookies.longitude}&latitude=${cookies.latitude}`);
          if (searchResponse.status === 202) {
            setAddress(searchResponse.data.candidates[0].formatted_address);
            setRating(searchResponse.data.candidates[0].rating);
            setIsOpen(searchResponse.data.candidates[0].opening_hours?.open_now);
          }
          // console.log(address);
          const getReviewsResp = await axios.get(`${env.API_URL}/getRestaurantReviews?street_address=${encodeURIComponent(searchResponse.data.candidates[0].formatted_address.trim())}`);
          // console.log("Get Reviews response: ", getReviewsResp)
          if (getReviewsResp.status === 202 && !reviewsLoaded) {
            // successful retrieval
            // console.log("reviews loaded: ", reviewsLoaded);
            // console.log("GetReviews return where reviews exist: ", getReviewsResp.data);
            const reviews: string[] = [];
            getReviewsResp.data.forEach((review: any) => {
              reviews.push(review.review);
            })
            setReviewsLoaded(true);
            setReviews(reviews);
          }
          else if (getReviewsResp.status === 201) {
            // status code for restaurant not yet being inserted into the database
            const requestBody = JSON.stringify({
              street_address: searchResponse.data.candidates[0].formatted_address,
              name: params.name
            })
            try {
              const insertResp = await axios.post(`${env.API_URL}/insertRestaurant`, requestBody);
              if (insertResp.status === 202) {
                // successful insertion
                setReviews([]);
              }
            } catch (err: any) {
              console.log("Insert restaurant error: ", err?.message);
            }
          }
          // console.log("getRestaurantReviews response: ", getReviewsResp)
          // console.log("search restaurant return: ", searchResponse);}
        } catch (err: any) {
          console.log("searchRestaurant endpoint error: ", err);
        }
      }
    }
    getRestaurantInfo();
  }, [])

  return (
    <div className="grid grid-cols-16 mt-16">
      <div className="col-start-4 col-span-10 rounded-lg px-4 py-6 bg-gray">
        <>
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
        </>

        <div className="mt-10">
          <h1 className="font-uber font-bold text-3xl">
            Items to Watch Out For:
          </h1>
          <div className="font-uber">
            {reviews.length === 0 && 
              <div className="py-4 px-6 mt-3 bg-white rounded-lg"> 
                <h1 className="font-uber font-medium text-lg my-3">No allergies recorded yet, be the first!</h1>
                <button 
                  className="bg-black text-white font-uber rounded-lg font-medium px-2 py-1 cursor-pointer"
                  onClick={()=>{setAddReview(true)}}
                >
                  Record an allergy
                </button>
              </div>
            }
            {reviews.map((review, index) => {
              return (
                <div key={index} className="py-4 px-6 mt-3 bg-white rounded-lg">
                  <h1 className="font-uber font-medium text-lg my-3">
                    {review}
                  </h1>
                </div>
              )
            })}
          </div>
        </div>
        {addReview &&
          <div className="w-full flex">
            <AddReviewForm address={address} setAddReview={setAddReview}/>
          </div>
        }
      </div>
    </div>
  )
}

export default RestaurantInfo