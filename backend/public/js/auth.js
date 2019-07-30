export const lock = new Auth0LockPasswordless(
    'eBuDdl0JieLr912SYGa0JyDIVWeBwor7',
    'treasureapp.auth0.com',
    {
        passwordlessMethod: "link",              // Sets Lock to use magic link
        auth: {
            redirectUrl: 'http://localhost:1991',
            responseType: 'token id_token'
        }
    }
);

lock.on('authenticated', function(authResult) {
    localStorage.setItem('id_token', authResult.idToken);
    localStorage.setItem('access_token', authResult.accessToken);
});
