import React, { useState } from 'react'
import axios from 'axios'

function Create(props) {
  const [url, setUrl] = useState('');
  const [error, setError] = useState(null)
  const urlPattern = new RegExp(/(http|https)?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,}/)

  const onChange = (e) => {
    setUrl(e.target.value)
  }

  const onSubmit = (e) => {
    e.preventDefault()
    if (url.trim() !== '') {
      if (url.match(urlPattern)) {
        setError(null)
        axios.post('/urls', {
          test: url.trim(),
        })
          .then()
          .catch()
  
        setUrl('')
        props.history.push('/')
        alert("The URL has been added successfully")
      } else {
        setError('URL must be like http://www.example.com')
      }
    } else {
      setError('URL must be like http://www.example.com')
    }
  }

  return (
    <div>
      <h1 className="block text-gray-700 text-xl font-bold">
        Add Your Own Configurations
      </h1>
      <form onSubmit={onSubmit} className="w-full max-w-lg mt-3 px-4">
        <div>
          <label htmlFor="url" className="block text-gray-700 text-sm font-bold">Url:</label>
          <input
            className={`mt-2 shadow appearance-none border rounded w-5/6 py-2 px-3 text-gray-700 leading-tight focus:outline-none ${error ? 'border-red-500' : 'focus:shadow-outline'}`}
            onChange={onChange}
            value={url}
            placeholder="http://www.example.com"
          /><br/>
          <span style={{display: error ? '' : 'none'}} className="text-red-600">{error}</span>
        </div>
        <div className="mt-6">
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg focus:outline-none focus:shadow-outline"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default Create
