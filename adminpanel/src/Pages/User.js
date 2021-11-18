import React, { Component } from 'react';
//import { withRouter } from "react-router";
//const axios = require('axios');



export class User extends Component {


    state = {
        userx: "1",
        dtmongo: [],
        ax:"1"
    };

    componentDidMount() {
        this.getJoke();

        this.interval = setInterval(() => {
            this.getJoke();
        }, 2000);

        this.handleChange = this.handleChange.bind(this);
    }

    getJoke() {
        // eslint-disable-next-line react-hooks/rules-of-hooks
        //const id = this.props.id;
        //console.log(id)
        //const { id } = this.props.match.params
        fetch("http://35.192.200.215.nip.io/node/user/" + this.state.userx)
            .then(res => {
                return res.json();
            })
            .then(res => {
                this.setState({
                    dtmongo: res,
                });
            });
    }

    clickeo(){
        this.setState({
            userx:this.state.ax
        })
    }

    handleChange(event) {    this.setState({userx: event.target.value}); this.getJoke() }

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
                    <h3 className="title is-3">Datos en vivo </h3>

                    <div class="field has-addons">
                        <div class="control">
                            <input class="input" type="text" placeholder="Find a repository" value={this.state.userx} onChange={this.handleChange}  />
                        </div>
                        <div class="control">
                            <button class="button is-info" on>
                                Search
                            </button>
                        </div>
                    </div>

                    <h3 className="title is-3">Datos del usuario: {this.state.userx} </h3>
                    <table className="table">
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