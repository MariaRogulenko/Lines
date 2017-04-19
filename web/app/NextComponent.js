// @flow
 
import React from 'react';
import {Group, Surface} from 'react-art';
import CircleComponent from './CircleComponent'


export default class NextComponent extends React.Component {
    constructor(props: {params: Object}) {
        super(props);
    };

    render() {
        var row = this.props.board.next_colors.map((value, i) => {
                var radius = this.props.radius;
                const r = value & 1;
                const g = value & 2;
                const b = value & 4;
                const color = "#" + (r||"f") + (g||"f") + (b||"f");
                return <CircleComponent key = {'next_' + i} x={i} y={0} radius={radius} color={color} style={{opacity: 1}}/>
        });
        return (
            
        <Surface width={this.props.radius * 6} height={this.props.radius * 2} style={{ display: 'block', margin: 'auto'}}>
            <Group>
                {row}
            </Group>
        </Surface>
        );
    }
}