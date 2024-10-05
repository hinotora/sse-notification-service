DEMOPATH = ${APP_NAME}

build:
	@docker rmi sse-notification-service-app
	@docker compose build

up:
	@docker compose up -d

logs:
	@docker compose logs --follow

down: 
	@docker compose down

restart: 
	@docker compose restart
