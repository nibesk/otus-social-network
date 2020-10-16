# To fresh run
```
> make fresh_run 
> make build-db
```
And wait till databases get ready for connections
Then in separate terminal
```
> make db_init
```
Stop databases containers and then start other containers 
```
docker-compose down
docker-compose up --build -d
```

Open url in browser http://localhost:8080/

You are all set:)
