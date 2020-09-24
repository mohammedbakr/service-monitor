import React from 'react'

function Result(props) {
  return (
    <React.Fragment>
      <td className="border px-4 py-2">{props.result.id}</td>
      <td className="border px-4 py-2">{props.result.title}</td>
    </React.Fragment>
  )
}

export default Result
