import React, { useEffect, useState } from 'react';
import Accordion from 'react-bootstrap/Accordion';
import BasicScene from './components/BasicScene';

function App() {
  const [entities, setEntities] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/api/game')
      .then(res => res.json())
      .then(data => setEntities(data))
      .catch(err => console.error("Error fetching message:", err));
  }, []);

  return (
    <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
        <Accordion.Header>Scene Info</Accordion.Header>
        <Accordion.Body>
          {entities ? (
            entities.map(entity => (
              <div key={entity.Id}>
                Position: ({entity.Position.X}, {entity.Position.Y}, {entity.Position.Z})
              </div>
            ))
          ) : (
            <div>Loading...</div>
          )}
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
