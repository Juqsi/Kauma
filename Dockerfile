FROM ghcr.io/johndoe31415/labwork-docker:master
WORKDIR /app
COPY . .
RUN ./build
CMD ["./kauma","main/Aufgaben/input/poly2block"]