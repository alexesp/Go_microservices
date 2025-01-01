FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp

##up: start todos los contenedores
up:
	@echo "Starting Docker imagesl..."
	docker-compose up -docker
	@echo "Docker images started!"

##up_build: stop docker-compose (si esta ejecutandose), compilar todo y start docker compose
up_build: build_broker build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when requered) and startin docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: construir broker binario como linux ejecutable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: construir auth binario como linux ejecutable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"
## build_front: construir front en binario
build_front:
	@echo "Building front and binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: start fron end
start:
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop front end
stop:
	@echo "Stopping fron end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped fron end!"