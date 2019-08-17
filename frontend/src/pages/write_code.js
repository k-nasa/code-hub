import React from "react";

import ErrorMessage from "../component/error_message";
import SuccessMessage from "../component/success_message";
import { postCode, handleError } from "../api";
import { useState } from "react";
import { compile } from "path-to-regexp";

const WriteCode = props => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [language, setLanguage] = useState("rust");
  const [status, setStatus] = useState("public");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [compileResult, setCompileResult] = useState("Execution result")

  const submitCode = async () => {
    const user_id = await props.user.getIdToken();

    const res = await postCode(user_id, title, body, status).catch(e => {
      setErrorMessage(e.toString());
    });

    if (!handleError(res, setSavingError)) {
      if (res.status === 400) {
        const json = await res.json();
        if (!json) {
          return;
        }
        setErrorMessage(json.message ? json.message : "failed save");
      }
      return;
    }

    setSavingsuccess();
  };

  const setSavingError = () => {
    setErrorMessage("Failed saving code");
    setTimeout(() => setErrorMessage(""), 2000);
  };

  const setSavingsuccess = () => {
    setSuccessMessage("Saving!!");
    setTimeout(() => setSuccessMessage(""), 2000);
  };

  const HeaderButton = () => (
    <div>
      <label className="label">Select use language</label>
      <div style={{ padding: "5px" }} className="select is-info is-medium">
        <select value={language} onChange={e => setLanguage(e.target.value)}>
          <option value="rust">Rust</option>
          <option value="golang">Golang</option>
          <option value="ruby"> Ruby</option>
        </select>
      </div>
      <button onClick={props.handler} className="button is-info is-medium">
        Run
      </button>
    </div>
  );

  const TextEditor = () => (
    <div>
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
  );

  const CompileOutput = () => (
    <div style={{ paddingTop: "20px" }}>
      <div className="control">
        <textarea className="textarea" readOnly>
          {compileResult}
        </textarea>
      </div>
    </div>
  );

  return (
    <div>
      <ErrorMessage error={errorMessage} handler={() => setErrorMessage("")} />

      <SuccessMessage
        message={successMessage}
        handler={() => setSuccessMessage("")}
      />

      <HeaderButton />
      <div style={{ padding: "30px" }} className="container-padding">
        <TextEditor />
        <CompileOutput />
      </div>
      <FooterButton handler={submitCode} changeStatus={setStatus} />
    </div>
  );
};

export default WriteCode;

const FooterButton = props => (
  <div
    style={{
      bottom: "50px",
      right: "100px",
      position: "fixed"
    }}
  >
    <div style={{ padding: "5px" }} className="select is-info is-medium">
      <select onChange={e => props.changeStatus(e.target.value)}>
        <option value="public">Public saving</option>
        <option value="private">Private storage</option>
        <option value="limited_release"> Limited release</option>
      </select>
    </div>
    <button
      onClick={props.handler}
      className="button is-info is-outlined is-medium"
    >
      Save!!
    </button>
  </div>
);
