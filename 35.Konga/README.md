# Konga

```CMD
https://hub.docker.com/r/pantsel/konga/


docker pull pantsel/konga

 docker run -p 1337:1337 \
             --network {{kong-network}} \ // optional
             --name konga \
             -e "NODE_ENV=production" \ // or "development" | defaults to 'development'
             -e "TOKEN_SECRET={{somerandomstring}}" \
             pantsel/konga


docker run -p 1337:1337 
          --network {{kong-network}} \ // optional
          -e "TOKEN_SECRET={{somerandomstring}}" \
          -e "DB_ADAPTER=the-name-of-the-adapter" \ // 'mongo','postgres','sqlserver'  or 'mysql'
          -e "DB_HOST=your-db-hostname" \
          -e "DB_PORT=your-db-port" \ // Defaults to the default db port
          -e "DB_USER=your-db-user" \ // Omit if not relevant
          -e "DB_PASSWORD=your-db-password" \ // Omit if not relevant
          -e "DB_DATABASE=your-db-name" \ // Defaults to 'konga_database'
          -e "DB_PG_SCHEMA=my-schema"\ // Optionally define a schema when integrating with prostgres
          -e "NODE_ENV=production" \ // or 'development' | defaults to 'development'
          --name konga \
          pantsel/konga
		  
		  

```
