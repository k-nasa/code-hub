import React from "react";
import { useState } from "react";
import { postCode } from "../api";
import ErrorMessage from "../component/error_message";
import SuccessMessage from "../component/success_message";
import "../style.css";

const WriteCode = props => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [status, setStatus] = useState("public");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const submitCode = async () => {
    const user_id = await props.user.getIdToken();

    const res = await postCode(user_id, title, body, status).catch(e => {
      setErrorMessage(e.toString());
    });

    if (res === undefined || res === null) {
      setErrorMessage("Failed save code");
      return;
    }

    if (res.status !== 201) {
      setErrorMessage("Failed saving code");
    }

    setSuccessMessage("Saving!!");
  };

  return (
    <div>
      <ErrorMessage error={errorMessage} handler={() => setErrorMessage("")} />

      <SuccessMessage
        message={successMessage}
        handler={() => setSuccessMessage("")}
      />

      <div style={{ padding: "30px" }} className="container-padding">
        <div className="field">
          <div className="control">
            <input
              style={{ padding: "30px", marginBottom: "20px" }}
              onChange={e => setTitle(e.target.value)}
              placeholder="A great title for this code"
              value={title}
              className="input is-large"
              type="text"
            />
            <textarea
              onChange={e => setBody(e.target.value)}
              value={body}
              placeholder="Great code here!"
              className="textarea is-info"
            />
          </div>
        </div>
      </div>
      <FooterButton handler={submitCode} changeStatus={setStatus} />
    </div>
  );
};

export default WriteCode;

const FooterButton = props => (
  <div
    style={{
      bottom: "20px",
      right: "20px",
      position: "fixed"
    }}
  >
    <div className="field">
      <div className="control">
        <div className="select is-rounded is-info is-medium">
          <select onChange={e => props.changeStatus(e.target.value)}>
            <option value="public">Public saving</option>
            <option value="private">Private storage</option>
            <option value="limited_release"> Limited release</option>
          </select>
          <button
            onClick={props.handler}
            className="button is-link is-outlined"
          >
            Save!!
          </button>
        </div>
      </div>
    </div>
  </div>
);
