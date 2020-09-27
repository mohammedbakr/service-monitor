import React from 'react'
import { Link } from 'react-router-dom'

function Result(props) {
  return (
    <React.Fragment>
      <td className="border px-4 py-2">
        <Link to={`details/${props.result.url}`}>{props.result.url}</Link>
      </td>
      <td className="border px-4 py-2">{props.result.timeresponse}</td>
      <td className="border px-4 py-2">{props.result.time}</td>
      {(props.result.code >= 200 && props.result.code < 400) ? <td className="border px-4 py-2 bg-green-500">{props.result.code}</td> : <td className="border px-4 py-2 bg-red-500">{props.result.code}</td>}
    </React.Fragment>
  )
}

export default Result
