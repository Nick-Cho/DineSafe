import React, {useState} from 'react'
import env from "react-dotenv";
import axios from "axios"

type Props = {
  showLogin: boolean;
}

function Login({ showLogin }: Props) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e: React.SyntheticEvent<EventTarget>)  => {
    e.preventDefault();

    const response = await axios.post(`${env.API_URL}/login`, {
      email: username,
      password: password
    });

    console.log("Response from login request: ", response);
  }
  return (
    <>
      <div className="bg-white w-full h-screen flex justify-center items-center overflow-hidden">
        <div className="block text-center">
          <h1 className="font-uber text-3xl">
            Welcome to Placeholder
          </h1>

          <form onSubmit={handleSubmit}>
            <div>
              <input
                type="text"
                className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                placeholder="Enter phone number or email"
                onChange={(e)=>{setUsername(e.target.value)}}
              />
              <input
                type="password"
                className="w-full bg-gray rounded-lg mt-4 py-3 px-4"
                placeholder="Enter password"
                onChange={(e)=>{setPassword(e.target.value)}}
              />
              <button className="bg-black text-white w-full rounded-lg mt-4 py-3 ">
                Continue
              </button>
            </div>
          </form>
          <h1 
            className="font-uber underline font-medium cursor-pointer mt-4"
          > 
            Forgot your password?
            </h1>
        </div>
      </div>
    </>
  )
}

export default Login