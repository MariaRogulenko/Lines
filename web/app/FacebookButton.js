// @flow
 
import React from 'react';
import FacebookLogin from 'react-facebook-login';
 
export default class MyComponent extends React.Component {
    constructor(props: {params: Object}) {
        super(props);
    };

    responseFacebook = (response: Object) => {
        this.props.userChange(response.name, response.userID);
        this.props.x();
        
    };

    render() {
        return (
            <FacebookLogin
                appId="190987348065054"
                autoLoad={false}
                fields="name,email,picture"
                callback={this.responseFacebook}
                icon="fa-facebook"
                data-auto-logout-link="true" 
                size="small"
            />
        )
    }
}