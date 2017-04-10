// @flow

import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import {browserHistory} from 'react-router';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import React, {Component} from "react";
import RaisedButton from 'material-ui/RaisedButton';
import FlatButton from 'material-ui/FlatButton';

export default class Play extends Component {
    state: {
        id: string,
        status: string,
        board: Object,
        username: string,
        bestScore: number
    }

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            id: this.props.location.state.userID,
            status: "",
            board: {},
            username: this.props.location.state.username,
            bestScore: 0
        };
        this.componentWillMount();
    }

    componentWillMount = () => {
        $.get("/api/v1/state/" + this.state.id, (result) => {
            this.setState({
                status: result.status,
                board: result.board,
                bestScore: result.best_score
            });
        });
        this.zero();
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
                    bestScore: result.state.best_score
                });
            },
        });
        this.zero();
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
                    bestScore: result.state.best_score
                });
            },
        });
        this.zero();
    }

    zero = () => {
        if ("undefined" === typeof this.state.bestScore) {
            this.state.bestScore = 0;
        }
        if ("undefined" === typeof this.state.board.score) {
            this.state.board.score = 0;
        }
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
                <RaisedButton label={"High Score: " + (this.state.bestScore || 0)} disabled={true} />
                <RaisedButton label={"Curr Score: " + (this.state.board.score || 0)} disabled={true} />
                <br />
                <RaisedButton label="Start New Game" onTouchTap={this.handleNewGame}/>
                {rowDivs}
            </div>
        );
    }
}
