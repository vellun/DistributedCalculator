import { React } from "react";
import "./ExpItem.css";

function ExpItem({ expression, status, result, started_at, ended_at }) {
  const s = {};
  if (status == "complete") s.color = "green";
  if (status == "process") s.color = "yellow";

  function timestamp_to_string(timestamp) {
    let dateObject = new Date(timestamp * 1000);
    return dateObject.toLocaleString();
  }

  started_at = timestamp_to_string(started_at);
  if (ended_at !== 0){
    ended_at = timestamp_to_string(ended_at);
  } else{
    ended_at = "?"
  }


  return (
    <div className="expression">
      <div className="exp_content">
        <h3>
          {expression}=<span>{result ? result : "?"}</span>
        </h3>
        <ul>
          <li>
            Status: <span style={s}>{status}</span>
          </li>
          <li>Started at: {started_at}</li>
          <li>Ended at: {ended_at}</li>
        </ul>
      </div>
    </div>
  );
}

export default ExpItem;
