import React, { Component } from 'react';
//const axios = require('axios');


export class Dashboard extends Component {

  state = {
    kafkaNum: 0,
    rabbitNum: 0,
    pubNum: 0,
    dtmongo:[],
    top:[]
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
    fetch("http://35.192.200.215.nip.io/node/dashboard")
      .then(res => {
        return res.json();
      })
      .then(res => {
        this.setState({
          kafkaNum: res.worker_report[0].count,
          pubNum: res.worker_report[1].count,
          rabbitNum: res.worker_report[2].count,
          dtmongo:res.data_mongo,
          top:res.top_3
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
          <h3 className="title is-3">Estadisticas del cluster </h3>
          <div className="tile is-ancestor has-text-centered">
            <div className="tile is-parent">
              <article className="tile is-child box">
                <p className="title">{this.state.rabbitNum}</p>
                <p className="subtitle">RabbitMQ</p>
              </article>
            </div>
            <div className="tile is-parent">
              <article className="tile is-child box">
                <p className="title">{this.state.pubNum}</p>
                <p className="subtitle">Pubsub</p>
              </article>
            </div>
            <div className="tile is-parent">
              <article className="tile is-child box">
                <p className="title">{this.state.kafkaNum}</p>
                <p className="subtitle">Kafka</p>
              </article>
            </div>
          </div>

          <h3 className="title is-3">Top juegos </h3>
          <div class="tags are-medium">
          {this.state.top.map((object, i) => <> 
            <span class="tag">{"Juego: "+object._id + " -  ganados "+object.count}</span>
          </>)}
          </div>


          <br></br>
          <h3 className="title is-3">Datos Mongo </h3>
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

//export default Dashboard;


/*export class Dashboard extends Component {

  a = 1;
  b = 1;
  c = 1;

  // eslint-disable-next-line no-useless-constructor



  constructor() {

    super();


    this.reqServer(this);

  }

  reqServer(ez){
    axios.get('http://35.192.200.215.nip.io/node/dashboard').then(function (response) {
    // handle success
      ez.setterx(response.data)
    })
  }

  setterx(response) {
    this.a = response.worker_report[2].count
    this.b = response.worker_report[1].count
    this.c = response.worker_report[0].count

    console.log(this.a)
  }



  render() {
    return (
      <div>
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
            <h3 className="title is-3">Estadisticas del cluster </h3>
            <div className="tile is-ancestor has-text-centered">
              <div className="tile is-parent">
                <article className="tile is-child box">
                  <p className="title">{this.a}</p>
                  <p className="subtitle">RabbitMQ</p>
                </article>
              </div>
              <div className="tile is-parent">
                <article className="tile is-child box">
                  <p className="title">{this.b}</p>
                  <p className="subtitle">Pubsub</p>
                </article>
              </div>
              <div className="tile is-parent">
                <article className="tile is-child box">
                  <p className="title">{this.c}</p>
                  <p className="subtitle">Kafka</p>
                </article>
              </div>
            </div>
          </div>
        </div>
      </div>

    );
  }
}
*/
