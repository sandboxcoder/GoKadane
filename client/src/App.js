import React, { useEffect, useState } from 'react';
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
    <div>
      <div key = "entityInfo">
          {entities ? (
            entities.map(entity => (
              <div key={entity.Id}>
                Position: ({entity.Position.X}, {entity.Position.Y}, {entity.Position.Z})
              </div>
            ))
          ) : (
            <div>Loading...</div>
          )}
      </div>
      <div key = "scene">
          {entities ? (
            <BasicScene entities={entities} />
          ) : (
            <div>Loading 3D Scene...</div>
          )}
      </div>
    </div>
  );
}

export default App;
