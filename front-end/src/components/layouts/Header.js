import React from 'react'
import Nav from './Nav'
import { Link } from "react-router-dom";

function Header() {
  return (
    <header className="border-b p-3 flex justify-between items-center shadow">
      <div className="container m-auto flex justify-between">
        <span className="font-bold">
          <Link 
            to='/'
          >
            Service Monitor
          </Link>
        </span>
        <Nav />
      </div>
    </header>
  )
}

export default Header
