import React from 'react';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Header from './components/layouts/Header';
import Index from './views/URLs/Index';
import Create from './views/URLs/Create';
import Show from './views/URLs/Show';
import Edit from './views/URLs/Edit';
import About from './views/About';
import Footer from './components/layouts/Footer';

function App() {
  return (
    <div className="relative pb-10 min-h-screen">
      <Router>
        <Header />
        <div className="container m-auto p-3 mt-3">
          <Switch>
            <Route exact path="/" component={Index} />
            <Route path="/create" component={Create} />
            <Route exact path="/urls/:id" component={Show} />
            <Route path="/urls/:id/edit" component={Edit} />
            <Route path="/about" component={About} />
          </Switch>
        </div>
        <Footer />
      </Router>
    </div>
  );
}

export default App;
