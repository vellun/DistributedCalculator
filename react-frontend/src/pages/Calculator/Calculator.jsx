import Button from "../../components/Button/Button";
import "./Calculator.css";
import axios from "axios";

import { React, useState } from "react";

function Calculator() {
  const [Exp, setExp] = useState("");
  const [message, setMessage] = useState();

  async function handleSubmit(e) {
    setMessage("");
    e.preventDefault();
    e.target.reset();
    try {
      const resp = await axios.post(
        "http://localhost:8000/expression/",
        {
          Exp: Exp,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      console.log(resp.status);
      if (resp.status === 201) {
        setMessage(
          <p>Your expression is counting... Check <a style={{color: "white"}} href="/history">history</a> to see details</p>
        );
      }
    } catch {
      setMessage("Invalid expression");
    }
    setExp("");
  }

  return (
    <div className="calculator">
      <h1>☆Calculator☆</h1>
      <br></br>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Enter your expression"
          onChange={(e) => setExp(e.target.value)}
        />

        <Button type="submit">Let's count!</Button>
      </form>

      <div className="message">
        <p>{message}</p>
      </div>

      <br></br>
      <div className="rulesDiv">
        <b>Binary operations with integers only</b>
        <p>Valid operations: + - / *</p>
        <h4>Examples of valid expressions:</h4>
        <li>5*9+2</li>
        <li>5*(9+2)</li>

        <h4>Examples of INvalid expressions:</h4>
        <li>5.2*9+2.5</li>
        <li>-5*(9+2)**2</li>
      </div>
      <br></br>
    </div>
  );
}

export default Calculator;
