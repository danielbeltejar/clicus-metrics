import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Dashboard from "./components/Dashboard";
import Login from "./components/Login";
import Register from "./components/Register";
import ShortenUrl from "./components/ShortenUrl";

const App = () => {
    return (
        <Router>
            <div className="max-w-4xl mx-auto">
                <h1 className="text-3xl font-bold text-center mt-4">Clicus Metrics</h1>
                <Routes>
                    <Route path="/" element={<ShortenUrl />} />
                    <Route path="/dashboard" element={<Dashboard />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </div>
        </Router>
    );
};

export default App;
