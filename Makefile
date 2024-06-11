init:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

go:
	docker exec -it golang /bin/bash

redis:
	docker exec -it redis /bin/bash

ps:
	docker ps