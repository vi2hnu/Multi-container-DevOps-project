import { useState } from "react";
import "./home.css";

export function Home() {
    const [url, setUrl] = useState("");
    const [shortenedUrl, setShortenedUrl] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch("http://localhost:5000/create", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ original_url: url }),
            });

            const data = await response.json();
            if (data.shortened_url) {
                setShortenedUrl(`shortened url: http://localhost:5000/${data.shortened_url}`);
            }
        } catch (err) {
            console.error("Error:", err);
            setShortenedUrl("An error occurred. Please try again.");
        }
    };

    return (
        <div>
            <div className="right-container"></div>

            <div className="text-box">
                <h1 className="text">
                    Build stronger digital <span className="connections">connections</span>
                </h1>
            </div>

            <div className="container">
                <div className="input">
                    {!shortenedUrl ? (
                        <form onSubmit={handleSubmit}>
                            <h1 className="url-text">
                                Please enter the <span className="connections">URL</span> to be shortened
                            </h1>
                            <input
                                className="url"
                                name="original_url"
                                placeholder="URL"
                                value={url}
                                onChange={(e) => setUrl(e.target.value)}
                            />
                            <br />
                            <button className="button" type="submit">
                                submit
                            </button>
                        </form>
                    ) : (
                        <h1 className="url-text">{shortenedUrl}</h1>
                    )}
                </div>
            </div>

            <div className="left-container"></div>
        </div>
    );
}
