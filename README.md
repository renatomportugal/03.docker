# Instalação do Docker
https://docs.docker.com/engine/install/ubuntu/

## Ubuntu
###  Passo 01
Atualizar o sistema<br>
```
sudo apt-get update
```
  
### Passo 02
Copiar e colar todas essas linhas (-y aceita tudo por padrão)<br>
```
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y
```

### Passo 03
Adicionar a chave oficial do Docker<br>
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

### Passo 04
Verificar a impressão digital<br>
```
sudo apt-key fingerprint 0EBFCD88
```
Tela de retorno:<br>
```
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

### Passo 05
Setup do repositório<br>
```
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

### Passo 06
Instalar o motor Docker<br>
```
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

### Passo 07
Verificar se instalou<br>
```
docker --help
```

### Passo 08
Rodar a imagem hello-world<br>
```
sudo docker run hello-world
```

### Passo 09
Ver quais containers existem<br>
```
docker ps
```
Dará um erro de permissão, descobriremos qual o nome do usuário utilizado.<br>
```
whoami
```
Colocaremos o usuário no grupo Docker.<br>
```
sudo usermod -aG docker usuario
```

### Passo 10
Reiniciar o sistema.<br>
```
reboot
```

## Raspberry Pi 3
https://docs.docker.com/engine/install/debian/#install-using-the-convenience-script<br>
```
cat /etc/ *-release
```
