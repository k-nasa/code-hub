import React from "react";
import { useState, useEffect } from "react";
import ErrorMessage from "../component/error_message";
import firebase from "../firebase";
import { fetchCodes } from '../api'
import { Link } from "react-router-dom";

const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

const Index = props => {
  const [codes, setUser] = useState([]);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    init();
  }, []);

  const init = async () => {
    const res = await fetchCodes().catch(e => {
      setErrorMessage(e.toString());
    });

    if (res === undefined || res === null) {
      return;
    }

    const json = await res.json();

    setUser(json);
  };

  return (
    <div className="App">
      {errorMessage ? (
        <ErrorMessage
          error={errorMessage}
          handler={() => setErrorMessage("")}
        />
      ) : (
        <div />
      )}
      {codes.map((c, i) => {
        return (
          <article key={i} className="media">
            <figure className="media-left">
              <UserIcon icon_url={c.icon_url} />
            </figure>
            <CodeContent code={c} />
          </article>
        );
      })}

      {props.user ? (
        <Link to="write_code">
          <FooterButton text="Write code!!" />
        </Link>
      ) : (
        <FooterButton text="Sign up!!" handler={firebase.login} />
      )}
    </div>
  );
};

export default Index;

const FooterButton = props => (
  <button
    style={{
      bottom: "20px",
      right: "20px",
      position: "fixed"
    }}
    onClick={props.handler}
    className="button is-link is-outlined"
  >
    <p style={{ fontSize: "24px" }}>{props.text}</p>
  </button>
);
const UserIcon = props => (
  <p className="image is-64x64">
    <img
      src={
        props.icon_url ? props.icon_url : process.env.REACT_APP_DUMMY_ICON_URL
      }
    />
  </p>
);

const CodeContent = props => {
  const c = props.code;

  return (
    <div className="media-content">
      <div className="content">
        <h2>
          <Link to={`/user/${c.user_id}`}>
            {c.username ? c.username : "名無しさん"}
          </Link>
        </h2>
        <p>
          <Link to={`/codes/${c.id}`}>
            <strong>{c.title}</strong>
          </Link>
          <br />
          <small>{new Date(c.created_at).toDateString()}</small>
        </p>
        <pre>
          <code>{c.body}</code>
        </pre>
      </div>
    </div>
  );
};
