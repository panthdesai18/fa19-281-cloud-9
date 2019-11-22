import React, { Component } from 'react';
import counterburgersymbol from './counterburgersymbol.png';
import cbsymbol from './cbsymbol.jpg';
import burgerdetails from './burgerdetails.png';
import { Link } from 'react-router-dom';
import axios from 'axios';
import { Redirect } from 'react-router';
//import {Redirect} from 'react-router';
import './Menu.css';
import { Card } from 'antd';
//const { Meta } = Card;
var swal = require('sweetalert')
var hostname = 'http://18.206.192.130:12345/items'
var hostnameOrder = 'http://kong-elb-234657806.us-west-1.elb.amazonaws.com:80/order/order'
class Menu extends Component {

    constructor(props) {
        super(props);
        this.state = {
            allmenu: [],
            redirectVar: null,
            itemdetail: [],
            price: "",
            item:[],
            total : 0,
            itemname: []
        }
        this.handleclick = this.handleclick.bind(this);
       
    }


    handleclick = (e, name, price) => {
         this.state.item = [name, price]
         console.log(this.state.item)
        this.state.itemdetail.push(this.state.item);
        this.state.itemname.push(name);
        console.log("state",this.state.itemdetail, this.state.total, this.state.itemname)
    }

    componentDidMount() {
        console.log("hi")
        axios.get(hostname)
            .then((response) => {
                console.log("Response data", response.data)
                this.setState({
                    allmenu: response.data
                })
            });
        console.log("Checking menu details", this.state.allmenu)
    }

    render() {
        console.log("hey",this.state.itemdetail)
        let redirectVar = null;
        if (!localStorage.getItem("user")) {
            redirectVar = <Redirect to="/home" />
        }
        let wholemenu = this.state.allmenu.map((wholemenu, j) => {
            return (
                <div className="Menu">
                    <Card
                        hoverable
                        style={{ width: 300 }}
                        cover={<img src={cbsymbol} height="320" width="550" alt=""></img>}
                        className="MenuCards"
                    >
                        <div className="ItemDescription">
                            <br></br>
                            <p><b>Item Type : </b>{wholemenu.ItemType}</p>
                            <p><b>Item Name : </b>{wholemenu.ItemName}</p>
                            <p><b>Description :</b> {wholemenu.Description}</p>
                            <p><b>Price : </b>{wholemenu.Price}$</p>
                        </div>
                
                        <button onClick={(e) => {
                                                this.state.total = this.state.total + parseInt(wholemenu.Price)
                                                this.handleclick(e, wholemenu.ItemName, wholemenu.Price )
                                            }} className="btn btn-danger cartButton ">Add to Cart</button>
                    </Card>
                </div>
            )
        })
        return (
            <div>

                <div className="backgroundwallimage">
                    <div className="counterburgersymbol">
                        <img src={counterburgersymbol} height="100" width="200" alt=""></img>
                        <div className="NavbarLinks">
                            &nbsp; &nbsp; <Link to="/home" style={{ "font-size": "20px", "font-weight": "800", "color": "black", "background-color": "white" }}>HOME</Link>
                            <Link to="/menu" style={{ "font-size": "20px", "font-weight": "800", marginLeft: "20px", "color": "black", "background-color": "white" }}>MENU</Link>
                            <button style={{ "font-size": "20px", "font-weight": "800", marginLeft: "20px", "color": "black", "background-color": "white" }} 
                            onClick={() => {
                                                try {
                                                    console.log("state",this.state.itemdetail)
                                                    this.props.history.push({
                                                        pathname: "/burgerOrder",
                                                        state: {
                                                            itemdetail: this.state.itemdetail,
                                                            total: this.state.total,
                                                            itemname: this.state.itemname
                                                        }
                                                    })
                                                } catch (e) { }
                                            }}><Link to='/burgerOrder'>CART</Link></button>
                        </div>
                        <div className="container MenuOustide">
                            <div className="storedetails">
                                &nbsp;&nbsp; <b style={{ "font-size": "40px", "font-weight": "800", marginBottom: "0px" }}>THE COUNTER</b>
                                <br></br>
                                &nbsp;&nbsp; <Link to="/location">Change Location</Link>
                                <br></br>
                                <p>&nbsp;&nbsp; Phone: (408) 423-9200</p>
                                <p> &nbsp;&nbsp; Pickup Hours: Open today 11am-10pm </p>
                                <p>&nbsp;&nbsp; Accepted Cards: Mastercard, American Express, Discover</p>

                                <div className="burgerdetails">
                                    <img src={burgerdetails} height="220" width="550" alt=""></img>
                                    <p style={{ "font-size": "15px", "font-weight": "400", marginLeft: "5px", marginBottom: '5px' }}>Selections vary by location and may have limited availability.<br></br>
                                        <a className="allergy" href="/nutrition">Nutritional information<br></br>
                                            Allergen information</a>
                                    </p>

                                </div>
                            </div>
                            <div className="container Menu2">
                                {wholemenu}
                                <br></br>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Menu;