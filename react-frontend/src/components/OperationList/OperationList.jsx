import { React, useState, useEffect } from "react";
import OpItem from "../OpItem/OpItem";
import axios from "axios";

function OperationList() {
  const [operations, setOp] = useState([]);

  async function fetchExps() {
    const resp = await axios.get("http://localhost:8000/operations/");
    console.log(resp);
    setOp(resp.data);
  }

  useEffect(() => {
    fetchExps();
  }, []);

  return (
    <div className="OpList">
      <ul>
        <div>
          {operations.map((op) => (
            <li key={op.id}>
              <OpItem name={op.name} duration={op.duration} id={op.id} />
            </li>
          ))}{" "}
        </div>
      </ul>
    </div>
  );
}

export default OperationList;
