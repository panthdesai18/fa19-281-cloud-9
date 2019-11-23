import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import Signup from './Login/Signup'
import Login from './Login/Login'
import AdminLogin from './Login/AdminLogin'
import LandingPage from './LandingPage/LandingPage'
import Menu from './Menu/Menu';
import Locations from './Locations/Locations';
class Main extends Component {
    render(){
        return(
            <div>
             <Route path="/signup" component={Signup}/>
             <Route path="/login" component={Login}/>
             <Route path="/adminlogin" component={AdminLogin}/>
             <Route path="/home" component={LandingPage}/>
             <Route path="/menu" component={Menu}/>
             <Route path="/locations" component={Locations}/>

            </div>
        )

    }
}
export default Main;
