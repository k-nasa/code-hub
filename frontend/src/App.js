import React from "react";
import { useState, useEffect } from "react";
import firebase from "./firebase";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Index from "./pages/index";
import ShowUser from "./pages/show_user";
import WriteCode from "./pages/write_code";
import ErrorMessage from "./component/error_message";

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
      <Router>
        <HeaderMenue />
        <ErrorMessage
          error={errorMessage}
          handler={() => setErrorMessage("")}
        />

        <br />
        <Route path="/" exact render={() => <Index user={user} />} />
        <Route path="/write_code" render={() => <WriteCode user={user} />} />
        <Route path="/user/:id" render={(p) => <ShowUser user={user} p={p}/>} />
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
