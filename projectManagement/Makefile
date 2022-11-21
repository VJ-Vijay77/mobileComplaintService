

up:
	@echo "starting the containers"
	docker compose up -d
	@echo "containers started!"
down:
	@echo "shutting down containers"
	docker compose down
	@echo "containers stopped!"

rmauth:
	docker stop authentication
	docker rmi -f projectmanagement-authentication:latest
rmcust:
	docker stop customer
	docker rmi -f projectmanagement-customer:latest
	
rmass:
	docker stop associate
	docker rmi -f projectmanagement-associate:latest

rmboth: rmauth rmcust rmass