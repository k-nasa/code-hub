import React from "react";
import { useState, useEffect } from "react";
import firebase from "./firebase";
import { BrowserRouter as Router, Route} from "react-router-dom";
import Index from "./pages/index";
import WriteCode from "./pages/write_code";
import ErrorMessage from "./component/error_message";

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
      <HeaderMenue />
      <ErrorMessage
        error={errorMessage}
        handler={() => setErrorMessage("")}
      />

      <br />
      <Router>
        <Route path="/" exact render={() => <Index user={user} />} />
        <Route path="/write_code" render={() => <WriteCode user={user} />} />
      </Router>
    </div>
  );
};

export default App;

const HeaderMenue = () => (
  <section className="hero is-dark">
    <div className="hero-head">
      <div className="container">
        {// FIXME テスト用でlogout追加してる
        }
        <p className="title" onClick={firebase.logout}>CodeHub</p>
        {/* <p style={{ textAlign: "right" }} onClick={firebase.logout}>
          logout
        </p> */}
      </div>
    </div>
  </section>
);
