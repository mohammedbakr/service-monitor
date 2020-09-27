import React from 'react'
import { useAxiosGet } from '../hooks/HttpRequests'
import Loader from '../components/Loader'

function UrlDetails(props) {
  const url = 'http://localhost:10000/api'
  // const url = 'https://jsonplaceholder.typicode.com/todos?_limit=10'
  // const results = useAxiosGet(url + props.match.params.id)
  // const id = props.match.params.id
  const results = useAxiosGet(url)

  let loader = null
  let content = null
  let error = null

  if (results.loading) {
    loader = <Loader></Loader>
  }

  if (results.error) {
    error = <p className="mt-3">Something went wrong...!</p>
  }

  if (results.data) {
    content =
      results.data.map((result, index) => 
        <tr key={index}>
          <td className="border px-4 py-2">{result.url}</td>
          <td className="border px-4 py-2">{result.timeresponse}</td>
          <td className="border px-4 py-2">{result.time}</td>
          {(result.code >= 200 && result.code < 400) ? <td className="border px-4 py-2 bg-green-500">{result.code}</td> : <td className="border px-4 py-2 bg-red-500">{result.code}</td>}
        </tr>
      )
  }

  return (
    <div className="text-center">
      {loader}
      {error}
      <div style={{display: content ? '' : 'none'}}>
        <div className="mt-3">
          <div className="font-bold text-2xl border-b w-1/3 m-auto">
            Results:
          </div>
          <table className="table-auto w-full mt-6">
            <thead>
              <tr>
                <th className="px-4 py-2">Url</th>
                <th className="px-4 py-2">Response Time</th>
                <th className="px-4 py-2">Timestamp</th>
                <th className="px-4 py-2">Response Code</th>
              </tr>
            </thead>
            <tbody>
              {content}
            </tbody>
          </table>
        </div>
        <hr/>
      </div>
    </div>
  )
}

export default UrlDetails
