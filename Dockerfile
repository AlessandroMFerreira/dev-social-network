FROM golang:1.21-alpine3.18
WORKDIR /app
COPY . .
RUN go mod download
RUN go build
CMD [ "./dev-social-network" ]
EXPOSE 5000