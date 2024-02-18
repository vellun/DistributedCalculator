import React from "react";
import "./Navbar.css";

function Navbar() {
  return <div className="navbar">
    <a className="logo_a" href="/"><h1 className="logo">Distributed<span>Calc☆</span></h1></a>

    <ul><li className="nav-link"><a className="nav-link_a" href="/">Calculator</a></li>
        <li className="nav-link"><a className="nav-link_a" href="/operations">Operations</a></li>
        <li className="nav-link"><a className="nav-link_a" href="/history">History</a></li>
        <li className="nav-link"><a className="nav-link_a" href="/comp-resourses">Computing resourses</a></li>
        </ul>
  </div>;
}

export default Navbar;