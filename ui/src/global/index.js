import React from "react";
import ReactDOM from "react-dom";
import App from "./app";
import "./styles.scss";

var mountNode = document.getElementById("app");
ReactDOM.render(
  <App message="Welcome to modern web development with Go" />,
  mountNode
);
