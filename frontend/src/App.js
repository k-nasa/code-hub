import React from "react";
import './style.css';


import ErrorMessage from "./component/error_message";
import Index from "./pages/index";
import ShowUser from "./pages/show_user";
import ShowCode from "./pages/show_code";
import WriteCode from "./pages/write_code";
import firebase from "./firebase";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { useState, useEffect } from "react";

const App = () => {
  const [user, setUser] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    firebase.auth().onAuthStateChanged(u => {
      if (u) {
        setUser(u);
      } else {
        setUser(null);
      }
    });
  });

  return (
    <div>
      <pre className="prettyprint">
        <code>puts("hello world")</code>
      </pre>
      <Router>
        <HeaderMenue />
        <ErrorMessage
          error={errorMessage}
          handler={() => setErrorMessage("")}
        />

        <br />
        <Switch>
          <Route exact path="/" exact render={() => <Index user={user} />} />
          <Route
            path="/write_code"
            render={() => <WriteCode user={user} />}
          />
          <Route
            path="/user/:id"
            render={p => <ShowUser user={user} p={p} />}
          />
          <Route path="/:username/:title" render={p => <ShowCode p={p} />} />
        </Switch>
      </Router>
    </div>
  );
};

export default App;

const HeaderMenue = () => (
  <section className="hero is-dark">
    <div className="hero-head">
      <div className="container">
        <Link className="title" to="/">
          CodeHub
        </Link>
        {
          <p style={{ textAlign: "right" }} onClick={firebase.logout}>
            logout
          </p>
        }
      </div>
    </div>
  </section>
);
