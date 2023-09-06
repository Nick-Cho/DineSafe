import React, { useState } from 'react'
import env from "react-dotenv";
import axios from "axios"

type Props = {
  showLogin: boolean;
  showSignup: boolean;
  setShowLogin: React.Dispatch<React.SetStateAction<boolean>>;
  setShowSignup: React.Dispatch<React.SetStateAction<boolean>>;
}

function AuthForm({ showLogin, showSignup, setShowLogin, setShowSignup }: Props) {
  const [name, setName] = useState("");
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
    setShowLogin(false);
  }

  const handleSignup = async (e: React.SyntheticEvent<EventTarget>) => {
    e.preventDefault();
    const requestBody = JSON.stringify(
      {
        name: name,
        email: username,
        password: password
      }
    )

    console.log("request body from signup request: ", requestBody)
    try {
      const response = await axios.post(`${env.API_URL}/registerUser`, requestBody)
      if (response.status === 200) {
        console.log(response);
      }

      if (response && response.status === 400) {
        console.log("Signup failed: ", response.data);
      }
    } catch (err: any) {
      console.log("Signup failed: ", err.response.data);
    }
    setShowSignup(false);
  }
  return (
    <>
      <div className="bg-white w-full h-screen flex justify-center items-center overflow-hidden">
        <div className="block text-center">
          <h1 className="font-uber font-medium text-2xl">
            Welcome to Placeholder
          </h1>
          <form onSubmit={showLogin ? handleLogin : handleSignup}>
            <div className="w-80">
              {showSignup &&
                <input
                  type="text"
                  className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                  placeholder="Enter your name"
                  onChange={(e) => { setName(e.target.value) }}
                />
              }
              <input
                type="text"
                className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                placeholder="Enter phone number or email"
                onChange={(e) => { setUsername(e.target.value) }}
              />
              <input
                type="password"
                className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                placeholder="Enter password"
                onChange={(e) => { setPassword(e.target.value) }}
              />
              {showSignup &&
                <input
                  type="password"
                  className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                  placeholder="Re-enter password"
                  onChange={(e) => { setConfirmPswd(e.target.value) }}
                />
              }
              <button
                className={`${(showSignup && ((password !== confirmPswd && password !== "") || name === "" || username === "")) ? "bg-btn-gray cursor-not-allowed" : "bg-black cursor-pointer"} 
                  text-white w-full rounded-lg mt-4 py-3`}
                disabled={showSignup && ((password !== confirmPswd && password !== "") || name === "" || username === "")}
              >
                {showSignup ?
                  "Sign up" : "Log in"
                }
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