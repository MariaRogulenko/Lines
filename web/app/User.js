// @flow

import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import {browserHistory} from 'react-router';
import RaisedButton from 'material-ui/RaisedButton';
import React, {Component} from "react";

export default class User extends Component {
    state: {
        username: string,
        userID: string
    }

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            username: props.params.username,
            userID: this.props.location.state.userID
        };
        console.log(this.state.username, this.state.userID, "!");
    }

    handleNewGame = () => {
        $.ajax({
            type: "POST",
            url: "/api/v1/new",
            data: JSON.stringify({user_name: this.state.username, id: this.state.userID}),
            success: (result) => {
                console.log(result.id);
                const location = {
                    pathname: "/play/" + this.state.username + "/",
                    state: {
                        userID: result.id
                    }
                }
                browserHistory.push(location);
            },
        });
    }

    render() {
        return (
            <div>
                <AppBar title={"Hello, " + this.state.userID + "!"} showMenuIconButton={false}/>
                <RaisedButton label="New game" onTouchTap={this.handleNewGame}/>
            </div>
        );
    }
}
