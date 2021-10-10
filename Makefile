build-compose:
	docker-compose -f docker-compose.yml up --build

build-compose-detached:
	docker-compose -f docker-compose.yml up -d --build

compose:
	docker-compose -f docker-compose.yml up
