import React, { useEffect, useState } from 'react'
import axios from 'axios'
import Loader from '../../components/Loader';

function Edit(props) {
  const id = props.match.params.id
  const [url, setUrl] = useState('');
  const [timeInterval, setTimeInterval] = useState(5)
  const [loader, setLoader] = useState(<Loader></Loader>)
  const [error, setError] = useState({ url: '', timeInterval: '' })
  const urlPath = `/urls/${id}`
  const urlPattern = new RegExp(/(http|https)?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,}/)

  useEffect(() => {
    axios.get(urlPath)
      .then(response => {
        setLoader(null)
        setUrl(response.data.pop().url)
      })
      .catch(error => {
        setLoader(null)
      })
  }, [urlPath])

  const validInputs = () => {
    if (url.trim() === '') {
      setError({...error, url: 'URL is required'})
    } else if (url.trim() !== '' && url.match(urlPattern)) {
      setError({...error, url: ''})
    } else {
      setError({...error, url: 'URL must be like http://www.example.com'})
    }

    if (!timeInterval) {
      setError({...error, timeInterval: 'Time Interval is required'})
    } else if (+timeInterval <= 0) {
      setError({...error, timeInterval: 'Time Interval must be greater than or equal 1'})
    }
  }

  const formIsValid = () => {
    validInputs()
    return url.trim() &&
           url.match(urlPattern) &&
           timeInterval &&
           +timeInterval > 0
  }

  const onSubmit = (e) => {
    e.preventDefault()
    if (!formIsValid) {
      return
    }

    axios.put(urlPath, {
      theurl: url.trim(),
      time: +timeInterval
    })
      .then()
      .catch()

    setUrl('')
    setError({ url: '', timeInterval: '' })
    alert("The configurations have been added successfully")
    props.history.push('/')
  }

  return (
    <div>
      {loader ? loader :
        <div>
          <h1 className="block text-gray-700 text-xl font-bold">
            Edit Configurations
          </h1>
          <form onSubmit={onSubmit} className="w-full max-w-lg mt-3 px-4">
            <div>
              <label className="block text-gray-700 text-sm font-bold">URL:</label>
              <input
                type="text"
                id="url"
                name="url"
                className={`mt-2 shadow appearance-none border rounded w-5/6 py-2 px-3 text-gray-700 leading-tight focus:outline-none ${error.url ? 'border-red-500' : 'focus:shadow-outline'}`}
                onChange={e => setUrl(e.target.value)}
                value={url}
                placeholder="http://www.example.com"
                required
              /><br/>
              <span style={{display: error.url ? '' : 'none'}} className="text-red-600">{error.url}</span>
            </div>
            <div className="mt-3">
              <label htmlFor="timeInterval" className="block text-gray-700 text-sm font-bold">Time Interval:</label>
              <input
                type="number"
                min="1"
                id="timeInterval"
                name="timeInterval"
                className={`mt-2 shadow appearance-none border rounded w-5/6 py-2 px-3 text-gray-700 leading-tight focus:outline-none ${error.timeInterval ? 'border-red-500' : 'focus:shadow-outline'}`}
                onChange={e => setTimeInterval(e.target.value)}
                value={timeInterval}
                placeholder="Enter Time Interval"
                required
              /><br/>
              <span style={{display: error.timeInterval ? '' : 'none'}} className="text-red-600">{error.timeInterval}</span>
            </div>
            <div className="mt-6">
              <button
                type="submit"
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg focus:outline-none focus:shadow-outline"
              >
                update
              </button>
            </div>
          </form>
        </div>
      }
    </div>
  );
};

export default Edit
