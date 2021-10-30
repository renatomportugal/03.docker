# DNS

```CMD
sudo apt -y install dnsmasq

sudo nano /etc/dnsmasq.conf

# line 19: uncomment (never forward plain names)
domain-needed
# line 21: uncomment (never forward addresses in the non-routed address spaces)
bogus-priv
# line 53: uncomment (query with each server strictly in the order in resolv.conf)
strict-order
# line 67: add if you need
# query the specific domain name to the specific DNS server
# the example follows means query [server.education] domain to the [192.168.1.106] server
server=/server.education/192.168.1.106
# line 135: uncomment (add domain name automatically)
expand-hosts
# line 145: add (define domain name)
domain=srv.world

Ctrl+O, Enter, Ctrl+X



sudo ln -fs /run/systemd/resolve/resolv.conf /etc/resolv.conf

sudo systemctl restart dnsmasq systemd-resolved

sudo nano /etc/hosts

# add records
192.168.1.106       dlp.srv.world dlp

Ctrl+O, Enter, Ctrl+X

sudo systemctl restart dnsmasq

cd /etc/netplan/
ls -la


sudo nano /etc/netplan/00-installer-config.yaml

adicione em nameservers:
addresses:
- 192.168.1.106

search:
- 192.168.1.106

Ctrl+O, Enter, Ctrl+X
  
sudo netplan apply

sudo systemd-resolve --status --no-pager | tail -9

Link 2 (enp1s0)
      Current Scopes: DNS
DefaultRoute setting: yes
       LLMNR setting: yes
MulticastDNS setting: no
  DNSOverTLS setting: no
      DNSSEC setting: no
    DNSSEC supported: no
         DNS Servers: 192.168.1.106
		 
dig dlp.srv.world

; <<>> DiG 9.16.1-Ubuntu <<>> dlp.srv.world.
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 11995
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;dlp.srv.world.                 IN      A

;; ANSWER SECTION:
dlp.srv.world.          0       IN      A       192.168.1.106

;; Query time: 0 msec
;; SERVER: 192.168.1.106#53(192.168.1.106)
;; WHEN: Wed Aug 25 20:44:02 JST 2020
;; MSG SIZE  rcvd: 58

dig -x 192.168.1.106

; <<>> DiG 9.16.1-Ubuntu <<>> -x 192.168.1.106
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 17844
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;30.0.0.10.in-addr.arpa.                IN      PTR

;; ANSWER SECTION:
30.0.0.10.in-addr.arpa. 0       IN      PTR     dlp.srv.world.

;; Query time: 0 msec
;; SERVER: 192.168.1.106#53(192.168.1.106)
;; WHEN: Wed Aug 25 20:45:32 JST 2020
;; MSG SIZE  rcvd: 78

```
