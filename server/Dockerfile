FROM node
RUN \
  apt-get update && \
  apt-get install -y apt-utils tor && \
  rm -rf /var/lib/apt/lists/*

# Create tor config
RUN mkdir -p /usr/local/etc/tor/hidden_service/ && mkdir -p /etc/tor
RUN echo "HiddenServiceDir /usr/local/etc/tor/hidden_service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:80" >> /etc/tor/torrc
RUN echo "echo 'Starting....\n'" >> /root/start.sh
RUN echo "tor &" >> /root/start.sh
RUN echo "sleep 10 && echo 'All started\n'" >> /root/start.sh
RUN echo "echo 'hostname' && cat /usr/local/etc/tor/hidden_service/hostname && echo '\n\n'" >> /root/start.sh
RUN echo "node ServerTor.js" >> /root/start.sh

WORKDIR /root/

# Define default command.
CMD sh /root/start.sh

# Expose ports.
EXPOSE 4434
EXPOSE 80
EXPOSE 8080
EXPOSE 8000
ADD private_key /usr/local/etc/tor/hidden_service/private_key
ADD * /root/
RUN npm install

