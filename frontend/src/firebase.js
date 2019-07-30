import firebase from "firebase/app";
import "firebase/auth";

const firebaseConfig = {
  apiKey: process.env.FIREBASE_APIKEY,
  authDomain: process.env.FIREBASE_AUTHDOMAIN,
  databaseURL: process.env.FIREBASE_DATABASEURL,
  projectId: process.env.FIREBASE_PROJECTID,
  messagingSenderId: process.env.FIREBASE_MESSAGINGSENDERID,
  appId: process.env.FIREBASE_APPID
};

// https://firebase.google.com/docs/auth/web/start
const githubProvider = new firebase.auth.GithubAuthProvider();

const FirebaseFactory = () => {
  firebase.initializeApp(firebaseConfig);
  let auth = firebase.auth();
  return {
    auth() {
      return auth;
    },

    login() {
      return auth.signInWithPopup(githubProvider);
    },

    logout() {
      return auth.signOut();
    }
  };
};

export default FirebaseFactory();
