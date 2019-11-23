import React, {Component} from 'react';
import "./signup.css";
import {Link} from 'react-router-dom';
import axios from 'axios';


class Login extends Component {
    constructor(props){
        super(props);
        this.state = {
            uname : "",
            password : ""
        }
        this.onChangeLogin = this.onChangeLogin.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }
    onChangeLogin (e) {
        e.preventDefault();
        this.setState({[e.target.name] : e.target.value})
    }
    async onSubmit(e){
        e.preventDefault();
        console.log("in submit")
        const loginData = {
            Username : this.state.uname,
            Password : this.state.password
        }
        console.log(loginData);

        try{
            const connectionReqResponse = await axios.post('https://gevnsiba07.execute-api.us-east-1.amazonaws.com/prod/usermanagement/login', loginData)
            console.log(connectionReqResponse);
            if (connectionReqResponse.status === 200){
                if(connectionReqResponse.data.error){
                    alert(connectionReqResponse.data.error)
                }else{
                    alert("Login successful!"); 
                    localStorage.setItem('username', connectionReqResponse.data.username);
                    localStorage.setItem('id', connectionReqResponse.data.username)
                    localStorage.setItem('name', connectionReqResponse.data.fullname)
                    this.props.history.push("/location");
                }
                
            }
        }
        catch(err) {
            if (err.response.status === 401){
                alert(err.response.data.Message)
            }
        }
    }
    render(){
        return(
            <div>
                <div className="signup">
                    <div className="signupNav">
                        <span className="signupSpan"> <a href="/" id="A_4"></a></span>
                    </div>
                </div>


                <div className="loginbox">
                    <h1 className = "signupheading">LOGIN</h1>
                    <center>Don't have an account?<Link to="/signup"> Create One!</Link> </center>
                    <form onSubmit = {this.onSubmit}>
                        <div className = "signUpDiv">
                            <label for="email" className="signUpLabel">
                                Username
                            </label>
                            <input type="text" name = "uname" onChange = {this.onChangeLogin} class="signUpInput"/>
                        </div>

                        <div className = "signUpDiv">
                            <label for="Password" className="signUpLabel">
                                Password
                            </label>
                            <input type="password" name = "password" onChange = {this.onChangeLogin} className="signUpInput"/>
                        </div>

                        <div className = "signUpDiv">
                            <input type="submit" className="signUpButton"  value="Login"/>
                        </div>
                    </form>
                </div>
            </div>
        )
    }


}
export default Login;