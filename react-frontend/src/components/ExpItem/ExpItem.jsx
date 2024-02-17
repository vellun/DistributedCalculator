import { React, useState } from "react";
import "./ExpItem.css"
import Button from "../Button/Button";


function ExpItem({expression, status, started_date}) {
  return <div className="expression">
    <div className="exp_content">
        <h3>{expression}=?</h3>
        <ul>
            <li>Status: {status}</li>
            <li>Started date: {started_date}</li>
            <li>Ended date: </li>
        </ul>
    </div>
    <div className="exp_btns">
        <Button>Show details</Button>
    </div>

  </div>;
}

export default ExpItem;
