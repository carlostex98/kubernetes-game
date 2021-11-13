//imports
const express = require('express')
const { MongoClient } = require('mongodb');
const app = express()
const port = 2500;

const url = 'mongodb://root:example@localhost:27017/';
const client = new MongoClient(url);
const dbName = 'Proyecto2Sopes';

client.connect();
console.log('Connected successfully to server');
const db = client.db(dbName);
const collection = db.collection('Logs');


//routes
app.get('/', (req, res) => {
    res.json({"id":'ct'})
})

app.get('/dashboard', async (req, res) => {
    //calcular reportes
    const f = await collection.find({}).toArray();
    //console.log(f)
   res.json(f);
})

//app listen
app.listen(port, async () => {
   
    
    console.log(`Example app listening at http://localhost:${port}`)
})

//helpers
