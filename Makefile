env:
	cp backend\example.env backend\.env

start:
	docker-compose up -d osn-users_mysql-master osn-users_backend osn_npm

preview:
	docker-compose up

fix-mysql-rights:
	sudo chmod 777 -R  database/master && sudo chmod 0444  database/master/conf.d/master.cnf

