import React, { Component } from "react";

export class Dashboard extends Component {
  render() {
    return (
        <nav className="navbar is-info">
        <div className="navbar-brand">
          <a className="navbar-item" href="https://bulma.io">
          <h3 class="title is-3">Adminpanel</h3>
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
    );
  }
}

