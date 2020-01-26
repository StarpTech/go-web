import React from "react";

export default function Header({ message }) {
  return (
    <header className="bg-white text-white">
      <div className="container text-center text-dark">
        <img src="http://127.0.0.1:8080/assets/images/go-web.png" />
        <h1>{message}</h1>
        <p className="lead">Welcome to modern web development with Go</p>
      </div>
    </header>
  );
}
