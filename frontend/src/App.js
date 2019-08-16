import React from "react";
import { useState, useEffect } from "react";
import firebase from "./firebase";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Index from "./pages/index";

const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

const App = () => {
  const [user, setUser] = useState(null);
  const [message, setMessage] = useState("");
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

  const getPrivateMessage = function(idToken) {
    return fetch(`${endpoint}/private`, {
      method: "get",
      headers: new Headers({
        Authorization: `Bearer ${idToken}`
      }),
      credentials: "same-origin"
    }).then(res => {
      if (res.ok) {
        return res.json();
      } else {
        throw Error(`Request rejected with status ${res.status}`);
      }
    });
  };


  const getPrivateMess = () => {
    user
      .getIdToken()
      .then(token => {
        return getPrivateMessage(token);
      })
      .then(resp => {
        setMessage(resp.message);
      })
      .catch(e => {
        setErrorMessage(e.toString());
      });
  };
  return (
    <div>
      <section className="hero is-dark">
        <div className="hero-head">
          <div className="container">
            <p className="title">CodeHub</p>
          </div>
        </div>
      </section>
      <p>{message}</p>

      <br />
      <Router>
        <Route path="/" exact component={Index} />
      </Router>
    </div>
  );
};

export default App;
