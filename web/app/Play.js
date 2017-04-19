// @flow
 
import React from 'react';
import BoardComponent from './BoardComponent'
import NextColors from './NextComponent'
import $ from "jquery";
import {browserHistory} from 'react-router';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';


export default class Play extends React.Component {

    state: {
        id: string,
        status: string,
        board: Object,
        username: string,
        bestScore: number,
        dialogOpen: boolean,
        aboutOpen: boolean,
    }

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            id: this.props.location.state ? this.props.location.state.userID : "",
            status: "",
            board: {},
            username: this.props.location.state ? this.props.location.state.username : "",
            bestScore: 0,
            dialogOpen: false,
            aboutOpen: false,
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

    handleAboutOpen = () => {
        this.setState({aboutOpen: true});
    };

    handleAboutClose = () => {
        this.setState({aboutOpen: false});
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

    handleRedirect() {
        browserHistory.push("/");
    }

    render() {
        if (!this.state.status || this.state.status === "NOT_FOUND") {
            return (
            <div>
                <Toolbar>
                    <ToolbarGroup firstChild={true}>
                    </ToolbarGroup>
                    <ToolbarGroup>
                        <ToolbarTitle text={"Color Lines"}  />
                    </ToolbarGroup>
                    <ToolbarGroup>
                    </ToolbarGroup>
                </Toolbar>
            <p>Error occurred. Please follow <a style= {{color: "blue"}} onClick = {this.handleRedirect}>this link</a> to start again</p>
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
        const actionsAbout = [
            <FlatButton
            label="Close"
            primary={true}
            onTouchTap={this.handleAboutClose}
        />,
        ];
        return (
            <div >
            <Toolbar>
                <ToolbarGroup firstChild={true}>
                    <FlatButton label="Game Rules" onTouchTap={this.handleAboutOpen}/>
                </ToolbarGroup>
                <ToolbarGroup>
                    <ToolbarTitle text={"Hello, " + this.state.username + "!"}  />
                </ToolbarGroup>
                <ToolbarGroup>
                    <FlatButton label="New Game" onTouchTap={this.handleNew}/>
                </ToolbarGroup>
            </Toolbar>
            <br/>


            <div style={{ width: 540, margin: 'auto', textAlign: 'center'}}>
            <NextColors board={this.state.board} radius={Math.min(window.innerWidth / 18, 30)}/>
            <p>{"High Score: " + (this.state.bestScore || 0)}</p>
            <p>{"Current Score: " + (this.state.board.score || 0)}</p>
            </div>
            <br/>
            <div>
            <BoardComponent style={{display: 'block', margin: 'auto'}} board={this.state.board} moveClick={this.click} radius={Math.min(window.innerWidth / 18, 30)}/>
            </div>
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
            <Dialog
                    title="Lines"
                    actions={actionsAbout}
                    modal={true}
                    open={this.state.aboutOpen}
                >
                <p>The game starts with a 9×9 board with five balls chosen out of seven different colours. 
                    The player can move one ball per turn, and the player may only move a ball to a particular place 
                    if there is a path (linked set of vertical and horizontal empty cells) between the current 
                    position of the ball and the desired destination. The goal is to remove balls by forming lines 
                    (horizontal, vertical or diagonal) of at least five balls of the same colour. 
                    If the player does form such lines of at least five balls of the same colour, 
                    the balls in those lines disappear, and he gains one turn, i.e. he can move another ball. 
                    If not, three new balls are added, and the game continues until the board is full.</p>
            </Dialog>
            </div>
        )
    }
}