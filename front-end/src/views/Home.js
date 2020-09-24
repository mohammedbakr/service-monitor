import React from 'react'
import Loader from '../components/Loader'
import Result from '../components/Result'
import BarChart from '../components/BarChart'
import { Link } from 'react-router-dom'
import { useAxiosGet } from '../hooks/HttpRequests'

function Home() {
  // const url = 'http://localhost:10000/api'
  const url = 'https://jsonplaceholder.typicode.com/todos?_limit=10'
  const results = useAxiosGet(url)

  let loader = null
  let content = null
  let chartContent = {}

  if (results.loading) {
    loader = <Loader></Loader>
  }

  if (results.data) {
    content =
      results.data.map((result, index) => 
        <tr key={index}>
            <Result result={result}/>
        </tr>
      )
    
    chartContent = {
      labels: results.data.map(result => result.title),
      datasets: [
        {
          label: 'Time Response (ms)',
          data: results.data.map(result => result.id),
          borderColor: results.data.map(() => 'rgba(54, 162, 235, 0.2)'),
          backgroundColor: results.data.map(() => 'rgba(54, 162, 235, 0.2)')
        }
      ]
    }
  }

  return (
    <div className="text-center">
      <div className="border-b">
        <h1 className="font-bold text-2xl">Health Check</h1>
        <Link
          to="/add-configurations"
          className="text-blue-500 py-3 block"
        >
          Add your own configurations
        </Link>
      </div>
      {loader}
      <div style={{display: content ? '' : 'none'}}>
        <div className="mt-6">
          <div className="font-bold text-2xl border-b w-1/3 m-auto">
            Results:
          </div>
          <table className="table-auto w-full mt-6">
            <thead>
              <tr>
                <th className="px-4 py-2">Id</th>
                <th className="px-4 py-2">Title</th>
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
