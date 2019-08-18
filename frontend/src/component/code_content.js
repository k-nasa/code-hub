import React from "react";
import { Link } from "react-router-dom";

const CodeContent = props => {
  const c = props.code;
  const content = props.is_ommit
    ? c.body.slice(0, 300) + (c.body.length > 300 ? "\n......\nomitted" : "")
    : c.body;

  return (
    <div>
      <p>
        <Link to={`/${c.username}/${c.title}`}>
          <strong>{c.title}</strong>
        </Link>
        <br />
        <small>{new Date(c.created_at).toDateString()}</small>
      </p>
      <pre className="prettyprint">
        <code>{content}</code>
      </pre>
    </div>
  );
};

export default CodeContent;
