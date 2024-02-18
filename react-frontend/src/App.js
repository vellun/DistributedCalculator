import React from "react";
import "./App.css";
import Navbar from "./components/Navbar/Navbar";
import Calculator from "./pages/Calculator/Calculator";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import History from "./pages/History/History";
import Operations from "./pages/Operations/Operations";
import Agents from "./pages/Agents/Agents";

function App() {
  return (
    <div className="App">
      <Navbar />
      <BrowserRouter className="App">
        <Routes>
          <Route path="/" element={<Calculator />} />
          <Route path="/history" element={<History />} />
          <Route path="/operations" element={<Operations />} />
          <Route path="/comp-resourses" element={<Agents />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
