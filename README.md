# Instalação do Docker
https://docs.docker.com/engine/install/ubuntu/

## Ubuntu
###  Passo 01
1. Atualizar o sistema
<br>
```
  sudo apt-get update
```
  
### Passo 02
1. Copiar e colar todas essas linhas (-y aceita tudo por padrão)
<br>
```
  sudo apt-get install \
      apt-transport-https \
      ca-certificates \
      curl \
      gnupg-agent \
      software-properties-common -y
```

### Passo 03
1. Adicionar a chave oficial do Docker
<br>
`  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -`

### Passo 04
1. Verificar a impressão digital
<br>
```
sudo apt-key fingerprint 0EBFCD88
```
Tela de retorno:
<br>
```
  pub   rsa4096 2017-02-22 [SCEA]
        9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
  uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
  sub   rsa4096 2017-02-22 [S]
```


### Passo xx
1. 
<br>
```
xx
```


## Raspberry Pi 3
https://docs.docker.com/engine/install/debian/#install-using-the-convenience-script
```cat /etc/ *-release```

