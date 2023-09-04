import React from 'react'

type Props = {
  showLogin: boolean;
}

function Login({ showLogin }: Props) {
  return (
    <>
      <div className="bg-white w-full h-screen flex justify-center items-center overflow-hidden">
        <div className="block text-center">
          <h1 className="font-uber text-3xl">
            Welcome to Placeholder
          </h1>

          <form>
            <div>
              <input
                type="text"
                className="w-full bg-gray focus:border-black rounded-lg mt-4 py-3 px-4"
                placeholder="Enter phone number or email"
              />
              <input
                type="password"
                className="w-full bg-gray rounded-lg mt-4 py-3 px-4"
                placeholder="Enter password"
              />
              <button className="bg-black text-white w-full rounded-lg mt-4 py-3 ">
                Continue
              </button>
            </div>
          </form>
          <h1 className="font-uber underline font-medium cursor-pointer mt-4"> Forgot your password?</h1>
        </div>
      </div>
    </>
  )
}

export default Login