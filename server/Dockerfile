FROM armhf/node
WORKDIR /root/
CMD node /root/Server.js
RUN npm install -g ngrok && ngrok authtoken 6pEALdcXicqP5THVYw5fE_6wegT2YcNHMuhuHTtWuVH
EXPOSE 4434
EXPOSE 8000
ADD * /root/
RUN npm install
# run with -it
# tikedzh6cg5unkrf.onion
