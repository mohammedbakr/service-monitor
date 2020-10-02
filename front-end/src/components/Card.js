import React from 'react'
import { Link } from 'react-router-dom'

function Card(props) {
  return (
    <React.Fragment>
      <div className="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-1 lg:w-1/3">
        <article className="overflow-hidden rounded-lg shadow-lg">
          <header className="flex items-center justify-between leading-tight p-2 md:p-4">
            <h1 className="text-lg">
              <Link to={`/urls/${props.result.id}`}>
                <span className="no-underline hover:underline text-black">
                  {props.result.url}
                </span>
              </Link>
            </h1>
            <div className="flex">
              {
                (props.result.code >= 200 && props.result.code < 400) ?
                <p className="text-white font-bold bg-green-400 rounded p-2 text-sm">
                  {props.result.code}
                </p> :
                <p className="text-white font-bold bg-red-400 rounded p-2 text-sm">
                  {props.result.code}
                </p>
              }
              <button
                className="text-xs ml-1 bg-red-500 hover:bg-red-700 text-white font-bold px-1 rounded focus:outline-none focus:shadow-outline"
                onClick={() => props.deleteUrl(props.result.id)}
              >
                Delete
              </button>
            </div>
          </header>
          <footer className="flex items-center justify-between leading-none p-2 md:p-4">
            <p className="ml-2 text-sm">
                {props.result.time}
            </p>
            <p className="ml-2 text-sm">
                {props.result.timeresponse}
            </p>
          </footer>
        </article>
      </div>
    </React.Fragment>
  )
}

export default Card
