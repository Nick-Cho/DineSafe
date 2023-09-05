import React, { useState } from 'react'
import env from "react-dotenv";
import axios from "axios"

import Login from "./forms/Login"
import Signup from "./forms/Signup"

type Props = {
  showLogin: boolean;
  showSignup: boolean;
}

function AuthForm({ showLogin, showSignup }: Props) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPswd, setConfirmPswd] = useState("");

  const handleLogin = async (e: React.SyntheticEvent<EventTarget>) => {
    e.preventDefault();
    const requestBody = JSON.stringify(
      {
        email: username,
        password: password
      }
    )

    console.log("request body for login request: ", requestBody);
    try {
      const response = await axios.post(`${env.API_URL}/login`, requestBody);
      console.log(response);
      if (response.status === 200) {
        //successful login, save user data in the redux store

      }
      if (response && response.status === 400) {
        console.log("Failed login: ", response.data);
      }
    } catch (err: any) {
      console.log("Login failed: ", err.response.data);
    }
  }
  return (
    <>
      <div className="bg-white w-full h-screen flex justify-center items-center overflow-hidden">
        <div className="block text-center">
          <h1 className="font-uber font-medium text-2xl">
            Welcome to Placeholder
          </h1>

          <form onSubmit={handleLogin}>
            <div className="w-80">
              <input
                type="text"
                className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                placeholder="Enter phone number or email"
                onChange={(e) => { setUsername(e.target.value) }}
              />
              <input
                type="password"
                className="w-full bg-gray rounded-lg mt-4 py-3 px-4"
                placeholder="Enter password"
                onChange={(e) => { setPassword(e.target.value) }}
              />
              {
                showSignup &&
                <input 
                  type="password"
                  className="w-full bg-gray rounded-lg mt-4 py-3 px-4"
                  placeholder="Re-enter password"
                  onChange={(e) => { setConfirmPswd(e.target.value) }}
                />
              }
              <button className="bg-black text-white w-full rounded-lg mt-4 py-3 ">
                {showSignup ? 
                "Sign up" : "Log in"}
              </button>
            </div>
          </form>
          {showLogin && 
            <h1 className="font-uber underline font-medium cursor-pointer mt-4">
              Forgot your password?
            </h1>
          }
        </div>
      </div>
    </>
  )
}

export default AuthForm