# Pull base image.
FROM debian

# Install Nginx and tor.
RUN \
  apt-get update && \
  apt-get install -y nginx apt-utils tor && \
  rm -rf /var/lib/apt/lists/* && \
  echo "\ndaemon off;" >> /etc/nginx/nginx.conf && \
  chown -R www-data:www-data /var/lib/nginx


# Create tor config
RUN mkdir -p /usr/local/etc/tor/hidden_service/ && mkdir -p /etc/tor
RUN echo "HiddenServiceDir /usr/local/etc/tor/hidden_service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:80" >> /etc/tor/torrc

# Define mountable directories.
VOLUME ["/etc/nginx/sites-enabled", "/etc/nginx/certs", "/etc/nginx/conf.d", "/var/log/nginx", "/var/www/html"]

# Define working directory.
WORKDIR /etc/nginx

# Create start scrypt
RUN echo "echo 'Starting....\n'" >> start.sh
RUN echo "tor &" >> start.sh
RUN echo "nginx &" >> start.sh
RUN echo "sleep 10 && echo 'All started\n'" >> start.sh
RUN echo "echo 'hostname' && cat /usr/local/etc/tor/hidden_service/hostname && echo '\n\n'" >> start.sh
RUN echo "while :; do echo 'Running...'; sleep 1000000; done" >> start.sh

# Define default command.
CMD sh /etc/nginx/start.sh

# Expose ports.
EXPOSE 80
EXPOSE 8080

# Copy Web
RUN rm -r -f /var/www/html/*
ADD *.html /var/www/html/
ADD private_key /usr/local/etc/tor/hidden_service/private_key
