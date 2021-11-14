//import logo from './logo.svg';
//import './App.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  //Link
} from "react-router-dom";
import { Dashboard } from './Pages/Dashboard';

function App() {
  return (
    <Router>
        <Routes>
          <Route exact path="/" element={<Dashboard></Dashboard>}> 
          </Route>
        </Routes>
    </Router>
  );
}

export default App;
