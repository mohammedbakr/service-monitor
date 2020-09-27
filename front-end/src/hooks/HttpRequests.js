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

    const reloadUrl = setInterval(() => {
      setRequest({
        loading: true,
        data: null,
        error: false
      })
      axios.get(url)
      .then(response => {
        setRequest({
          loading: false,
          data: response.data,
          error: false
        })
      })
      .catch(error => {
        console.log(error)
        setRequest({
          loading: false,
          data: null,
          error: true
        })
      })
    }, 30000)

    return () => clearInterval(reloadUrl)
  }, [url])

  return request
}