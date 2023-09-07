import React, { useState } from 'react'
import { Link } from 'react-router-dom'

import AuthForm from './AuthForm'

function Navbar() {
  const [showLogin, setShowLogin] = useState<boolean>(false);
  const [showSignup, setShowSignup] = useState<boolean>(false);

  const displayForm = async (request: string) => {
    if (request === "login") {
      if (showLogin) setShowLogin(false);
      else {
        setShowLogin(true);
        if (showSignup) setShowSignup(false);
      }
    }
    else if (request === "signUp") {
      if (showSignup) setShowSignup(false)
      else{
        setShowSignup(true);
        if (showLogin) setShowLogin(false);
      }
    }
  }
  return (
    <>
      <nav>
        <div className="sticky top-0 bg-black text-white font-uber font-medium text-l flex lg:justify-center z-50">
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
          <div className="flex absolute right-0 mr-2 sm:mr-10 lg:static lg:ml-50%" >
            <div
              className="my-3 py-2 px-3 rounded-full hover:bg-btn-gray cursor-pointer ease-in-out duration-300"
              onClick={() => { displayForm("login")}}
            >
              <h1>
                Login
              </h1>
            </div>

            <div
              className="my-3 py-2 px-3 ml-2 rounded-full bg-white cursor-pointer font-small"
              onClick={() => { displayForm("signUp")}}
            >
              <h1 className="text-black">
                Sign up
              </h1>
            </div>
          </div>
        </div>
      </nav>
      {
      <div className={`absolute overflow-hidden justify-center w-full h-screen ${(showLogin || showSignup) ? "translate-y-0 visible" : "-translate-y-full invisible"}  duration-300`}>
        <AuthForm showLogin={showLogin} showSignup={showSignup} setShowLogin={setShowLogin} setShowSignup={setShowSignup}/>
      </div>}
    </>
  )
}

export default Navbar