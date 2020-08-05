env:
	cp example.env .env

build-frontend:
	cd frontend && npm install && npm run build

build-backend:
	go install && go build -o ./build

start-backend-service:
	sudo service gosocialotus start
