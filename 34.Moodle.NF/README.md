
# moodle

```CMD
https://hub.docker.com/r/bitnami/moodle/


User and Site configuration
MOODLE_USERNAME: Moodle application username. Default: user
MOODLE_PASSWORD: Moodle application password. Default: bitnami
MOODLE_EMAIL: Moodle application email. Default: user@example.com
MOODLE_SITE_NAME: Moodle site name. Default: New Site
MOODLE_SKIP_BOOTSTRAP: Do not initialize the Moodle database for a new deployment. This is necessary in case you use a database that already has Moodle data. Default: no
Use an existing database
MOODLE_DATABASE_TYPE: Database type. Valid values: mariadb, mysqli. Default: mariadb
MOODLE_DATABASE_HOST: Hostname for database server. Default: mariadb
MOODLE_DATABASE_PORT_NUMBER: Port used by database server. Default: 3306
MOODLE_DATABASE_NAME: Database name that Moodle will use to connect with the database. Default: bitnami_moodle
MOODLE_DATABASE_USER: Database user that Moodle will use to connect with the database. Default: bn_moodle
MOODLE_DATABASE_PASSWORD: Database password that Moodle will use to connect with the database. No defaults.
ALLOW_EMPTY_PASSWORD: It can be used to allow blank passwords. Default: no
Create a database for Moodle using mysql-client
MYSQL_CLIENT_FLAVOR: SQL database flavor. Valid values: mariadb or mysql. Default: mariadb.
MYSQL_CLIENT_DATABASE_HOST: Hostname for MariaDB server. Default: mariadb
MYSQL_CLIENT_DATABASE_PORT_NUMBER: Port used by MariaDB server. Default: 3306
MYSQL_CLIENT_DATABASE_ROOT_USER: Database admin user. Default: root
MYSQL_CLIENT_DATABASE_ROOT_PASSWORD: Database password for the database admin user. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_NAME: New database to be created by the mysql client module. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_USER: New database user to be created by the mysql client module. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_PASSWORD: Database password for the MYSQL_CLIENT_CREATE_DATABASE_USER user. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_CHARACTER_SET: Character set to use for the new database. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_COLLATE: Database collation to use for the new database. No defaults.
MYSQL_CLIENT_CREATE_DATABASE_PRIVILEGES: Database privileges to grant for the user specified in MYSQL_CLIENT_CREATE_DATABASE_USER to the database specified in MYSQL_CLIENT_CREATE_DATABASE_NAME. No defaults.
MYSQL_CLIENT_ENABLE_SSL_WRAPPER: Whether to force SSL connections to the database via the mysql CLI tool. Useful for applications that rely on the CLI instead of APIs. Default: no
MYSQL_CLIENT_ENABLE_SSL: Whether to force SSL connections for the database. Default: no
MYSQL_CLIENT_SSL_CA_FILE: Path to the SSL CA file for the new database. No defaults
MYSQL_CLIENT_SSL_CERT_FILE: Path to the SSL CA file for the new database. No defaults
MYSQL_CLIENT_SSL_KEY_FILE: Path to the SSL CA file for the new database. No defaults
ALLOW_EMPTY_PASSWORD: It can be used to allow blank passwords. Default: no
SMTP Configuration
To configure MoodleTM to send email using SMTP you can set the following environment variables:

MOODLE_SMTP_HOST: SMTP host.
MOODLE_SMTP_PORT: SMTP port.
MOODLE_SMTP_USER: SMTP account user.
MOODLE_SMTP_PASSWORD: SMTP account password.
MOODLE_SMTP_PROTOCOL: SMTP protocol.
PHP configuration
PHP_ENABLE_OPCACHE: Enable OPcache for PHP scripts. No default.
PHP_EXPOSE_PHP: Enables HTTP header with PHP version. No default.
PHP_MAX_EXECUTION_TIME: Maximum execution time for PHP scripts. No default.
PHP_MAX_INPUT_TIME: Maximum input time for PHP scripts. No default.
PHP_MAX_INPUT_VARS: Maximum amount of input variables for PHP scripts. No default.
PHP_MEMORY_LIMIT: Memory limit for PHP scripts. Default: 256M
PHP_POST_MAX_SIZE: Maximum size for PHP POST requests. No default.
PHP_UPLOAD_MAX_FILESIZE: Maximum file size for PHP uploads. No default.


moodle:
    ...
    environment:
      - MOODLE_DATABASE_USER=bn_moodle
      - MOODLE_DATABASE_NAME=bitnami_moodle
      - ALLOW_EMPTY_PASSWORD=yes
      - MOODLE_SMTP_HOST=smtp.gmail.com
      - MOODLE_SMTP_PORT=587
      - MOODLE_SMTP_USER=your_email@gmail.com
      - MOODLE_SMTP_PASSWORD=your_password
      - MOODLE_SMTP_PROTOCOL=tls
  ...

  docker run -d --name moodle -p 80:8080 -p 443:8443 \
 --env MOODLE_DATABASE_USER=bn_moodle \
 --env MOODLE_DATABASE_NAME=bitnami_moodle \
 --env MOODLE_SMTP_HOST=smtp.gmail.com \
 --env MOODLE_SMTP_PORT=587 \
 --env MOODLE_SMTP_USER=your_email@gmail.com \
 --env MOODLE_SMTP_PASSWORD=your_password \
 --env MOODLE_SMTP_PROTOCOL=tls \
 --network moodle-tier \
 --volume /path/to/moodle-persistence:/bitnami \
 bitnami/moodle:latest

 moodle:
...
  # image: 'bitnami/moodle:3' # remove this line !
  build:
    context: .
    dockerfile: Dockerfile
    args:
      - EXTRA_LOCALES=fr_FR.UTF-8 UTF-8, de_DE.UTF-8 UTF-8, it_IT.UTF-8 UTF-8, es_ES.UTF-8 UTF-8
...

docker build -t bitnami/moodle:latest --build-arg EXTRA_LOCALES="fr_FR.UTF-8 UTF-8, de_DE.UTF-8 UTF-8, it_IT.UTF-8 UTF-8, es_ES.UTF-8 UTF-8" .

moodle:
...
  # image: 'bitnami/moodle:3' # remove this line !
  build:
    context: .
    dockerfile: Dockerfile
    args:
      - WITH_ALL_LOCALES=yes
...

docker build -t bitnami/moodle:latest --build-arg WITH_ALL_LOCALES=yes .

FROM bitnami/moodle
RUN echo "es_ES.UTF-8 UTF-8" >> /etc/locale.gen && locale-gen

docker logs moodle

docker-compose logs moodle

Maintenance
Backing up your container
To backup your data, configuration and logs, follow these simple steps:

Step 1: Stop the currently running container
$ docker stop moodle
Or using Docker Compose:

$ docker-compose stop moodle
Step 2: Run the backup command
We need to mount two volumes in a container we will use to create the backup: a directory on your host to store the backup in, and the volumes from the container we just stopped so we can access the data.

$ docker run --rm -v /path/to/moodle-backups:/backups --volumes-from moodle busybox \
  cp -a /bitnami/moodle /backups/latest

Restoring a backup
Restoring a backup is as simple as mounting the backup as volumes in the containers.

For the MariaDB database container:

 $ docker run -d --name mariadb \
   ...
-  --volume /path/to/mariadb-persistence:/bitnami/mariadb \
+  --volume /path/to/mariadb-backups/latest:/bitnami/mariadb \
   bitnami/mariadb:latest
For the MoodleTM container:

 $ docker run -d --name moodle \
   ...
-  --volume /path/to/moodle-persistence:/bitnami/moodle \
+  --volume /path/to/moodle-backups/latest:/bitnami/moodle \
   bitnami/moodle:latest

   UPDATE

   docker pull bitnami/moodle:latest

   docker-compose stop moodle

   Depois de fazer o backup

   docker-compose rm -v moodle

   docker-compose up -d

   ettings that can be adapted using environment variables. For instance, you can change the ports used by Apache for HTTP and HTTPS, by setting the environment variables APACHE_HTTP_PORT_NUMBER and APACHE_HTTPS_PORT_NUMBER respectively.

FROM bitnami/moodle
LABEL maintainer "Bitnami <containers@bitnami.com>"

## Change user to perform privileged actions
USER 0
## Install 'vim'
RUN install_packages vim
## Revert to the original non-root user
USER 1001

## Enable mod_ratelimit module
RUN sed -i -r 's/#LoadModule ratelimit_module/LoadModule ratelimit_module/' /opt/bitnami/apache/conf/httpd.conf

## Modify the ports used by Apache by default
# It is also possible to change these environment variables at runtime
ENV APACHE_HTTP_PORT_NUMBER=8181
ENV APACHE_HTTPS_PORT_NUMBER=8143
EXPOSE 8181 8143

  moodle:
-    image: bitnami/moodle:latest
+    build: .
     ports:
-      - '80:8080'
-      - '443:8443'
+      - '80:8181'
+      - '443:8143'
     environment:
       ...
+      - PHP_MEMORY_LIMIT=512m
     ...



```

## TAREFAS

```CMD
[ ] Descobrir como instalar em português, com codificação para português.
```
