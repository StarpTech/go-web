import React from "react";
import ReactDOM from "react-dom";
import App from "./app";
import LikeButton from "./components/like-button";
import Header from "./components/header";
import "./styles.scss";

var mountNode = document.getElementById("app");
ReactDOM.render(<App />, mountNode);

document.querySelectorAll(".like-button-component").forEach(domContainer => {
  ReactDOM.render(<LikeButton {...domContainer.dataset} />, domContainer);
});
document.querySelectorAll(".header-component").forEach(domContainer => {
  ReactDOM.render(<Header {...domContainer.dataset} />, domContainer);
});
