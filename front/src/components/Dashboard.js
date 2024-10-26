import React, { useEffect, useState } from "react";

const Dashboard = () => {
    const [urls, setUrls] = useState([]);

    useEffect(() => {
        const fetchUrls = async () => {
            const token = localStorage.getItem("token"); // Retrieve JWT from local storage

            const response = await fetch("https://cliclus.danielbeltejar.es/api/v1/dashboard", {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${token}`, // Send JWT for authorization
                },
            });

            if (response.ok) {
                const data = await response.json();
                setUrls(data);
            } else {
                alert("Failed to fetch URLs.");
            }
        };

        fetchUrls();
    }, []);

    return (
        <div className="p-4">
            <h2 className="text-xl font-bold">Dashboard</h2>
            <table className="min-w-full mt-4">
                <thead>
                    <tr>
                        <th className="border px-4 py-2">Original URL</th>
                        <th className="border px-4 py-2">Shortened URL</th>
                    </tr>
                </thead>
                <tbody>
                    {urls.map((url) => (
                        <tr key={url.id}>
                            <td className="border px-4 py-2">{url.originalUrl}</td>
                            <td className="border px-4 py-2">{url.shortenedUrl}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default Dashboard;
