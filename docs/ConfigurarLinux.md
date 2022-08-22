# ConfigurarLinux

## standby

```CMD
01. Impedir que o computador entre em standby quando fechar a tampa.
Se você usa um notebook como servidor linux e quer deixar ele ligado quando fecha a tampa, basta editar o item abaixo:

sudo nano /etc/systemd/logind.conf
Descomentar a linha, atribuindo o valor ignore.
HandleLidSwitch=ignore
Ctrl+O, Enter, Crtl+X
sudo reboot

Pronto, após apertar o botão de ligar pode baixar a tampa imediatamente.

```

## Mudar_Host

```CMD
sudo nano /etc/hosts
Mudar a linha:
127.0.1.1 novo_host
Ctrl+O, Enter, Ctrl+X

sudo nano /etc/hostname
novo_host
Ctrl+O, Enter, Ctrl+X

sudo nano /proc/sys/kernel/hostname
novo_host
Ctrl+O, Enter, Ctrl+X
```

## Rede

```CMD
sudo nmtui

```
