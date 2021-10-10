build-compose:
	docker-compose -f docker-compose.yml up --build --force-recreate

build-compose-detached:
	docker-compose -f docker-compose.yml up -d --build --force-recreate

compose:
	docker-compose -f docker-compose.yml up --force-recreate
