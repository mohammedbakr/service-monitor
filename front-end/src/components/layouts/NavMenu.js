import React from 'react'
import { Link } from 'react-router-dom'

function NavMenu(props) {
  return (
    <div>
      <div className="font-bold py-3">
        Service Monitor
      </div>
      <ul>
        <li>
          <Link
            to="/"
            className="text-blue-500 border-t border-b py-3 block"
            onClick={props.closeMenu}
          >
            Home
          </Link>
        </li>
        <li>
          <Link
            to="/add-configurations"
            className="text-blue-500 border-b py-3 block"
            onClick={props.closeMenu}
          >
            Add Configurations
          </Link>
        </li>
        <li>
          <Link
            to="/about"
            className="text-blue-500 border-b py-3 block"
            onClick={props.closeMenu}
          >
            About
          </Link>
        </li>
      </ul>
    </div>
  )
}

export default NavMenu
