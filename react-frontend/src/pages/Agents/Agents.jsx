import { React } from "react";
import AgentList from "../../components/AgentList/AgentList";

function Agents() {
  return (
    <div className="agents">
      <h2 style={{textAlign: "center", marginTop: "30px"}}>Computing resources</h2>
      <AgentList />
    </div>
  );
}

export default Agents;
