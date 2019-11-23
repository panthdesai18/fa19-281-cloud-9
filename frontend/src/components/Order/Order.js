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

class Order extends Component {
    constructor(props) {
        super(props);
        this.state = { 
            listed:[],
            error_status:" ",
            item_delete:"available",
            message:"",
            order: this.props.location.state.itemdetail,
            total_amount: this.props.location.state.total,
            names: this.props.location.state.itemname
        }
        
         this.placeorder =this.placeorder.bind(this)
    }

    async componentDidMount(){
        console.log("in component did mount",this.props.location.state.itemdetail)
        console.log("Inside cmpdidmount")
        console.log("itemdetail",this.state.order, this.state.total_amount, this.state.names)
        var v1 = this.state.order.toString();
        console.log(v1)

    }

   

    placeorder=(e)=>{
        try{
            var total1 = this.state.total_amount;
            var total2 = total1.toString();
            const data = {
                UserId:localStorage.getItem("id"),
                OrderStatus:"placed",
                Items: this.state.names,
                TotalAmount: total2,
            }
         console.log("Data : ",data);
         //set the with credentials to true
         axios.defaults.withCredentials = false;
         //make a post request with the user data
         axios.post('https://gevnsiba07.execute-api.us-east-1.amazonaws.com/prod/ordermanagement/placeOrder', data)
             .then((response) => {
                 console.log("Status Code : ", response);
                 let redirectVar = null;
                 if (response.status == 200) {
                     this.setState({
                         authFlag: true,
                         message: "Your order is successfully placed!"
                     })
                 }
                 else if (response.status == 201) {
                     this.setState({
                         authFlag: false,
                         message:""
                     })
                 }
             }
             );
         }catch(e){}
     }
  

    render() {
        //console.log(this.state.listed)
        //console.log(this.state.listed)
        let redirectVar = null;
        if(!localStorage.getItem("user")){
            redirectVar = <Redirect to= "/home"/>
        }
      const templates = this.state.listed;
      console.log("length of listed:",this.state.listed.length)
      const fullrecord = this.state.listed;
           
        var amount=0.0
        

        let details = this.state.order.map(orders => {
            return(
                <div class="u-list-heading ">
                <div class="u-flex u-flex-justify u-flex-align">
                    <div class = "box1">
                                    <div class="order">Order Items:</div>
                                    <div class = "box2">
                                    <div class="order">{orders[0]}</div> <div class = "order1">${orders[1]}</div>
                                    <br/>
                    </div>
                                    <br/>  
                                </div></div>
                                </div>
            )
        })
        return(
            <div>
           
                <img src = {counterburgersymbol} height="100" width="200" alt=""></img>
                <div className="NavbarLinks">
                    &nbsp; &nbsp; <Link to="/home" style={{"font-size": "20px", "font-weight" : "800" , "color":"black", "background-color": "white" }}>HOME</Link> 
                    <Link to="/menu" style={{"font-size": "20px", "font-weight" : "800" , marginLeft: "20px", "color":"black", "background-color": "white"  }}>MENU</Link>
                    <Link to="/pastorders" style={{"font-size": "20px", "font-weight" : "800" , marginLeft: "20px","color":"black", "background-color": "white"  }}>PAST ORDERS</Link>
                    </div>
                <div class="ml-5">
                <div>{this.state.message}</div>
                <h1 className="cart">Your Order Cart:</h1>
                <div>hey</div>
                {details}
                </div>
                <h3>Total Amount: <span className="price" >${this.state.total_amount}</span></h3>
                <button className="btn-primary submit_btn" 
                onClick={this.placeorder}
                > Place Order</button>
            
            </div>
        )
    }
}

export default Order;