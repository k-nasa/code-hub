import React from "react"
import { Link } from "react-router-dom";

const CodeContent = props => {
  const c = props.code;

  return (
    <div>
      <p>
        <Link to={`/codes/${c.username}/${c.title}`}>
          <strong>{c.title}</strong>
        </Link>
        <br />
        <small>{new Date(c.created_at).toDateString()}</small>
      </p>
      <pre>
        <code>{c.body}</code>
      </pre>
    </div>
  );
};

export default CodeContent;