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

fix-mysql-rights:
	sudo chmod 777 -R  database/master && sudo chmod 0444  database/master/conf.d/master.cnf

build: build-frontend build-backend backend-service-restart
