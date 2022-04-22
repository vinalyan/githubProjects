FROM golang

WORKDIR /home/bot

COPY . .

EXPOSE 8001

RUN go build -o /bin

