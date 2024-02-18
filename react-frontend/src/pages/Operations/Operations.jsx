import { React } from "react";
import "./Operations.css";
import OperationList from "../../components/OperationList/OperationList";

function Operations() {
  return (
    <div className="operations">
      <h2 style={{textAlign: "center", marginTop: "30px"}}>Operations</h2>
      <OperationList />
    </div>
  );
}

export default Operations;
