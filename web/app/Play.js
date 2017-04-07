// @flow

import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import {browserHistory} from 'react-router';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import React, {Component} from "react";
import RaisedButton from 'material-ui/RaisedButton';

export default class Play extends Component {
    state: {
        id: string,
        status: string,
        board: Object,
        username: string
    }

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            id: this.props.location.state.userID,
            status: "",
            board: {},
            username: this.props.location.state.username
        };
        this.componentWillMount();
    }

    componentWillMount = () => {
        $.get("/api/v1/state/" + this.state.id, (result) => {
            this.setState({
                status: result.status,
                board: result.board,
            });
        });
    }

    handleNewGame = () => {
        $.ajax({
            type: "POST",
            url: "/api/v1/new",
            data: JSON.stringify({id: this.state.id}),
            success: (result) => {
                this.setState({
                    status: result.state.status,
                    board: result.state.board,
                });
            },
        });
    }

    click = (i: number, j: number) => {
        $.ajax({
            type: "POST",
            url: "/api/v1/move/" + this.state.id,
            data: JSON.stringify({to: {x : i, y : j}}),
            success: (result) => {
                this.setState({
                    status: result.state.status,
                    board: result.state.board,
                });
            },
        });
    }

    render() {
        if (!this.state.status) {
            return (<AppBar title={"Game " + this.state.id} showMenuIconButton={false}/>);
        }
        let rows = [];
        for (let i = 0; i < 9; i++) {
            rows.push(this.state.board.table.slice(i*9, (i+1)*9));
        }
        const rowDivs = rows.map((row, idx) => {
            let cols = [];
            for (let j = 0; j < 9; j++) {
                const r = row[j] & 1;
                const g = row[j] & 2;
                const b = row[j] & 4;
                const color = "#" + (r||"f") + (g||"f") + (b||"f");
                cols.push(
                    <FloatingActionButton backgroundColor={color} key={j} onTouchTap={()=>this.click(idx, j)}/>
                );
            }
            return (
                <div key={idx}>
                    {cols}
                </div>
            );
        });
        return (
            <div>
                <AppBar title={"Hello, " + this.state.username + "!"} showMenuIconButton={false}/>
                {JSON.stringify(this.state.status)}
                <RaisedButton label="Start New Game" onTouchTap={this.handleNewGame}/>
                {rowDivs}
            </div>
        );
    }
}
