import { Router } from "express";
import {MongoClient} from 'mongodb'
const router=Router()
const user=encodeURIComponent("root")
const pass=encodeURIComponent("1234")
const uri= `mongodb://${user}:${pass}@35.223.93.218:27017`
const mongo=new MongoClient(uri)

router.get("/getLogs",async(req,res)=>{
    let col=mongo.db("mydb").collection("artista")
    let a=await col.find({}).sort({_id:-1}).limit(20).project({_id:0,date:1,name:1,album:1}).toArray()
    console.log(a)
    res.send(a)
})

export default router