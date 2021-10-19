
let MadWalletJS = require("../../../MadNetWallet-JS/index.js");

require('dotenv').config({ path: process.cwd()+'/texttx/.env'});
let privateKey;
if(process.env.PRIVATE_KEY) {
    privateKey = process.env.PRIVATE_KEY;
}
else{
    privateKey = "6aea45ee1273170fb525da34015e4f20ba39fe792f486ba74020bcacc9badfc1"
}
let madWallet = new MadWalletJS();
madWallet.Account.addAccount(privateKey, 1);
madWallet.Transaction.createDataStore(madWallet.Account.accounts[0]["address"], 1, 1,"COFFEE")
.catch( error => {
    if (typeof error === 'string'){
        console.log(error)
    }
})
.then(()=>{
    //hash it up
    madWallet.Transaction.Tx._createTx()
})
.catch(()=> {
    console.error("Error sending transaction");
});


