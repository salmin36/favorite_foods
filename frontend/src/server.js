const { response} = require('express')
const express = require('express')
const app = express()
const port = 8080
const request = require("request");


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




app.listen(port, () => console.log(`Example app listening on port ${port}!`))