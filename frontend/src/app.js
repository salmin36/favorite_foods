import React, {Component} from 'react';
import ReactDOM from 'react-dom';

class App extends Component{
    
    componentDidMount(){
        this.fetchDataFromBackend();
    }

    constructor(props) {
        super(props);
        this.state = {data: 'false'}
    
    }
    
    fetchDataFromBackend(){
        fetch('/foods', { method: 'get', mode: 'no-cors', })
        .then(res => {
            if (res.ok) {
                return res
            } else {
                let errorMessage = '${res.status(${res.statusText})}',
                error = new Error(errorMessage);
                console.log(errorMessage)
                throw(error)
            }
        })
        .then(res => res.json())
        .then(json =>{
            console.log(json);
            this.setState({data: json.data})
        })
    }

    render (){
        return( 
        <div className="App">
            <header>
                <p>Hello World!</p>
            </header>
        </div>);
    
    }
}

export default App;