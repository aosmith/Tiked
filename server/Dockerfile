FROM debian
WORKDIR /root/
ADD * /root/
RUN \
  apt-get update && \
  apt-get install -y node nginx apt-utils tor && \
  rm -rf /var/lib/apt/lists/* && \
  echo "\ndaemon off;" >> /etc/nginx/nginx.conf && \
  chown -R www-data:www-data /var/lib/nginx

# Create tor config
RUN mkdir -p /usr/local/etc/tor/hidden_service/ && mkdir -p /etc/tor
RUN echo "HiddenServiceDir /usr/local/etc/tor/hidden_service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:80" >> /etc/tor/torrc

RUN echo "echo 'Starting....\n'" >> start.sh
RUN echo "tor &" >> start.sh
RUN echo "nginx &" >> start.sh
RUN echo "sleep 10 && echo 'All started\n'" >> start.sh
RUN echo "echo 'hostname' && cat /usr/local/etc/tor/hidden_service/hostname && echo '\n\n'" >> start.sh
RUN echo "node ServerTor.js" >> start.sh

# Define default command.
CMD sh /root/start.sh

# Expose ports.
EXPOSE 4434
EXPOSE 80
EXPOSE 8080
EXPOSE 8000
RUN npm install

