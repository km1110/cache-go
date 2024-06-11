init:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

go:
	docker exec -it golang /bin/bash

redis:
	docker compose exec -it redis bash

ps:
	docker ps