import React from "react";

import CodeContent from "../component/code_content";
import ErrorMessage from "../component/error_message";
import UserIcon from "../component/user_icon";
import { Link } from "react-router-dom";
import { getCode, getCommentsApi, postCommentApi } from "../api";
import { useState, useEffect } from "react";

const ShowCode = props => {
  const username = props.p.match.params.username;
  const title = props.p.match.params.title;

  const [code, setCode] = useState({});
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState("");

  const [errorMessage, setErrorMessage] = useState("");
  const uid = props.user ? props.user.uid : null;

  useEffect(() => {
    fetchCode().catch(e => setErrorMessage(e.toString()));

    if (code.id) {
      fetchComments(code.id).catch(e => setErrorMessage(e.toString()));
    }
  }, [code.id]);

  const fetchCode = async () => {
    const res = await getCode(username, title).catch(e => {
      setErrorMessage(e.toString());
      return;
    });

    const json = await res.json();
    setCode(json);
  };

  const fetchComments = async code_id => {
    const res = await getCommentsApi(code_id).catch(e => {
      setErrorMessage(e.toString());
    });

    if (!res.ok || res.status !== 200) {
      return;
    }

    const json = await res.json();
    console.log(json);
    setComments(json);
  };

  const postComment = async () => {
    if(newComment === "") {
      setErrorMessage("Please input commnet");
      return
    }

    if(props.user === null) {
      setErrorMessage("Please login");
      return
    }
    const user_id = await props.user.getIdToken();
    const res = await postCommentApi(user_id, code.id, newComment);

    if (res.status !== 201) {
      return;
    }
    setNewComment("");
    const json = await res.json();

    console.log(json);
    setComments(comments.concat(json));
  };

  const ShowCode = () => (
    <article className="media">
      <figure className="media-left">
        <UserIcon icon_url={code.icon_url} />
      </figure>
      <div className="media-content">
        <div className="content">
          <h2>
            <Link to={`/user/${code.user_id}`}>
              {code.username ? code.username : "名無しさん"}
            </Link>
          </h2>
        </div>
        {code ? (
          <CodeContent code={code} show_edit={uid === code.firebase_uid} />
        ) : (
          <div />
        )}
      </div>
    </article>
  );

  const Comments = () => (
    <div style={{ padding: "30px" }}>
      {comments.map((comment, i) => {
        return (
          <article key={i} className="media">
            <figure className="media-left">
              <UserIcon icon_url={comment.icon_url} />
            </figure>
            <div className="media-content">
              <div className="content">
                <strong>
                  <Link to={`/user/${comment.user_id}`}>
                    {comment.username ? comment.username : "名無しさん"}
                  </Link>
                </strong>
              </div>
              {comment.body}
            </div>
          </article>
        );
      })}
    </div>
  );

  return (
    <div>
      <ErrorMessage
        error={errorMessage}
        handler={() => setErrorMessage("")}
      />
      <div style={{ padding: "16px" }}>
        <ShowCode />

        <Comments />
        <CommentInput
          clickHandler={() => postComment()}
          textHandler={e => setNewComment(e.target.value)}
          commentContent={newComment}
        />
      </div>
    </div>
  );
};

const CommentInput = props => (
  <div style={{ padding: "24px" }} className="media-content">
    <div className="field">
      <p className="control">
        <textarea
          className="textarea"
          value={props.commentContent}
          onChange={e => props.textHandler(e)}
          placeholder="Add a comment..."
        />
      </p>
    </div>
    <nav className="level">
      <div className="level-left">
        <div className="level-item">
          <button onClick={props.clickHandler} className="button is-info">
            Submit
          </button>
        </div>
      </div>
    </nav>
  </div>
);

export default ShowCode;
