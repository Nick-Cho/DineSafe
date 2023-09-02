import React from 'react'
import { Link } from 'react-router-dom'
function Navbar() {
  return (
    <nav>
      <div className="bg-black text-white font-uber font-medium text-l flex lg:justify-center">
        <div className="ml-2 sm:ml-10 flex md:space-x-4">
          <Link to="/">
            <h1 className="text-2xl font-light py-4">
              Placeholder
            </h1>
          </Link>

          <Link to="/about">
            <div className="my-3 py-2 px-3 invisible lg:visible hover:bg-btn-gray rounded-full ease-in-out duration-300">
              <h1>
                About
              </h1>
            </div>
          </Link>
        </div>
        <div className="flex absolute right-0 mr-2 sm:mr-10 lg:static lg:ml-50% " >
          <div className="my-3 py-2 px-3 rounded-full hover:bg-btn-gray cursor-pointer ease-in-out duration-300">
            <h1>
              Login
            </h1>
          </div>

          <div className="my-3 py-2 px-3 ml-2 rounded-full bg-white cursor-pointer font-small">
            <h1 className="text-black ">
              Sign up
            </h1>
          </div>
        </div>
      </div>
    </nav >
  )
}

export default Navbar