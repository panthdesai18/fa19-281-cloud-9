import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import SignUp from './Login/Signup'
import Login from './Login/Login'
import LandingPage from './LandingPage/LandingPage'
import Menu from './Menu/Menu';
import Order from './Order/Order';
import Pastorders from './Order/Pastorders';
import Locations from './Locations/Locations';

class Main extends Component {
    render(){
        return(
            <div>
             <Route path="/signup" component={SignUp}/>
             <Route path="/login" component={Login}/>
             <Route path="/home" component={LandingPage}/>
             <Route path="/menu" component={Menu}/>
             <Route path="/burgerOrder" component={Order} />
             <Route path="/pastorders" component={Pastorders} />
             <Route path="/location" component={Locations}/>
            </div>
        )

    }
}
export default Main;
