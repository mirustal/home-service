ps:
	docker-compose ps

up-d:
	docker-compose down
	docker-compose up -d

logs:
	docker-compose logs -f

up-b:
	docker-compose up --build

down:
	docker-compose down