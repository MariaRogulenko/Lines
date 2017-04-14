// @flow
 
import React from 'react';
import FacebookLogin from 'react-facebook-login';
 
export default class MyComponent extends React.Component {
    constructor(props: {params: Object}) {
        super(props);
    };

    responseFacebook = (response: Object) => {
        this.props.FBLogin(response.name, response.userID);
    };

    render() {
        return (
            <FacebookLogin
                appId="190987348065054"
                autoLoad={false}
                fields="name,email,picture"
                callback={this.responseFacebook}
                data-auto-logout-link="true"
            />
        )
    }
}