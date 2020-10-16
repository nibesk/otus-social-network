# To fresh run
```
> make fresh_run 
> make build-db
```
And wait till they get ready to connections
Then in separate terminal
```
> make db_init
```
Stop database containers and then start other containers 
```
docker-compose down
docker-compose up -- build -d
```

Open url in browser http://localhost:8080/

You are all set:)
