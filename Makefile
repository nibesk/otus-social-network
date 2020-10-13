env:
	cp backend\example.env backend\.env

start:
	docker-compose up -d osn-users_mysql-master osn-users_backend osn_npm

preview:
	docker-compose up

fix-rights:
	sudo chmod 777 -R service-users/database/master
	sudo chmod 0444  service-users/database/master/conf.d/master.cnf
	sudo chmod 777 -R service-chat/database/data
	sudo chmod 777 -R service-chat/database/logs

mongo-clear:
	sudo rm -rf service-chat/database/data/*
	touch service-chat/database/data/.gitkeep
	sudo rm -rf service-chat/database/logs/*
	touch service-chat/database/logs/.gitkeep

force_recreate:
	docker-compose up --build --force-recreate $(n)
