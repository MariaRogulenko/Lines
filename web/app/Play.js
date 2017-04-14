// @flow
 
import React from 'react';
import BoardComponent from './BoardComponent'
import NextColors from './NextComponent'
import $ from "jquery";
import AppBar from 'material-ui/AppBar';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';

export default class Play extends React.Component {

    state: {
        id: string,
        status: string,
        board: Object,
        username: string,
        bestScore: number,
        dialogOpen: boolean,
    }

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            id: this.props.location.state.userID,
            status: "",
            board: {},
            username: this.props.location.state.username,
            bestScore: 0,
            dialogOpen: false,
        };
    }

    componentWillMount = () => {
        $.get("/api/v1/state/" + this.state.id, (result) => {
            console.log(result.status);
            this.setState({
                status: result.status,
                board: result.board,
                bestScore: result.best_score
            });
        });
    }

    handleDialogOpen = () => {
        this.setState({dialogOpen: true});
    };

    handleDialogClose = () => {
        this.setState({dialogOpen: false});
    };

    handleNew = () => {
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
                    bestScore: result.state.best_score,
                    dialogOpen: result.state.status === "GAME_OVER" ? true : false
                });
            },
        });
    }

    render() {
        if (!this.state.status || this.state.status === "NOT_FOUND") {
            return (
            <div>
            <AppBar title={"Game " + this.state.id} showMenuIconButton={false}/>
            <p>User is not found. Try to reload the page</p>
            </div>);
        }
        const actions = [
        <FlatButton
            label="Yes"
            primary={true}
            onTouchTap={(event) => { this.handleNew(); this.handleDialogClose();} }
        />,
        <FlatButton
            label="Cancel"
            primary={true}
            onTouchTap={this.handleDialogClose}
        />,
        ];
        return (
            <div>
            <AppBar title={"Hello, " + this.state.username + "!"} showMenuIconButton={false}/>
            <RaisedButton label="Start New Game"  onTouchTap={this.handleNew}/>
            <RaisedButton label={"High Score: " + (this.state.bestScore || 0)} disabled={true} />
            <RaisedButton label={"Curr Score: " + (this.state.board.score || 0)} disabled={true} />
            <br/>
            <NextColors board={this.state.board}/>
            <br/>
            <BoardComponent board={this.state.board} moveClick={this.click}/>
            <Dialog
                    title="Game Over"
                    actions={actions}
                    modal={true}
                    open={this.state.dialogOpen}
                >
                <p>Score: {this.state.board.score || 0}</p>
                <p>Best: {this.state.bestScore || 0}</p>
                <p>Would you like to start new game?</p>
            </Dialog>
            </div>
        )
    }
}