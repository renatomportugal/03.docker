```
a base de dados redmine deve ser criada.

cd docker
docker-compose up
 


senha
admin
admin

Crie conta para o admin aprovar.
Funcionou.

http://192.168.1.109/



/var/www/html/data

_______________________________________________________________________________
__CONFIGURAÇÕES________________________________________________________________
Autorizar domínio

docker ps -a
docker exec -it id-do-container bash
apt-get update
apt-get install nano

nano /var/www/html/config/config.php

'trusted_domains' =>
array(
    0 => 'localhost',
    1 => 'seuSite.com'
)

_______________________________________________________________________________
__PROBLEMAS____________________________________________________________________
localhost:8080/index.php?route=/server/variables

Edite para 512000
sort buffer size
https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_sort_buffer_size


```