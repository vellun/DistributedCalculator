import { React, useState } from "react";
import "./History.css";
import ExpressionsList from "../../components/ExpressionList/ExpressionsList";

function History() {
  return (
    <div className="history">
      <h2>Expressions History</h2>
      <ExpressionsList />
    </div>
  );
}

export default History;
