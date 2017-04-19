// @flow

import $ from "jquery";
import {browserHistory} from 'react-router';
import RaisedButton from 'material-ui/RaisedButton';
import React, {Component} from "react";
import TextField from 'material-ui/TextField';
import FacebookButton from './FacebookButton';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
import Paper from 'material-ui/Paper';

export default class Login extends Component {
    state: {
        username: string,
        userID: string,
        dialogOpen: boolean,
        errorText: string
    };

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            username: "",
            userID: "",
            dialogOpen: false,
            errorText: ""
        };
    }

    handleUsernameChange = (e: Object, value: string) => {
        this.setState({username: value});
    }

    handleUserChange = (name: string, id: string) => {
        this.setState({username: name, userID: id});
    }

    handleLogin = () => {
        $.ajax({
            type: "POST",
            url: "/api/v1/login",
            data: JSON.stringify({user_name: this.state.username, id: this.state.userID}),
            success: (result) => {
                this.setState({userID: result.id});
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

    handleDialogOpen = () => {
        this.setState({dialogOpen: true});
    };

    handleDialogClose = () => {
        this.setState({dialogOpen: false});
    };

    handleError = () => {
        if (this.state.username === "") {
            this.setState({
            errorText: "This field is required"
            })
        } else {
            this.handleDialogClose();
            this.handleLogin()
        };
    }

    render() {

        const style = {
            height: 210,
            width: 300,
            margin: 20,
            textAlign: 'center',
            display: 'inline-block',
        };

        const actions = [
        <FlatButton
            label="Continue"
            onTouchTap={this.handleError}
            primary={true}
        />,
        <FlatButton
            label="Cancel"
            onTouchTap={ this.handleDialogClose}
            secondary={true}
        />,
        ];
        return (
            <div>
                <Toolbar>
                    <ToolbarGroup>
                </ToolbarGroup>
                <ToolbarGroup>
                    <ToolbarTitle text={"Color Lines"}  />
                </ToolbarGroup>
                <ToolbarGroup>
                </ToolbarGroup>
                </Toolbar>
            <div style={{textAlign: "center", margin: "auto"}}>
                <div style={{width: 600, display: "inline-block"}}>
                <p>Color Lines is a web version of the once popular desktop game. 
                   The object of the game is to align as often as possible five (or more) balls 
                   of the same color causing them to disappear.
                </p>
                </div>

                <br/>
                <Paper 
                    style={style} 
                    zDepth={2} 
                    children={
                    <div>
                        <div style={{marginTop: 20, marginBottom: 40}}>
                        <FacebookButton userChange={this.handleUserChange} x={this.handleLogin}/>
                        </div>
                        <p>or</p>
                        <RaisedButton 
                            style={{marginTop: 20}} 
                            label="Play without registration" 
                            onTouchTap={(event) => {this.handleDialogOpen()}}
                        />
                    </div>}
                />
            </div>
                <Dialog
                    actions={actions}
                    modal={true}
                    open={this.state.dialogOpen}>
                <p>Please, enter your name</p>
                <TextField errorText={this.state.errorText} onChange={this.handleUsernameChange}/>
            </Dialog>
            </div>
        );
    }
}
