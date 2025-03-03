# ConfigurarLinux

## Atualizar Hora

```CMD
[Atualizar a hora manualmente (PC sem a bateria da BIOS)]
date
sudo date -s MM/DD/AAAA
sudo date -s HH:MM:SS

[Atualizar a hora automaticamente]
sudo date -s "$(wget -qSO- --max-redirect=0 google.com 2>&1 | grep Date: | cut -d' ' -f5-8)Z"
sudo dpkg-reconfigure tzdata
Escolher America, São Paulo
timedatectl
Retorno:
Local time: Wed 2024-01-17 14:02:56 -03
Universal time: Wed 2024-01-17 17:02:56 UTC
RTC time: Sat 2005-01-01 01:00:38
Time zone: America/Sao_Paulo (-03, -0300)
System clock synchronized: no
NTP service: active
RTC in local TZ: no

[Atualizar a hora automaticamente com crontab]
sudo su
crontab -l
no crontab for root
crontab -e
Select an editor.  To change later, run 'select-editor'.
  1. /bin/nano        <---- easiest
  2. /usr/bin/vim.tiny

Choose 1-2 [1]: 1

@reboot date -s "$(wget -qSO- --max-redirect=0 google.com 2>&1 | grep Date: | cut -d' ' -f5-8)Z"

Ctrl+O, Enter, Ctrl+X

crontab: installing new crontab


crontab -l
# Edit this file to introduce tasks to be run by cron.
#
# Each task to run has to be defined through a single line
# indicating with different fields when the task will be run
# and what command to run for the task
#
# To define the time you can provide concrete values for
# minute (m), hour (h), day of month (dom), month (mon),
# and day of week (dow) or use '*' in these fields (for 'any').
#
# Notice that tasks will be started based on the cron's system
# daemon's notion of time and timezones.
#
# Output of the crontab jobs (including errors) is sent through
# email to the user the crontab file belongs to (unless redirected).
#
# For example, you can run a backup of all your user accounts
# at 5 a.m every week with:
# 0 5 * * 1 tar -zcf /var/backups/home.tgz /home/
#
# For more information see the manual pages of crontab(5) and cron(8)
#
# m h  dom mon dow   command
@reboot date -s "$(wget -qSO- --max-redirect=0 google.com 2>&1 | grep Date: | cut -d' ' -f5-8)Z"

-------------------------------------------------------------------------------
```

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
