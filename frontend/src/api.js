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
    if (!res.ok) {
      throw Error(`Request rejected with status ${res.status}`);
    }
    return res;
  });
};
