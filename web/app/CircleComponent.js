// @flow
 
import React from 'react';
import {Shape, Path} from 'react-art';


export default class CircleComponent extends React.Component {

    constructor(props: {params: Object}) {
        super(props);
    };

    makePath = (radius: number , cx: number, cy: number) => {
        var path = Path().moveTo(cx, cy-radius).arc(0, radius * 2, radius)
        .arc(0, radius * -2, radius).close();
        return path;
    }

    handleClick = () => {
        this.props.moveClick(this.props.x, this.props.y);
    }

    render() {
        var radius = this.props.radius;
        var cx = this.props.x * 2 * radius + radius;
        var cy = this.props.y * 2 * radius + radius;
        var path = this.makePath(radius, cx, cy);
        return (
                <Shape 
                    d={path} 
                    fill={this.props.color} 
                    onClick={this.handleClick} 
                    opacity={this.props.style.opacity}>
                </Shape>
        )
    }
}