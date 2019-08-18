import React from "react";
import { Link } from "react-router-dom";
import { deleteCodeApi } from "../api";

const CodeContent = props => {
  const c = props.code;

  const uid = props.get_user ? props.get_user.firebase_uid : null;
  const firebase_uid = props.login_user ? props.login_user.uid : null;

  const content = props.is_ommit
    ? c.body.slice(0, 300) + (c.body.length > 300 ? "\n......\nomitted" : "")
    : c.body;

  const deleteCode = async () => {
    const user_id = await props.login_user.getIdToken();

    const res = await deleteCodeApi(user_id, c.id).catch(e => {
      console.log(e);
    });

    if (res.ok) {
      window.location.reload();
    }
  };

  const Links = () => (
    <div>
      <Link
        style={{
          padding: "10px"
        }}
        to={{
          pathname: "/write_code",
          state: { title: c.title, body: c.body }
        }}
      >
        edit
      </Link>
      <a
        style={{
          padding: "10px"
        }}
        onClick={deleteCode}
      >
        delete
      </a>
    </div>
  );
  return (
    <div>
      <p>
        <Link to={`/${c.username}/${c.title}`}>
          <strong>{c.title}</strong>
        </Link>
        <br />
        <small>{new Date(c.created_at).toDateString()}</small>
      </p>
      {uid !== null && uid !== undefined && uid === firebase_uid ? (
        <Links />
      ) : (
        <div />
      )}
      <pre className="prettyprint">
        <code>{content}</code>
      </pre>
    </div>
  );
};

export default CodeContent;
