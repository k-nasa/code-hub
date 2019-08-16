import React from "react";
import { useState, useEffect } from "react";
import firebase from "./firebase";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Index from "./pages/index";

const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

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
      <section className="hero is-dark">
        <div className="hero-head">
          <div className="container">
            <p className="title">CodeHub</p>
            <p style={{ textAlign: "right" }} onClick={firebase.logout}>
              logout
            </p>
          </div>
        </div>
      </section>

      {errorMessage ? <errorMessage error={errorMessage} /> : <div />}

      <br />
      <Router>
        <Route path="/" exact render={() => <Index user={user} />} />
      </Router>
    </div>
  );
};

export default App;

const errorMessage = error => (
  <div class="notification is-danger">
    <button class="delete" />
    <p>{error}</p>
  </div>
);
