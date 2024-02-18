import { React, useState, useEffect } from "react";
import ExpItem from "../ExpItem/ExpItem";
import Button from "../Button/Button";
import axios from "axios";
import AgentItem from "../AgentItem/AgentItem";

function AgentList() {
  const [agents, setAgent] = useState([]);

  async function fetchExps() {
    const resp = await axios.get("http://localhost:8000/agents/");
    setAgent(resp.data);
    console.log(agents)
  }

  useEffect(() => {
    fetchExps();
  }, []);

  return (
    <div className="AgentList">
      <ul>
        <div>
          {agents.map((agent) => (
            <li key={agent.id}>
              <AgentItem
                id={agent.id}
                status={agent.status}
                last_active={agent.last_active}
              />
            </li>
          ))}{" "}
        </div>
      </ul>
    </div>
  );
}

export default AgentList;
