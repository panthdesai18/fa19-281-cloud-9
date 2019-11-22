import React, {Component} from 'react';
import "./signup.css";
import {Link} from 'react-router-dom';
import axios from 'axios';



class SignUp extends Component {
    constructor(props){
        super(props);
        this.state = {
            uname : "",
            fname : "",
            emailid : "",
            password : ""
        }
        this.onChangeSignUp = this.onChangeSignUp.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }
    onChangeSignUp (e) {
        e.preventDefault();
        this.setState({[e.target.name] : e.target.value})
    }

    async onSubmit(e){
        e.preventDefault();
        const signUpData = {
            Username: this.state.uname,
            Fullname: this.state.fname,
            Emailid : this.state.emailid,
            Password : this.state.password
        }
        try{
            const connectionReqResponse = await axios.post('http://login-env.u67gpznbsg.us-east-1.elasticbeanstalk.com/signup', signUpData)
            if (connectionReqResponse.status === 201){
                alert("User has been successfully created!");
                this.props.history.push("/login");
            }
        } catch(err) {
            if (err.response.status === 409){
                alert("User already exists!")
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

                <div className="signupbox">
                    <h1 className = "signupheading">SIGNUP</h1>
                    <center>Already have an account? <Link to="/login"> LOGIN!</Link></center>
                    <form onSubmit = {this.onSubmit}>
                        <div className = "signUpDiv">
                            <label for="Username" className="signUpLabel">
                                USERNAME
                            </label>
                            <input type="text" name = "uname" onChange = {this.onChangeSignUp} className="signUpInput" placeholder="Username" required/>
                        </div>

                        <div className = "signUpDiv">
                            <label for="Fullname"  className="signUpLabel">
                                FULL NAME
                            </label>
                            <input type="text" name = "fname" onChange = {this.onChangeSignUp} className="signUpInput" placeholder="Full Name"required/>
                        </div>

                        <div className = "signUpDiv">
                            <label for="emailid" className="signUpLabel">
                                EMAIL ID
                            </label>
                            <input type="email" name = "emailid" onChange = {this.onChangeSignUp} className="signUpInput" placeholder="Email id"required/>
                        </div>

                        <div className = "signUpDiv">
                            <label for="Password" className="signUpLabel">
                                PASSWORD
                            </label>
                            <input type="password" name = "password" onChange = {this.onChangeSignUp} className="signUpInput" placeholder="Password"required/>
                        </div>

                        <div className = "signUpDiv">
                            <input type="submit" className="signUpButton" value="Sign Up"/>
                        </div>
                    </form>
                </div>
            </div>
        )
    }
}
export default SignUp;
