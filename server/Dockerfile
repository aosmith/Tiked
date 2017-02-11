FROM armhf/node
WORKDIR /root/
ADD * /root/
RUN npm install && npm install -g ngrok && ngrok authtoken 4w9pro22iTnSessDjreEm_3nvNXXA2fR7iZdKwMt2Ek
CMD node Server.js
