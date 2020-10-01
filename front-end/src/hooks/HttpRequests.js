import {useEffect, useState} from 'react'
import axios from 'axios'

// Custom Hook
export function useAxiosGet(url){
  const [request, setRequest] = useState({
      loading: false,
      data: null,
      error: false,
  })

  const backendCall = (url) => {
    setRequest({
      loading: true,
      data: null,
      error: false,
    })
    axios.get(url)
      .then(response => {
        setRequest({
          loading: false,
          data: response.data,
          error: false,
        })
      })
      .catch(error => {
        setRequest({
          loading: false,
          data: null,
          error: true,
        })
      })
  }

  useEffect(() => {
    backendCall(url)

    const reloadUrl = setInterval(() => {
      backendCall(url)
    }, 30000)

    return () => {
      clearInterval(reloadUrl)
    }
  }, [url])

  return request
}