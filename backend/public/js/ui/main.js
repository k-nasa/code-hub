import { h, render, Component } from 'preact';

class Clock extends Component {
    render() {
        let time = new Date().toLocaleTimeString();
        return h('span', null, time);
    }
}

export function renderMain(element) {
    render(h(Clock),element)
}
