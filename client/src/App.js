import React, { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('/api/game')
      .then(res => res.json())
      .then(data => setMessage(data.text))
      .catch(err => console.error("Error fetching message:", err));
  }, []);

  return (
    <div style={{ padding: "2rem", fontFamily: "sans-serif" }}>
      <h1>React + Go</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;
