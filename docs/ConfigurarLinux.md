# ConfigurarLinux

```CMD
Impedir que o computador entre em standby

sudo nano /etc/systemd/logind.conf
Descomentar a linha, atribuindo o valor ignore.
HandleLidSwitch=ignore

```
