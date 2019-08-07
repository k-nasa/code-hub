import { h, Component } from "preact";
import firebase from "./firebase";
import { getPrivateMessage, getPublicMessage } from "./api";

class App extends Component {
  constructor() {
    super();
    this.state.user = null;
    this.state.message = "";
    this.state.errorMessage = "";
  }

  componentDidMount() {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        this.setState({ user });
      } else {
        this.setState({
          user: null
        });
      }
    });
  }

  getPrivateMessage() {
    this.state.user
      .getIdToken()
      .then(token => {
        return getPrivateMessage(token);
      })
      .then(resp => {
        this.setState({
          message: resp.message
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  }

  render(props, state) {
    if (state.user === null) {
      return <button onClick={firebase.login}>Please login</button>;
    }
    return (
      <div>
        <div>{state.message}</div>
        <p style="color:red;">{state.errorMessage}</p>
        <button onClick={this.getPrivateMessage.bind(this)}>
          Get Private Message
        </button>
        <button onClick={firebase.logout}>Logout</button>
      </div>
    );
  }
}

export default App;
