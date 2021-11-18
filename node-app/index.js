//imports
const express = require('express')
const { MongoClient } = require('mongodb');
const app = express()
const port = 2500;

const url = 'mongodb://root:example@35.184.26.14:27017/';
const client = new MongoClient(url);
const dbName = 'Proyecto2Sopes';

client.connect();
console.log('Connected successfully to server');
const db = client.db(dbName);
const collection = db.collection('Logs');


//routes
app.get('/', (req, res) => {
    res.json({ "id": 'ct' })
})

app.get('/dashboard', async (req, res) => {
    //calcular reportes
    const f = {
        "data_mongo": await allData(),
        "worker_report": await workerReport(),
        "top_3": await top3()
    }
    //console.log(f)
    res.json(f);
})

app.get('/redis', async (req, res) => {
    //calcular reportes
    const f = {
        "latest": await getLast10(),
        "best": await best10()
    }
    //console.log(f)
    res.json(f);
})

app.get('/user/:id', async (req, res) => {
    //calcular reportes
    const f = await userReport(req.params.id);
    //console.log(f)
    res.json(f);
})

//app listen
app.listen(port, async () => {

    console.log(`Example app listening at http://localhost:${port}`)
})

//helpers

async function allData() {
    const f = await collection.find({}).toArray();
    return f;
}

async function top3() {
    
    const c = await collection.aggregate([{ $group: { _id: "$game", count: { $sum: 1 } } }])
    return c;
}

async function workerReport() {
    // :)
    const c = await collection.aggregate([{ $group: { _id: "$worker", count: { $sum: 1 } } }])
    return c.toArray();
}

async function userReport(id){
    //get games
    const f = await collection.find({"winner":id}).toArray();
    return f;
}

async function getLast10(){
    const f = await collection.find({}).sort({_id:-1}).limit(10);
    return f.toArray();
}

async function best10(){
    const c = await collection.aggregate([{ $group: { _id: "$winner", count: { $sum: -1 } } }]).sort({count:1}).limit(10);
    return c.toArray();
}