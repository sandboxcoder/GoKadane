import React, { useEffect, useState } from 'react';
import Accordion from 'react-bootstrap/Accordion';
import BasicScene from './components/BasicScene';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('/api/game')
      .then(res => res.json())
      .then(data => setMessage(data.text))
      .catch(err => console.error("Error fetching message:", err));
  }, []);

  return (
    <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
        <Accordion.Header>Scene Info</Accordion.Header>
        <Accordion.Body>
          {message}
        </Accordion.Body>
      </Accordion.Item>
      <Accordion.Item eventKey="1">
        <Accordion.Header>3D Scene</Accordion.Header>
        <Accordion.Body>
          <BasicScene></BasicScene>
        </Accordion.Body>
      </Accordion.Item>
    </Accordion>
  );
}

export default App;
