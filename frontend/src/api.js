export const endpoint = process.env.REACT_APP_ENDPOINT_HOST;

export const fetchCodes = () => fetch(`${endpoint}/users/codes`);

export const postCode = (idToken, title, body, status) => {
  return fetch(`${endpoint}/codes`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    body: JSON.stringify({
      title: title,
      body: body,
      status: status
    }),
    credentials: "same-origin"
  }).then(res => {
    return res;
  });
};

export const fetchUserCodes = user_id => {
  return fetch(`${endpoint}/users/${user_id}/codes`).then(res => {
    if (!res.ok) {
      throw Error(`Request rejected with status ${res.status}`);
    }

    return res;
  });
};

export const handleError = (res, errorSetter) => {
  if (res === undefined || res === null) {
    errorSetter();
    return false;
  }

  if (res.status !== 201) {
    errorSetter();
    return false;
  }

  return true;
};

export const getCode = (username, title) => {
  return fetch(`${endpoint}/${username}/${title}`).then(res => {
    if (!res.ok) {
      throw Error(`Request rejected with status ${res.status}`);
    }
    return res;
  });
};

export const compileCode = (language, body) => {
  return fetch(`${endpoint}/compile`, {
    method: "POST",
    body: JSON.stringify({
      language: language,
      body: body
    }),
    credentials: "same-origin"
  }).then(res => {
    return res;
  });
};

export const deleteCodeApi = (idToken, id) => {
  return fetch(`${endpoint}/codes/${id}`, {
    method: "DELETE",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin"
  }).then(res => {
    return res;
  });
};

export const getCommentsApi = code_id => {
  return fetch(`${endpoint}/codes/${code_id}/comments`).then(res => {
    if (!res.ok) {
      throw Error(`Request rejected with status ${res.status}`);
    }
    return res;
  });
};

export const postCommentApi = (idToken, code_id, body) => {
  return fetch(`${endpoint}/comments`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    body: JSON.stringify({
      code_id: code_id,
      body: body
    }),
    credentials: "same-origin"
  }).then(res => {
    return res;
  });
};
