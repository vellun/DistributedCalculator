import React from "react";
import "./App.css";
import ExpressionsList from "./components/ExpressionList/ExpressionsList";
import Navbar from "./components/Navbar/Navbar";
import Calculator from "./pages/Calculator/Calculator";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import History from "./pages/History/History";

function App() {
  return (
    <div className="App">
      <Navbar />
      <BrowserRouter className="App">
        <Routes>
          <Route path="/" element={<Calculator />} />
          <Route path="/history" element={<History />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
