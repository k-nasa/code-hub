import React from "react";

import CodeContent from "../component/code_content";
import ErrorMessage from "../component/error_message";
import UserIcon from "../component/user_icon";
import firebase from "../firebase";
import { Link } from "react-router-dom";
import { fetchCodes } from "../api";
import { useState, useEffect } from "react";

const Index = props => {
  const [codes, setCodes] = useState([]);
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
    setCodes(json);
  };

  return (
    <div style={{ padding: "20px" }} className="App">
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
              <Link to={`/user/${c.user_id}`}>
                <UserIcon icon_url={c.icon_url} />
              </Link>
            </figure>
            <div className="media-content">
              <div className="content">
                <h2>
                  <Link to={`/user/${c.user_id}`}>
                    {c.username ? c.username : "名無しさん"}
                  </Link>
                </h2>
              </div>
              <CodeContent code={c} />
            </div>
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
