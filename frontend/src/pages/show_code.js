import React from "react";

import CodeContent from "../component/code_content";
import ErrorMessage from "../component/error_message";
import UserIcon from "../component/user_icon";
import { Link } from "react-router-dom";
import { getCode } from "../api";
import { useState, useEffect } from "react";

const ShowCode = props => {
  const username = props.p.match.params.username;
  const title = props.p.match.params.title;

  const [code, setCode] = useState({});

  const [errorMessage, setErrorMessage] = useState("");
  const uid = props.user ? props.user.uid : null;

  useEffect(() => {
    fetchCode().catch(e => setErrorMessage(e.toString()));
  }, []);

  const fetchCode = async () => {
    const res = await getCode(username, title).catch(e => {
      setErrorMessage(e.toString());
    });

    const json = await res.json();
    setCode(json);
  };

  return (
    <div style={{ padding: "16px" }}>
      <ErrorMessage error={errorMessage} handler={() => setErrorMessage("")} />

      <article className="media">
        <figure className="media-left">
          <UserIcon icon_url={code.icon_url} />
        </figure>
        <div className="media-content">
          <div className="content">
            <h2>
              <Link to={`/user/${code.user_id}`}>
                {code.username ? code.username : "名無しさん"}
              </Link>
            </h2>
          </div>
          {code ? (
            <CodeContent code={code} show_edit={uid === code.firebase_uid} />
          ) : (
            <div />
          )}
        </div>
      </article>
    </div>
  );
};

export default ShowCode;
