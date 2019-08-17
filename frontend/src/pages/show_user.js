import React from "react";

import CodeContent from "../component/code_content";
import ErrorMessage from "../component/error_message";
import SuccessMessage from "../component/success_message";
import UserIcon from "../component/user_icon";
import { Link } from "react-router-dom";
import { fetchUserCodes, handleError } from "../api";
import { useState, useEffect } from "react";

const ShowUser = props => {
  const [codes, setCodes] = useState([]);
  const [user, setUser] = useState({});
  
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
      fetchCodes().then(json => {
        setCodes(json.codes);
        setUser(json.user);
      });
  }, [])

  const fetchCodes = async () => {
      const res = await fetchUserCodes(props.p.match.params.id).catch((e) => {
          setErrorMessage(e.toString())
      })

    if (res === undefined || res === null) {
      setErrorMessage("failed get user data");
      return ;
    }
      return await res.json();
  }

  return (
    <div style={{padding: "20px"}}>
      <ErrorMessage error={errorMessage} handler={() => setErrorMessage("")} />
      <article className="media">
        <figure className="media-left">
          <UserIcon icon_url={user.icon_url} />
        </figure>
        <div className="media-content">
          <div className="content">
            <h2>{user.username ? user.username : "名無しさん"}</h2>
          </div>
          {
              codes.map((c, i) => {
                c.username = user.username;
                  return (
                    <div style={{padding: "16px"}}key={i}>
                      <article className="media">
                        <div className="media-content">
                          <div className="content">
                            <CodeContent code={c} />
                          </div>
                        </div>
                      </article>
                    </div>
                  );
              })
          }
        </div>
      </article>
    </div>
  );
};

export default ShowUser;
