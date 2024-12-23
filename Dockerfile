FROM ghcr.io/johndoe31415/labwork-docker:master AS build

WORKDIR /kauma

COPY . /kauma

RUN sed -i 's/\r$//' build
RUN bash ./build

ENTRYPOINT ["/bin/bash", "/run.sh"]
