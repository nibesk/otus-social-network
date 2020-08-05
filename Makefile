env:
	cp example.env .env

build-frontend:
	npm install --prefix ./frontend && \
 	npm run build --prefix ./frontend

build-backend:
	go install && go build -o ./build

start-backend-service:
	sudo service gosocialotus start

build: build-frontend build-backend
