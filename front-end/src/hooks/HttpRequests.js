import {useEffect, useState} from 'react'
import axios from 'axios'

// Custom Hook
export function useAxiosGet(url){
  const [request, setRequest] = useState({
      loading: false,
      data: null,
      error: false
  })

  useEffect(() => {
    setRequest({
      loading: false,
      data: null,
      error: false
    })
    // axios.get(url, {
    //   headers: { 'Content-Type': 'application/json' }
    // })
    axios.get(url)
      .then(response => {
        console.log(response)
        // setRequest({
        //   loading: false,
        //   data: response.data,
        //   error: false
        // })
      })
      .catch((error) => {
        console.log(error)
        // setRequest({
        //   loading: false,
        //   data: null,
        //   error: true
        // })
      })
  }, [url])

  return request
}