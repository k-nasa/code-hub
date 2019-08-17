import React from "react";

import { useState, useEffect } from "react";
import ErrorMessage from "../component/error_message";
import SuccessMessage from "../component/success_message";
import { Link } from "react-router-dom";
import { fetchUserCodes, handleError } from "../api";
import UserIcon from "../component/user_icon";
import CodeContent from "../component/code_content";

const ShowUser = props => {
  const [codes, setCodes] = useState([]);
  const [user, setUser] = useState({});
  
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  useEffect(() => {
      fetchCodes()
  }, [])

  const fetchCodes = async () => {
      console.log(props)
      const res = await fetchUserCodes(props.p.match.params.id).catch((e) => {
          setErrorMessage(e.toString())
      })

    if (res === undefined || res === null) {
      setErrorMessage("failed get user data");
      return ;
    }
      const json = await res.json()

      console.log(json)
      setCodes(json.codes)
      setUser(json.user)
  }

  return (
    <div>
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
                  return (
                    <div key={i}>
                      <CodeContent code={c} />
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
