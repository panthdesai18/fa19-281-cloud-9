import React, {Component} from 'react';
import "./locations.css";
import {Link} from 'react-router-dom';
import axios from 'axios';
import {Redirect} from 'react-router';



class Locations extends Component {
    constructor(props){
        super(props);
        this.state = {
            location : "",
            locationInfo: []
        }
        this.onChangeLocation = this.onChangeLocation.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.settingLocation = this.settingLocation.bind(this);
    }

    onChangeLocation (e) {
        e.preventDefault();
        this.setState({[e.target.name] : e.target.value})
    }
    async onSubmit(e){
        e.preventDefault();
        const locationData = this.state.location
        try{
            const connectionReqResponse = await axios.get(`http://localhost:3000/locations/zipcode/${locationData}`)
            console.log("checkpoint")
            console.log(connectionReqResponse.data)
            if (connectionReqResponse.data === null) {
                this.setState({
                    locationInfo: [],
                    nolocation: "No near locations found"
                })
            }
            else {

            this.setState({
                locationInfo : this.state.locationInfo.concat(connectionReqResponse.data)
            })
            }
            console.log(this.state.locationInfo)

        } catch(err) {

        }

    }

    settingLocation(locationName) {
        console.log("this")
        localStorage.setItem("locationName",locationName)
        this.props.history.push("/menu")
    }

    render(){
        let redirectVar = null;
        //localStorage.setItem("user","ramya")
        if(!localStorage.getItem("user")){
            redirectVar = <Redirect to= "/login"/>
        }
        if (this.state.locationInfo.length > 0) {
            var details = this.state.locationInfo.map((value, i) => {
                return (
                    <div>

                        <div id="borderDemo">
                            <br></br>
                            <table>
                                <br></br>
                                <tr>
                                    <button type="button" onClick={() => this.settingLocation(value.locationName)}
                                            class="myButton">{i + 1}. {value.locationName}</button>
                                </tr>
                                <br></br>
                                <tr>
                                    <tr>Restaurant name :- {value.locationName}</tr>
                                    <tr>Address :- {value.address}, {value.city}, {value.country}, {value.zipcode}</tr>
                                    <tr>Phone:- {value.phone}</tr>
                                    <tr>Email:- {value.email}</tr>
                                </tr>

                            </table>
                        </div>
                    </div>
                )

            })
        }
        else {
            var details =  <p> Not Found </p>
        }
        return(
            <div>
                {redirectVar}
                <div className="location">
                    <div className="locationNav">
         <span className="locationSpan"> <a href="/" id="Logo"></a>
         </span>
                        <div id="menu-outer-location">
                            <div className="tableLocation">
                                <ul id="horizontal-list-location">
                                    <li><Link to="/home"><font color = "black" face="Impact" size="4">HOME</font></Link></li>
                                    <li><Link to="/menu"><font color = "black" face="Impact" size="4">VIEW MENUS</font></Link></li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="locationBox">
                    <h1 className="headFont">
                        Locations of our Hungry Burgers Near You
                    </h1>
                    <h2 >
                        Enter Zip Code:
                    </h2>
                    <form action="/locations" id="FORM_8" onSubmit = {this.onSubmit}  >
                        <input type="text" id="zipText" name="location" placeholder="enter zip code" onChange = {this.onChangeLocation} />
                        <br></br>
                        <input type="submit" className="searchButton" value="Ok"/>
                    </form>
                    <br></br>

                    {details}

                </div>

            </div>
        )

    }

}
export default Locations;