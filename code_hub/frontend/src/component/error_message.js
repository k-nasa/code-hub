import React from "react";

const ErrorMessage = props => (
  <div>
    {props.error ? (
      <div className="notification is-danger">
        <button onClick={props.handler} className="delete" />
        <p>{props.error}</p>
      </div>
    ) : (
      <div />
    )}
  </div>
);

export default ErrorMessage;
