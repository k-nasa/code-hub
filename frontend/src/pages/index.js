import React from "react";
import { useState, useEffect } from "react";
import ErrorMessage from "../component/error_message";

const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

const Index = props => {
  const [codes, setUser] = useState([]);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    init();
  }, []);

  const init = async () => {
    const res = await fetch(`${endpoint}/users/codes`).catch(e => {
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
        {
          // TODO あとでコンポーネントとして切り出す
        }
        return (
          <article key={i} className="media">
            <figure className="media-left">
              <UserIcon icon_url={c.icon_url} />
            </figure>
            <div className="media-content">
              <div className="content">
                <h2>{c.username ? c.username : "名無しさん"}</h2>
                <p>
                  <strong>{c.title}</strong>
                  <br />
                  <small>{new Date(c.created_at).toDateString()}</small>
                </p>
                <pre>
                  <code>{c.body}</code>
                </pre>
              </div>
            </div>
          </article>
        );
      })}

      {props.user ? (
        <FooterButton text="Write code!!" />
      ) : (
        <FooterButton text="Sign up!!" />
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
    className="button is-link outlined"
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
