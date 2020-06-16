# Instalação do Docker
https://docs.docker.com/engine/install/ubuntu/

## Ubuntu
###  Passo 01
1. Atualizar o sistema <br>
```
sudo apt-get update
```
  
### Passo 02
1. Copiar e colar todas essas linhas (-y aceita tudo por padrão) <br>
```
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y
```

### Passo 03
1. Adicionar a chave oficial do Docker<br>
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

## Raspberry Pi 3
https://docs.docker.com/engine/install/debian/#install-using-the-convenience-script
```cat /etc/ *-release```

