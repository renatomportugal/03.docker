# Host

## Mudar_Nome

```CMD
hostnamectl
hostnamectl set-hostname lab-094
```

## Mudar_IP
```CMD
sudo nano /etc/netplan/00-installer-config.yaml

# This is the network config written by 'subiquity'
network:
  ethernets:
    ens32:
      addresses:
      - 192.168.1.94/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 181.213.132.5
        search:
        - 181.213.132.5
  version: 2
  
```