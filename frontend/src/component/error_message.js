import React from "react";

const ErrorMessage = props => (
  <div class="notification is-danger">
    <button onClick={props.handler} class="delete" />
    <p>{props.error}</p>
  </div>
);

export default ErrorMessage;
