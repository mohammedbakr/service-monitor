// import axios from 'axios'
import React, { useState } from 'react'
import { produce } from "immer";
import { generate } from "shortid";
// import { useAxiosGet } from '../hooks/HttpRequests';
// import Result from './Result';

function AddConfigurations() {
  // const url = 'https://jsonplaceholder.typicode.com/todos?_limit=2'
  const [urls, setUrls] = useState([
    { id: '1', url: "https://www.google.com" }
  ]);

  // let ids = null
  // let titles = null
  // const results = useAxiosGet(url)
  // if (results.data) {
  //   ids = results.data.map(result => result.id)
  //   titles = results.data.map(result => result.title)
  //   setUrls([
  //     {id: '2', url: 'http://yahoo.com'}
  //   ])
  // }
  // console.log(urls)

  const onSubmit = (e) => {
    e.preventDefault()
    console.log(urls)
    // axios.post('http://localhost:10000/api/configurations', {
    //   // test: JSON.stringify(urls, null, 2)
    //   // test: urls
    //   test: JSON.stringify(urls.map(url => url.url))
    // })
    //   .then(response => console.log(['success', response]))
    //   .catch(error => console.log(['error', error]))
  }

  return (
    <div>
      <h1 className="block text-gray-700 text-xl font-bold">
        Add Your Own Configurations
      </h1>
      <form onSubmit={onSubmit} className="w-full max-w-lg mt-3 px-4">
        <label htmlFor="url" className="block text-gray-700 text-sm font-bold">Url:</label>
        {urls.map((p, index) => {
          return (
            <div key={p.id} className="mt-2">
              <input
                className="mt-2 shadow appearance-none border rounded w-5/6 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                onChange={e => {
                  const url = e.target.value;
                  setUrls(currenturls =>
                    produce(currenturls, v => {
                      v[index].url = url;
                    })
                  );
                }}
                value={p.url}
                placeholder="Enter url"
              />
              <button disabled={p.id === '1'}
                className="text-xs mx-4 bg-red-500 hover:bg-red-700 text-white font-bold px-1 rounded focus:outline-none focus:shadow-outline"
                onClick={() => {
                  setUrls(currenturls =>
                    currenturls.filter(x => x.id !== p.id)
                  );
                }}
              >
                X
              </button>
            </div>
          );
        })}
        <div
          className="w-1/6 flex justify-center items-center m-1 font-medium py-1 px-2 bg-white rounded-full text-blue-700 bg-blue-100 border border-blue-300"
        >
          <div
            className="text-xs font-normal leading-none max-w-full flex-initial"
            style={{cursor: 'pointer'}}
            placeholder="add new url"
            onClick={() => {
              setUrls(currenturls => [
                ...currenturls,
                {
                  id: generate(),
                  url: "",
                }
              ]);
            }}
          >
            (+) url
          </div>
        </div>
        <div className="mt-6">
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          >
            Submit Configurations
          </button>
        </div>
      </form>
      {/* <div>{JSON.stringify(urls, null, 2)}</div> */}
    </div>
  );
};

export default AddConfigurations
