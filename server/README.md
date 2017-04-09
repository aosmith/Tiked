# Server
#### Server.js is the C&C server written in node js, its sends the commands to the users and is the only interface to the users

# Use
#### To use it you will neeed:
- nodejs
- ngrok setup in your computer, cretate a account in their web page
- A pastebin URL and API keys

# How it works
- Ngrok gives you a url and port which you can use to communicate and thus avoiding to foward ports
- With the Pastebin API the Ngrok url is posted on a pastebin note (This is done so that the client doesn't connect to a fixed ip and instead fetches it at runtime, was the best aproach for my needs. Additionally it should encrypt the url when posted on pastebin and decrypted by the client for more security)
- Server listens for connections and gives a promt

### TODO
- [ ] Use encryption 

### In progress
- Due to the problems of using pastebin I have inplemented a docker-contained onion server to dispatch url and files because of the advantages of not having to forward ports and it's anonimity.
