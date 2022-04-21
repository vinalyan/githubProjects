FROM golang

WORKDIR /bin/read-adviser-bot

COPY . .

EXPOSE 8001
