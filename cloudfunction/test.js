const axios = require('axios');

async function getWords(){
    try{
        console.log("_____________________")
        response = await axios.get('https://gohellow-duuq6naepa-uc.a.run.app/firestore/collection/words');
        console.log("data", response.data)

        // data = [{a:"3"}, {a:"1"}]
        meaningArr = response.data.map( d => d.meaning);

        palabras = meaningArr.join(", ")
        console.log(meaningArr, palabras)
        return palabras;
    }catch(err){
        console.log(err)
    }
}

getWords()