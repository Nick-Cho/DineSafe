import React from 'react'

function AutoComplete() {
  return (
    <div className="w-96 lg:w-full relative">
        <div className="absolute w-full text-center bg-white rounded-lg py-10 drop-shadow-lg">
            <h1 className="text-btn-gray">
                No Results Found
            </h1>
        </div>
    </div>
  )
}

export default AutoComplete