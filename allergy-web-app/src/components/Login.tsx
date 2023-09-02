import React from 'react'
import PropTypes from 'prop-types'

type Props = {
  showLogin: boolean;
}

function Login({ showLogin }: Props) {
  return (
    <>
      <div className="bg-black w-full h-screen">
        Login
      </div>
    </>
  )
}

export default Login