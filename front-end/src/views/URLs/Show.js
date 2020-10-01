import React from 'react'
import { Link } from 'react-router-dom'
import Loader from '../../components/Loader'
import { useAxiosGet } from '../../hooks/HttpRequests'

function Show(props) {
  const id = props.match.params.id
  const results = useAxiosGet(`/urls/${id}`)

  let loader = null
  let content = null
  let error = null
  let currentUrl = null

  if (results.loading) {
    loader = <Loader></Loader>
  }

  if (results.error) {
    error = (
      <div className="bg-red-300 p-3 text-center">
        <p>Something went wrong...!</p>
      </div>
    )
  }

  if (results.data) {
    currentUrl = results.data.pop().url
    content =
      results.data.map((result, index) => 
        <tr key={index}>
          <td className="border py-2">{index + 1}</td>
          <td className="border py-2">{result.timeresponse}</td>
          <td className="border py-2">{result.time}</td>
          {(result.code >= 200 && result.code < 400) ? <td className="border py-2 bg-green-500">{result.code}</td> : <td className="border py-2 bg-red-500">{result.code}</td>}
        </tr>
      )
  }

  return (
    <div>
      {loader}
      {error}
      <div style={{display: content ? '' : 'none'}}>
        <div className="flex justify-between">
          <h1 className="text-xl">
            <span className="font-bold">Results of: </span>
            <span className="text-blue-500">{currentUrl}</span>
          </h1>
          <Link to={`/urls/${id}/edit`}>
            <button className="bg-blue-500 rounded-lg hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Edit URL
            </button>
          </Link>
        </div>
        <table className="table-fixed w-full mt-6 text-center">
          <thead>
            <tr>
              <th className="border w-1/12 py-2">#</th>
              <th className="border w-3/12 py-2">Response Time</th>
              <th className="border w-6/12 py-2">Timestamp</th>
              <th className="border w-2/12 py-2">Status</th>
            </tr>
          </thead>
          <tbody>
            {content}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Show
