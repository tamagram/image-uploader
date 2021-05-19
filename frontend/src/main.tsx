import React from "react"
import { render } from "react-dom"

const App: React.FC = () => {
    return (
        <>
            <h1>Hello, world</h1>
            <p>I'm tama</p>
        </>
    )
}

render(<App />, document.getElementById('root'))