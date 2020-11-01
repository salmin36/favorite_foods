const { response} = require('express')
const express = require('express')
const app = express()
const port = 8080
const request = require("request");

/*
var whitelist = ['http://backend', "https://backend", "https://frontend", "http://127.0.0.1", "https://127.0.0.1"]
var corsOptions = {
  origin: function (origin, callback) {
    if (whitelist.indexOf(origin) !== -1) {
      callback(null, true)
    } else {
      callback(new Error('Not allowed by CORS'))
    }
  }
}

res.
app.use(cors(corsOptions))
*/

app.use('/', express.static('./dist', {
  index: "index.html"
}))


app.get('/foods', function(req, res) {
  let json = ""
  request.get('http://backend:8081/foods', (error, response, body) => {
    if (!error){
      json = JSON.parse(body)
      console.log(json)
      res.json(json)
    } else {
      console.error(error)
      res.json("{'Error':'Could not resolve host'}")
    }
    
  })
})



/*request('http://backend:8081/foods',function(error, response, body){
    if(error.body != null) {
      console.error("Error happened")
      res.json("{'Error':'Could not find endpoint'}")
    } else {
      res.send(response)
    }  
    
  })
  */ 
/*

app.get('/foods', async (req, res) => {
  const uri = 'http://backend:8081/foods'
  const resp =  await reque(uri).catch(r => {
    
    res.json("{'error':'service not found'}")
  })
  res.json(resp)
  

})
*/

/*
app.use('/foods', createProxyMiddleware({
  target: 'http://backend:8081/foods',
  
  }));
*/

app.listen(port, () => console.log(`Example app listening on port ${port}!`))