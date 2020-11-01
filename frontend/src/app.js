import React, {Component} from 'react';
import {
    TaggedContentCard
  } from 'react-ui-cards';

class App extends Component{
    
    componentDidMount(){
        this.fetchDataFromBackend();
    }

    constructor(props) {
        super(props);
        this.state = {
            foods : []
        }
    
    }
    
    fetchDataFromBackend(){
        fetch('/foods', { method: 'get', mode: 'no-cors', })
        .then(res => {
            if (res.ok) {
                return res
            } else {
                console.log(res)
                
            }
        }).catch( err => {
            console.err(err)
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
            <TaggedContentCard
          //href='https://github.com/nukeop'
          thumbnail='https://cdn.valio.fi/mediafiles/bcf5ebd5-a1c9-435f-a5cd-6b1692adcd13/1000x752-recipe-hero/4x3/kinkkukiusaus.jpg'
          title='Kinkkukiusaus'
          description='Kinkkukiusaus'
          tags={[
            'Liha',
            'kiusaus'
          ]}
        />
        </div>);
    
    }
}

export default App;