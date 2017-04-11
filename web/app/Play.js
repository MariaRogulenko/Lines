// @flow

import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import {browserHistory} from 'react-router';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import React, {Component} from "react";
import RaisedButton from 'material-ui/RaisedButton';
import FlatButton from 'material-ui/FlatButton';
import CSSTransitionGroup from 'react-transition-group/CSSTransitionGroup';

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
    }


    render() {
        if (!this.state.status) {
            return (<AppBar title={"Game " + this.state.id} showMenuIconButton={false}/>);
        }
        let rows = [];
        for (let i = 0; i < 9; i++) {
            rows.push(this.state.board.table.slice(i*9, (i+1)*9));
        }
        const st = {
            borderStyle: 'dotted',
           
        } 
        const rowDivs = rows.map((row, idx) => {
            let cols = [];
            for (let j = 0; j < 9; j++) {
                const r = row[j] & 1;
                const g = row[j] & 2;
                const b = row[j] & 4;
                const color = "#" + (r||"f") + (g||"f") + (b||"f");
                if (this.state.board.active.x !== -1 && idx === (this.state.board.active.x || 0) && j === (this.state.board.active.y || 0)) {
                    cols.push(
                    <FloatingActionButton backgroundColor={color} style={st} key={j} onTouchTap={()=>this.click(idx, j)}/>
                )} else {
                    cols.push(
                    <FloatingActionButton backgroundColor={color} key={j} onTouchTap={()=>this.click(idx, j)}/>
                );
                }
                
                
            }
            return (
                <div key={idx}>
                    {cols}
                </div>
            );
        });
        let colors = this.state.board.next_colors;
        let cols2 = [];
            for (let j = 0; j < 3; j++) {
                const r = colors[j] & 1;
                const g = colors[j] & 2;
                const b = colors[j] & 4;
                const color = "#" + (r||"f") + (g||"f") + (b||"f");
                cols2.push(
                        <FloatingActionButton backgroundColor={color} mini={true} key={j}/>
                );
            }
        const colorsDivs = 
                <div >
                    {cols2}
                </div>
               
        return (
            <div>
                <AppBar title={"Hello, " + this.state.username + "!"} showMenuIconButton={false}/>
                <RaisedButton label={"High Score: " + (this.state.bestScore || 0)} disabled={true} />
                <RaisedButton label={"Curr Score: " + (this.state.board.score || 0)} disabled={true} />
                {colorsDivs}
                <br />
                <RaisedButton label="Start New Game"  onTouchTap={this.handleNewGame}/>
                {rowDivs}
            </div>
        );
    }
}
