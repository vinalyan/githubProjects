FROM golang

WORKDIR /bin/bot

COPY . .

EXPOSE 8001

#RUN go build -o ..
