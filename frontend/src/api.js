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
