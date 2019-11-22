import React, {Component} from 'react';
import {Link} from 'react-router-dom';


class LandingPage extends Component {
    render(){
    
        return(
            <div>
            <div className="signup"> 
            <div className="signupNav">
         <span className="signupSpan"> <a href="/" id="A_4"></a>
         </span>
         <div id="menu-outer">
  <div className="tableLangingPage">
    <ul id="horizontal-list">
        <li><Link to="/home"><font color = "black" face="Impact" size="4">HOME</font></Link></li>
      <li><Link to="/signup"><font color = "black" face="Impact" size="4">CREATE ACCOUNT</font></Link></li>
      <li><Link to="/login"><font color = "black" face="Impact" size="4 ">SIGN IN</font></Link></li>
    </ul>
  </div>
</div>  
 </div>
  </div>
<div id="DIV_3">
 </div>
 </div>
)
    }
}
export default LandingPage;

