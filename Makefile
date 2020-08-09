env:
	cp example.env .env

build-frontend:
	npm install --prefix ./frontend && \
 	npm run build --prefix ./frontend

build-backend:
	go install && \
 	go build -o ./build

backend-service-start:
	sudo service gosocialotus start

backend-service-restart:
	sudo service gosocialotus restart

build: build-frontend build-backend backend-service-restart
