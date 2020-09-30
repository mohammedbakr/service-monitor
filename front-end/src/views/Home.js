import React from 'react'
import Loader from '../components/Loader'
import Result from '../components/Result'
import BarChart from '../components/BarChart'
import { useAxiosGet } from '../hooks/HttpRequests'

function Home() {
  const url = 'http://localhost:10000/api'
  const results = useAxiosGet(url)

  let loader = null
  let content = null
  let error = null
  let chartContent = {}

  if (results.loading) {
    loader = <Loader></Loader>
  }

  if (results.error) {
    error = <p className="mt-3">Something went wrong...!</p>
  }

  if (results.data) {
    content =
      results.data.map((result) => 
        <tr key={result.id}>
            <Result result={result}/>
        </tr>
      )
    
    chartContent = {
      labels: results.data.map(result => result.url),
      datasets: [
        {
          label: 'Time Response (ms)',
          data: results.data.map(result => parseFloat(result.timeresponse)),
          borderColor: results.data.map(() => 'rgba(54, 162, 235, 0.2)'),
          backgroundColor: results.data.map((result) => result.code === 200 ? 'rgba(0,128, 0, 0.5)' : 'rgba(225, 0, 0, .5)')
        }
      ]
    }
  }

  return (
    <div className="text-center">
      <div className="border-b">
        <h1 className="font-bold text-2xl">Health Check</h1>
      </div>
      {loader}
      {error}
      <div style={{display: content ? '' : 'none'}}>
        <div className="mt-6">
          <div className="font-bold text-2xl border-b w-1/3 m-auto">
            Results:
          </div>
          <table className="table-auto w-full mt-6">
            <thead>
              <tr>
                <th className="px-4 py-2">ID</th>
                <th className="px-4 py-2">URL</th>
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
        <div className="mt-3">
          <BarChart chartData={chartContent} />
        </div>
      </div>
    </div>
  )
}

export default Home
