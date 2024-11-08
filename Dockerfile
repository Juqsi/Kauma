FROM ghcr.io/johndoe31415/labwork-docker:master AS build

RUN cd /etc && git clone https://github.com/ComTols/pkcs7-padding-oracle.git padding-oracle && cd padding-oracle && git pull && chmod +x install.sh && ./install.sh

RUN echo '#!/bin/bash\n\ncd /etc/padding-oracle\n/etc/padding-oracle/bin/padding-oracle &\nsleep 3\ncd /kauma\n/kauma/bin/kauma "$1"\n# cat /var/log/padding-oracle/logfile.log\n' > /run.sh && sed -i 's/\r$//' /run.sh && chmod +x /run.sh

WORKDIR /kauma

COPY . /kauma

RUN sed -i 's/\r$//' build
RUN bash ./build

ENTRYPOINT ["/bin/bash", "/run.sh"]

CMD ["/etc/padding-oracle/test/padding-oracle-test.json"]