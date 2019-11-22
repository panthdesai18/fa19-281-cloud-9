import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import SignUp from './User/SignUp'
import Login from './User/Login'
import LandingPage from './LandingPage/LandingPage'
import Menu from './Menu/Menu';
import Payments from './Payments/Payments';
import Order from './Order/Order';
import Location from './Location/Location';

class Main extends Component {
    render(){
        return(
            <div>
             <Route path="/signup" component={SignUp}/>
             <Route path="/login" component={Login}/>
             <Route path="/home" component={LandingPage}/>
             <Route path="/menu" component={Menu}/>
             <Route path="/payments" component={Payments}/>
             <Route path="/burgerOrder" component={Order} />
             <Route path="/location" component={Location}/>
            </div>
        )

    }
}
export default Main;
