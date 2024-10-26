import React, { useState } from "react";

const Register = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch("https://cliclus.danielbeltejar.es/api/v1/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            alert("User registered successfully!");
        } else {
            alert("Registration failed.");
        }
    };

    return (
        <div className="p-4">
            <h2 className="text-xl font-bold">Register</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    className="border px-4 py-2 w-full mt-2"
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="border px-4 py-2 w-full mt-2"
                    required
                />
                <button type="submit" className="bg-blue-500 text-white px-4 py-2 mt-2">
                    Register
                </button>
            </form>
        </div>
    );
};

export default Register;
