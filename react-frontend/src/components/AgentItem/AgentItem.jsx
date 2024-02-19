import { React, useState } from "react";
import "./AgentItem.css";
import Button from "../Button/Button";
import axios from "axios";

function AgentItem({ id, status, last_active }) {
  const [stat, setStatus] = useState(status);
  const s = {};
  if (stat == "running") s.color = "green";
  if (stat == "missing") s.color = "yellow";
  if (stat == "dead") s.color = "red";

  function timestamp_to_string(timestamp) {
    let dateObject = new Date(timestamp * 1000);
    return dateObject.toLocaleString();
  }

  if (last_active !== 0) {
    last_active = timestamp_to_string(last_active);
  } else {
    last_active = "?";
  }

  const [message, setMsg] = useState("");

  async function handleDisconnect() {
    try {
      const resp = await axios.post(
        "http://localhost:8000/disconnect/",
        {
          Id: id,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      console.log(resp.status);
      if (resp.status === 200) {
        setMsg(
          <p>
            Disconnection from the agent {id} done. It will be replaced soon
          </p>
        );
        setStatus("dead");
      }
    } catch {
      setMsg(<p>Could not disconnect from the agent {id}</p>);
    }
  }

  return (
    <div className="agent">
      <div className="agent_content">
        <h2>Agent {id}</h2>
        <ul>
          <li>
            Status: <span style={s}>{stat}</span>
          </li>
          <li>Last active: {last_active}</li>
        </ul>
        {status === "dead" ? <p>Disconnection from the agent {id} done. It will be replaced soon</p> :
        <p>{message}</p>}
      </div>
      <div className="agent_btns">
        <Button onClick={handleDisconnect} className="disconnect_btn">
          Disconnect
        </Button>
      </div>
    </div>
  );
}

export default AgentItem;
