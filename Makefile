build:
	docker build -t vinalyan/bot-huelpet:latest .
build_vds:
	docker build --platform linux/amd64 -t vinalyan/bot-huelpet:amd64 .
run:
	docker run --rm -it --name botHuelpet vinalyan/bot-huelpet:latest
stop:
	docker stop botHuelpet
push:
	docker push vinalyan/bot-huelpet:latest

pull:
	docker pull vinalyan/bot-huelpet:latest



