import React from 'react'
import Loader from '../../components/Loader'
import Card from '../../components/Card'
import BarChart from '../../components/BarChart'
import { Link } from 'react-router-dom'
import { useAxiosGet } from '../../hooks/HttpRequests'
import axios from 'axios'

function Index() {
  const url = '/stats'
  const results = useAxiosGet(url)

  let loader = null
  let cardContent = null
  let chartContent = {}
  let error = null

  const deleteUrl = (id) => {
    alert('The URL has been deleted successfully')
    console.log(id)
    axios.delete(`/urls/${id}`)
      .then()
      .catch()
  }

  if (results.loading) {
    loader = <Loader></Loader>
  }

  if (results.error) {
    error = (
      <div className="bg-red-300 p-3 mt-3 text-center">
        <p>Something went wrong...!</p>
      </div>
    )
  }

  if (results.data) {
    cardContent =
    results.data.map(result => 
      <Card result={result} key={result.id} deleteUrl={deleteUrl}/>
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
    <React.Fragment>
      <div className="font-bold text-2xl border-b w-1/2 text-center m-auto">
        Health Check
      </div>
      {loader}
      {error}
      <div style={{display: results.data ? '' : 'none'}} className="mt-6">
        <div className="flex justify-around border-b">
          <h1 className="font-bold text-2xl">Result</h1>
          <Link to="/create">
            <button className="bg-blue-500 rounded-lg hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Add URLs
            </button>
          </Link>
        </div>
        <div className="container my-12 mx-auto px-1 md:px-1">
          <div className="flex flex-wrap -mx-1 lg:-mx-1">
            {cardContent}
          </div>
        </div>
        <hr/>
        <div className="mt-3">
          <BarChart chartData={chartContent} />
        </div>
      </div>
    </React.Fragment>
  )
}

export default Index
