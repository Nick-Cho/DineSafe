import React from 'react'
import {Link} from 'react-router-dom'
function AutoComplete({searchResults}: any) {
    // console.log("Auto complete search results: ", searchResults)
  return (
    <div className="w-96 lg:w-full relative z-10">
        <div className={
            `absolute w-full text-center bg-white rounded-lg z-10
            ${searchResults.length === 0 ? "py-10" : "py-2"} 
            drop-shadow-lg`}>
        {searchResults.length === 0 &&  
            <h1 className="text-light-gray">
                No Results Found
            </h1>            
        }
        {searchResults.length > 0 && searchResults.map((result: any) => {
            return (
                <div id={result.formatted_address}>
                    <Link to={`/review/${result?.name?.replace(/\s/g, "%20")}`}>
                        <div className="hover:bg-gray py-3 duration-200">
                            <h1 className="text-black cursor-pointer">
                                {result?.name}
                            </h1>
                        </div>
                    </Link>
                </div>
            )
        })

        }
        </div>
    </div>
  )
}

export default AutoComplete