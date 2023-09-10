import React from 'react'

function AutoComplete({searchResults}: any) {
  return (
    <div className="w-96 lg:w-full relative">
        <div className={
            `absolute w-full text-center bg-white rounded-lg 
            ${searchResults.length=== 0 ? "py-10" : "py-2"} 
            drop-shadow-lg`}>
        {searchResults.length === 0 &&  
            <h1 className="text-light-gray">
                No Results Found
            </h1>            
        }
        {searchResults.length > 0 && searchResults.map((result: any) => {
            return (
                <div className="hover:bg-gray py-3 duration-200">
                    <h1 className="text-black cursor-pointer">
                        {result.name}
                    </h1>
                </div>
            )
        })

        }
        </div>
    </div>
  )
}

export default AutoComplete