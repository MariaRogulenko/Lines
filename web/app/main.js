// @flow

"use strict"

import $ from "jquery";
import React from "react";
import ReactDOM from "react-dom";
import {Router, Route, browserHistory} from "react-router"
import injectTapEventPlugin from "react-tap-event-plugin";
import getMuiTheme from "material-ui/styles/getMuiTheme";
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";
import Login from "./Login";
import Play from "./Play";

// Executes the code below when the document is ready.
$(() => {
  // Needed for onTouchTap
  // Check this repo:
  // https://github.com/zilverline/react-tap-event-plugin
  injectTapEventPlugin();

  ReactDOM.render((
    <MuiThemeProvider muiTheme={getMuiTheme()}>
      <Router history={browserHistory}>
        <Route path="/" component={Login}/>
        <Route path="/play" component={Play}/>
      </Router>
    </MuiThemeProvider>
  ), document.getElementById("root"));
})
