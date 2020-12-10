FROM faddat/sos-full

RUN cd ~/ && \
      curl -s https://api.github.com/repos/ovrclk/akash/releases/latest \
      | grep "linux_amd64.zip" \
      | cut -d : -f 2,3 \
      | tr -d \" \
      | curl -o akash.zip -  \
      
      
