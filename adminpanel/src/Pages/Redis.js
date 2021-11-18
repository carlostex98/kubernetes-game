import React, { Component } from 'react';
//const axios = require('axios');


export class Redis extends Component {

  state = {
    dtmongo:[],
    best:[]
  };

  // eslint-disable-next-line no-useless-constructor
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    this.getJoke();
    this.interval = setInterval(() => {
      this.getJoke();
    }, 5000);
  }

  getJoke() {
    fetch("http://35.192.200.215.nip.io/node/redis")
      .then(res => {
        return res.json();
      })
      .then(res => {
        this.setState({
          dtmongo:res.latest,
          best:res.best
        });
      });
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return <div>
      <nav className="navbar is-info">
        <div className="navbar-brand">
          <a className="navbar-item" href="https://bulma.io">
            <h3 className="title is-3">Adminpanel</h3>
          </a>
          <div className="navbar-burger" data-target="navbarExampleTransparentExample">
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>

        <div id="navbarExampleTransparentExample" className="navbar-menu">
          <div className="navbar-start">
            <a className="navbar-item" href="https://bulma.io/">
              Inicio
            </a>
            <a className="navbar-item" href="https://bulma.io/">
              Redis
            </a>
            <a className="navbar-item" href="https://bulma.io/">
              Mongo
            </a>
          </div>

          <div className="navbar-end">
            <div className="navbar-item">
              <div className="field is-grouped">
                <p className="control">
                  <a className="button" href="/">
                    <span>
                      Aprobar ✅
                    </span>
                  </a>
                </p>
              </div>
            </div>
          </div>
        </div>
      </nav>
      <br />

      <div className="container" >
        <div className="box">
          <h3 className="title is-3">Estadisticas del cluster en vivo </h3>
          

          <h3 className="title is-3">Top juegos </h3>
          <div class="tags are-medium">
          {this.state.best.map((object, i) => <> 
            <span class="tag">{"Jugador: "+object._id + " -  ganados "+object.count}</span>
          </>)}
          </div>


          <br></br>
          <h3 className="title is-3">Ultimos 10 </h3>
        <table class="table">
          <thead>
            <tr>
              <th>#</th>
              <th>ID</th>
              <th>numero del serv</th>
              <th>juego num </th>
              <th>nombre del juego</th>
              <th>ganador</th>
              <th>jugadores</th>
              <th>worker</th>
            </tr>
          </thead>
          
          <tbody>
          {this.state.dtmongo.map((object, i) => <> 
            <tr>
              <th>{i}</th>
              <th>{object._id}</th>
              <td>{object.number_req}</td>
              <td>{object.game}</td>
              <td>{object.name_game}</td>
              <td>{object.winner}</td>
              <td>{object.players}</td>
              <td>{object.worker}</td>
            </tr>
          
          </>)}
            
          
          </tbody>
        </table>
        </div>
        
      </div>
    </div>;
  }
}
