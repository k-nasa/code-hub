import React from "react";

const SuccessMessage = props => (
  <div>
    {props.message ? (
      <div className="notification is-primary">
        <button onClick={props.handler} className="delete" />
        <p>{props.message}</p>
      </div>
    ) : (
      <div />
    )}
  </div>
);

export default SuccessMessage;
