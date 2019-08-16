import React from "react";
import { useState, useEffect } from "react";

const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

const Index = props => {
  const [codes, setUser] = useState([]);
  const [errorMessage, setMessage] = useState("");

  useEffect(() => {
    init();
  }, []);

  const init = async () => {
    const res = await fetch(`${endpoint}/users/codes`).catch(e => {
      setMessage(e.toString());
    });

    const json = await res.json();

    setUser(json);
  };

  return (
    <div className="App">
      <p>{errorMessage}</p>
      {codes.map((c, i) => {
        {
          // TODO あとでコンポーネントとして切り出す
        }
        return (
          <article key={i} className="media">
            <figure className="media-left">
              <p className="image is-64x64">
                <img
                  src={
                    c.icon_url
                      ? c.icon_url
                      : process.env.REACT_APP_DUMMY_ICON_URL
                  }
                />
              </p>
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
    </div>
  );
};

export default Index;
