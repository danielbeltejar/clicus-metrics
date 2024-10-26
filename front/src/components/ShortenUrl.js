import React, { useState } from "react";

const ShortenUrl = () => {
    const [url, setUrl] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem("token"); // Retrieve JWT from local storage

        const response = await fetch("https://cliclus.danielbeltejar.es/api/v1/shorten", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${token}`, // Send JWT for authorization
            },
            body: JSON.stringify({ originalUrl: url }),
        });

        if (response.ok) {
            const data = await response.json();
            alert(`Shortened URL: ${data.shortenedUrl}`);
        } else {
            alert("Error shortening the URL.");
        }
    };

    return (
        <div className="p-4">
            <h2 className="text-xl font-bold">Shorten URL</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Enter URL"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    className="border px-4 py-2 w-full mt-2"
                    required
                />
                <button type="submit" className="bg-blue-500 text-white px-4 py-2 mt-2">
                    Shorten
                </button>
            </form>
        </div>
    );
};

export default ShortenUrl;
