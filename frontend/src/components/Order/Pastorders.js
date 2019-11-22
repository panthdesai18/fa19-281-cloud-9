import React, { Component } from 'react';
import {NavLink} from 'react-router-dom';
import axios from 'axios';
import {Link} from 'react-router-dom';
import {Redirect} from 'react-router';
import { array } from 'prop-types';
import counterburgersymbol from './counterburgersymbol.png';
import './Order.css';
var result=[]
var hostname=""

class Pastorders extends Component {
    constructor(props) {
        super(props);
        this.state = { 
            orders:[]
        }
         
    }

    componentDidMount(){
    
        axios.get('http://a9bdd7cf00d1911ea85f50e68d07da0d-1018085420.us-east-1.elb.amazonaws.com/orders')
                .then((response) => {
                //update the state with the response data
                this.setState({
                    orders : response.data
                });
                console.log(response.data)
                console.log(this.state.orders)
            });
    }

  

    render() {
       // console.log(this.state.listed)
        //console.log(this.state.listed)
        let redirectVar = null;
        if(!localStorage.getItem("user")){
            redirectVar = <Redirect to= "/home"/>
        }
      
        
        let details = this.state.orders.map(order => {
            var items1 = order.Items
            console.log("Items : ",items1)
            var str1 = items1.toString();
            console.log("string: ",str1)
            return(
                <div class="u-clickable u-list">
                <div class="u-flex u-flex-justify u-flex-align">
                                        <div class = "box1">
                                        <div class="order">Order id: {order._id}</div>
                                        <div class="order">Status: {order.OrderStatus}</div>
                                        <div class="order">Items: {str1} </div>
                                        <div class="order">Price: {order.TotalAmount}</div>
                                        </div>
                                        </div></div>
                                   
            )
        })
        return(
            <div>
           
                <img src = {counterburgersymbol} height="100" width="200" alt=""></img>
                <div className="NavbarLinks">
                    &nbsp; &nbsp; <Link to="/home" style={{"font-size": "20px", "font-weight" : "800" , "color":"black", "background-color": "white" }}>HOME</Link> 
                    <Link to="/menu" style={{"font-size": "20px", "font-weight" : "800" , marginLeft: "20px", "color":"black", "background-color": "white"  }}>MENU</Link>
                    <Link to="/pastorders" style={{"font-size": "20px", "font-weight" : "800" , marginLeft: "20px","color":"black", "background-color": "white"  }}>PAST ORDERS</Link>
                    {/* <Link to="/" style={{"font-size": "20px", "font-weight" : "800" , marginLeft: "20px","color":"black", "background-color": "white"  }}></Link> */}
                    </div>
                <div class="ml-5">
                <h1 className="cart">Past Orders: </h1>
                {details}
                </div>
            
            </div>
        )
    }
}

export default Pastorders;