//import logo from './logo.svg';
//import './App.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  //Link
} from "react-router-dom";
import { Dashboard } from './Pages/Dashboard';
import { Redis } from "./Pages/Redis";
import { User } from "./Pages/User";

function App() {
  return (
    <Router>
        <Routes>
          <Route exact path="/" element={<Dashboard></Dashboard>}> 
          </Route>
          <Route path="/user" element={<User />}> 
          </Route>
          <Route path="/redis" element={<Redis />}> 
          </Route>
        </Routes>
    </Router>
  );
}

export default App;
