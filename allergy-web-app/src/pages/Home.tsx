import React, {useState} from 'react'
import homeBanner from '../assets/images/home_banner.jpg'
function Home() {
    const [location, setLocation] = useState("");

    const handleLocation = (e: any) => {
        
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
                            onChange={(e)=>{handleLocation(e)}}
                        />
                        <button 
                            className="bg-black cursor-pointer text-white font-medium rounded-lg mt-4 py-3 px-10"
                        >
                            Search
                        </button>
                    </form>
                </div>
            </div>
            <div className="h-screen w-full flex justify-center items-center col-span-8">
                <img src={homeBanner} alt="home-banner" className="w-full"/> 
            </div>
        </div>
    )
}

export default Home