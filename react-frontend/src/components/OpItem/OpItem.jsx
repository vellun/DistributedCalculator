import { React, useState } from "react";
import "./OpItem.css";
import Button from "../Button/Button";
import axios from "axios";

function OpItem({ id, name, duration }) {
  const [Duration, setDuration] = useState(duration);

  async function handleSubmit(e) {
    e.preventDefault();
    e.target.reset();

    const resp = await axios.post(
      "http://localhost:8000/operation/",
      {
        Id: id,
        Duration: Number(Duration),
        Name : name,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  }
  return (
    <form className="operation" onSubmit={handleSubmit}>
      <div className="op_content">
        <h1 className="h">{name}</h1>
        <h3>Duration(seconds):</h3>
        <input
          style={{ width: "150%" }}
          type="number"
          placeholder={Duration}
          onChange={(e) => setDuration(e.target.value)}
        />
      </div>
      <div className="exp_btns">
        <Button type="submit">Apply</Button>
      </div>
    </form>
  );
}

export default OpItem;
