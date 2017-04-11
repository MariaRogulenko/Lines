// @flow

import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import {browserHistory} from 'react-router';
import RaisedButton from 'material-ui/RaisedButton';
import React, {Component} from "react";
import TextField from 'material-ui/TextField';
import FacebookButton from './FacebookButton';

export default class Login extends Component {
    state: {
        username: string,
        userID: string
    };

    //handleUsernameChange = (e: Object, value: string) => {
    //    this.setState({username: value});
    //}

    //handleLogin = () => {
    //    browserHistory.push("/user/" + this.state.username + "/");
    //}

    handleNew = (name: string, id: string) => {
        this.setState({username: name, userID: id});
        $.ajax({
            type: "POST",
            url: "/api/v1/login",
            data: JSON.stringify({user_name: this.state.username, id: this.state.userID}),
            success: (result) => {
                const location = {
                    pathname: "/play/",
                    state: {
                        username: this.state.username,
                        userID: this.state.userID
                    }
                }
                browserHistory.push(location);
            },
        });
    }
    //<TextField hintText="Enter your name" onChange={this.handleUsernameChange}/><br/>
    //            <RaisedButton label="Login" onTouchTap={this.handleLogin}/>
    render() {
        return (
            <div>
                <AppBar title="Lines" showMenuIconButton={false}/>
                <FacebookButton ha={this.handleNew}/>
                
                
            </div>
        );
    }
}
