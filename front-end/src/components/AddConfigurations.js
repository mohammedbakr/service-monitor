import React, { Component } from 'react'

class AddConfigurations extends Component {
  state = {
    url: ''
  }

  onChange = (e) => {
    this.setState({
      [e.target.name]: e.target.value
    })
  }

  

  render() {
    return (
      <div>
        <h1 className="block text-gray-700 text-xl font-bold">
          Add Your Own Configurations
        </h1>
        <form onSubmit={this.onSubmit} className="w-full max-w-lg mt-3 px-4">
          <div className="mt-4">
            <label htmlFor="url" className="block text-gray-700 text-sm font-bold">Url:</label>
            <input
              type="text"
              className="mt-2 shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Enter Url"
              name="url"
              value={this.state.url}
              onChange={this.onChange}
            />
          </div>
          <div className="mt-6">
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            >Add</button>
          </div>
        </form>
      </div>
    )
  }
}

export default AddConfigurations
