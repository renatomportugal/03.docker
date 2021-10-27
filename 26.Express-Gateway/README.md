# Express-Gateway

```CMD```
docker run -d --name express-gateway-data-store \
                -p 6379:6379 \
                redis:alpine
				
docker run -d --name express-gateway \
    --link eg-database:eg-database \
    -v /my/own/datadir:/var/lib/eg \
    -p 8080:8080 \
    -p 9876:9876 \
    express-gateway
```
