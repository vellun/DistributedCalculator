import { React } from "react";
import "./Button.css";

function Button(props) {
  return <button {...props} className="myBtn">{props.children}</button>;
}

export default Button;
