import React from "react";

const UserIcon = props => (
  <p className="image is-64x64">
    <img
      src={
        props.icon_url ? props.icon_url : process.env.REACT_APP_DUMMY_ICON_URL
      }
    />
  </p>
);

export default UserIcon;
