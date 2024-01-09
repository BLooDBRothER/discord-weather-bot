FROM golang:1.21.4-bullseye
 
WORKDIR /app
 
COPY . .
 
RUN go mod download
 
RUN go build -o /discord-weather-bot
 
EXPOSE 8080
 
CMD ["/discord-weather-bot"]
