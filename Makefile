docker-up:
	docker-compose up -d

docker-clean:
	docker compose down --remove-orphans
	docker system prune -f

docker-restart:
	make docker-clean 
	make docker-up