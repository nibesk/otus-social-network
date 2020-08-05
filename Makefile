env:
	cp .env.example .env

build-frontend:
	cd frontend && npm run build

build-backend:
	go build -o ./build
