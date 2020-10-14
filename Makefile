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

db-init:
	docker exec -i -t otn__chat_mongodb sh -c "mongo -u root -p mongo < /init/init"
	docker exec -i -t osn__users_mysql-master sh -c "mysql -uroot -pmysql dbase < /init_sql/initdb.sql"
	docker exec -i -t osn__users_mysql-master sh -c "mysql -uroot -pmysql dbase < /init_sql/mock_users.sql"

