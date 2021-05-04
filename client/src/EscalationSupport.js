import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";
import "./EscalationApp.css";

let endpoint = "http://localhost:8080";

export class EscalationSupport extends Component {
  constructor(props) {
    super(props);
    this.state = ({
        nombre:'',
        ticketJira:'',
        estudio: false
      })
    this.procesar = this.procesar.bind(this);
    this.nombreApp = this.nombreApp.bind(this);
    this.cambioTicket = this.cambioTicket.bind(this);    
    this.cambioEstudio = this.cambioEstudio.bind(this);
  }


  render() {
    return (
      <div>
        <form onSubmit={this.procesar}>
          <p>Ingrese nombre App:<input type="text" value={this.state.nombre} onChange={this.nombreApp} /></p>
          <p>Ingrese ticket Jira:<input type="text" value={this.state.ticketJira} onChange={this.cambioTicket} /></p>
          <p>Estudios:<input type="checkbox" value={this.state.estudio} onChange={this.cambioEstudio} /></p>          
          <p><input type="submit" /></p>
        </form>
        <hr />
      </div>
    );
  }

  procesar(e) {
    e.preventDefault();
    alert('Dato cargado '+this.state.nombre + ' ' + 
                         +this.state.edad + ' ' + 
                         +this.state.estudio);
  }

  nombreApp(e) {
    this.setState( {
      nombre: e.target.value
    })
  }

  cambioTicket(e) {
    this.setState( {
        ticketJira: e.target.value
    })
  }  

  cambioEstudio(e) {
    this.setState( {
      estudio: !this.state.estudio
    })
  }    
}

export default EscalationSupport;