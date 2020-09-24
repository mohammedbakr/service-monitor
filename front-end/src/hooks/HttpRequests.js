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
      loading: true,
      data: null,
      error: false
    })
    axios.get(url)
      .then(response => {
        // console.log(response.data)
        setRequest({
          loading: false,
          data: response.data,
          error: false
        })
      })
      .catch((error) => {
        console.log(error)
        setRequest({
          loading: false,
          data: null,
          error: true
        })
      })
  }, [url])

  return request
}