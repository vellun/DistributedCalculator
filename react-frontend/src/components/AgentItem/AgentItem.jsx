import { React } from "react";
import "./AgentItem.css";
import Button from "../Button/Button";

function AgentItem({ id, status, last_active }) {
  const s = {};
  if (status == "running") s.color = "green";
  if (status == "missing") s.color = "yellow";
  if (status == "dead") s.color = "red";

  function timestamp_to_string(timestamp) {
    let dateObject = new Date(timestamp * 1000);
    return dateObject.toLocaleString();
  }

  if (last_active !== 0) {
    last_active = timestamp_to_string(last_active);
  } else {
    last_active = "?";
  }

  return (
    <div className="agent">
      <div className="agent_content">
        <h2>Agent {id}</h2>
        <ul>
          <li>
            Status: <span style={s}>{status}</span>
          </li>
          <li>Last active: {last_active}</li>
        </ul>
      </div>
      <div className="agent_btns">
        <Button className="disconnect_btn">Disconnect</Button>
      </div>
    </div>
  );
}

export default AgentItem;
