build:
	docker build  --platform linux/amd64 -t vinalyan/bot-huelpet:latest .
run:
	docker run --rm --name botHuelpet vinalyan/bot-huelpet:i386
run-dev:
	docker run -it --rm --name botHuelpet vinalyan/bot-huelpet
stop:
	docker stop botHuelpet
push:
	docker push vinalyan/bot-huelpet:latest

pull:
	docker pull vinalyan/bot-huelpet:latest



