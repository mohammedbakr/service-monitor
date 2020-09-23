import axios from 'axios'
import React, { useState } from 'react'

function AddConfigurations(props) {
  const [url, setUrl] = useState('')
  // const [increaseUrl, setIncreaseUrl] = useState([])
  // let input = 1

  const onChange = (e) => {
    setUrl(e.target.value)
  }

  // const increaseUrlHandler = () => {
  //   // console.log('clicked')
  //     console.log('clicked')
  //     input = 2
  //     console.log(input)
  //     return input
  // }

  const onSubmit = (e) => {
    e.preventDefault()
    if (url !== '') {

      console.log(url)

      axios.post('http://localhost:10000/api/configurations', {
        Test: url
      })
        .then(res => console.log(res))
        .catch(err => console.log(err))

      setUrl('')
      props.history.push('/')
    }
  }

  return (
    <div>
      <h1 className="block text-gray-700 text-xl font-bold">
        Add Your Own Configurations
      </h1>
      <form onSubmit={onSubmit} className="w-full max-w-lg mt-3 px-4">

        <div className="mt-4">
          <label htmlFor="url" className="block text-gray-700 text-sm font-bold">Url:</label>
          <input
            type="text"
            id="url"
            className="mt-2 shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            placeholder="Enter Url"
            name="url"
            value={url}
            onChange={onChange}
          />
        </div>
        {/* {input} */}

        {/* {increaseUrl} */}
        {/* <div className="w-1/6 flex justify-center items-center m-1 font-medium py-1 px-2 bg-white rounded-full text-blue-700 bg-blue-100 border border-blue-300">
          <div
            className="text-xs font-normal leading-none max-w-full flex-initial"
            style={{cursor: 'pointer'}}
            onClick={increaseUrlHandler}
          >
            (+)Url
          </div>
        </div> */}

        <div className="mt-6">
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          >
            Submit Configurations
          </button>
        </div>

      </form>
    </div>
  )
}

export default AddConfigurations
