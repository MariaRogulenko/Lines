// @flow
 
import React from 'react';
import {Group, Surface} from 'react-art';
import CircleComponent from './CircleComponent'
import {Motion, spring} from 'react-motion';


export default class BoardComponent extends React.Component {
    state: {
        opacity: number
    };

    constructor(props: {params: Object}) {
        super(props);
        this.state = {
            opacity: 1
        }
    };

    componentDidMount = () => {
        setInterval(() => {
        this.setState((prevState) => ({
            opacity: prevState.opacity === 1 ? 0 : 1
        }))
        }, 500)
    };

    convertToMatrix = () => {
        var rows = [];
        for (var i = 0; i < 9; i++) {
            var row = [];
            for (var j = 0; j < 9; j++) {
                row.push(this.props.board.table[i * 9 + j])
            }
            rows.push(row);
        }
        return rows;
    }

    handleClick = (i: number, j: number)=> {
        this.props.moveClick(i, j);
    }

    render() {
        var rows = this.convertToMatrix();
        var grid = rows.map((row, i) => {
            var rowElements = row.map((value, j) => {
                var radius = 30;
                const r = value & 1;
                const g = value & 2;
                const b = value & 4;
                const color = "#" + (r||"f") + (g||"f") + (b||"f");
                if (this.props.board.active.x !== -1 && i === (this.props.board.active.x || 0) && j ===  (this.props.board.active.y || 0) ) {
                    return(
                        <Motion defaultStyle={{ opacity: 0 }} style={{ opacity: spring(this.state.opacity) }}>
                        { (style) => <CircleComponent key = {'tale_' + j} x={i} y={j} radius={radius} color={color} moveClick={this.handleClick} style={style}/>}
                        </Motion>
                )} else {
                    return(
                    <CircleComponent key = {'tale_' + j} x={i} y={j} radius={radius} color={color} moveClick={this.handleClick} style={{opacity: 1}}/>
                );
                }
            });
            return <Group key={'row_' + i}> 
                    {rowElements}
                </Group>
        });
        return (
            
            <Surface width='540' height='540' style={{backgroundColor: 'whitesmoke'}}>
				<Group>
					{grid}
				</Group>
            </Surface>
			);
    }
}