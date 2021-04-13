import { useState, useEffect } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';
import Spinner from './Spinner';

function App() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState([]);

  useEffect(() => {
    axios.get('http://localhost/').then((res, err) => {
      setLoading(false);
      if (res.status === 200) {
        setData(res.data);
      } else {
        console.log(err);
      }
    });
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />

        <a
          className="App-link"
          href="https://github.com/r-spacex/SpaceX-API/"
          target="_blank"
          rel="noopener noreferrer"
        >
          SpaceX Launches
        </a>

        {loading && <Spinner />}

        <ul className="list">
          {data.map(item => (
            <li key={item.id} className="list__item">
              {item.links.patch.small && (
                <img src={item.links.patch.small} alt="" height="50" />
              )}
              <div>
                #{item.flight_number} <small>{item.name}</small>
              </div>
              <div className="list__details">{item.details}</div>
              {item.failures.length > 0 && (
                <div className="list__failure">
                  Failure: {item.failures[0].reason}
                </div>
              )}
              <div className="list__date">Date: {item.date_utc}</div>
            </li>
          ))}
        </ul>
      </header>
    </div>
  );
}

export default App;
