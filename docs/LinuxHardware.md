# LinuxHardware

```CMD
listar o hardware
sudo lshw -short
Retorno:
H/W path             Device     Class          Description
==========================================================
                                system         AOD270 (Type1Sku0)
/0                              bus            JE01_CT
/0/0                            memory         128KiB BIOS
/0/20                           memory         2GiB System Memory
/0/20/0                         memory         [empty]
/0/20/1                         memory         2GiB SODIMM DDR3 Synchronous 666 MHz (1,5 ns)
/0/28                           processor      Intel(R) Atom(TM) CPU N2600   @ 1.60GHz
/0/29                           memory         512KiB L2 cache
/0/2a                           memory         32KiB L1 cache
/0/100                          bridge         Atom Processor D2xxx/N2xxx DRAM Controller
/0/100/2             /dev/fb0   display        Atom Processor D2xxx/N2xxx Integrated Graphics Controller
/0/100/1b            card0      multimedia     NM10/ICH7 Family High Definition Audio Controller
/0/100/1b/0          input15    input          HDA Intel HDMI/DP,pcm=3
/0/100/1b/1          input16    input          HDA Intel Mic
/0/100/1b/2          input17    input          HDA Intel Headphone
/0/100/1c                       bridge         NM10/ICH7 Family PCI Express Port 1
/0/100/1c/0          enp1s0     network        RTL810xE PCI Express Fast Ethernet controller
/0/100/1c.1                     bridge         NM10/ICH7 Family PCI Express Port 2
/0/100/1c.1/0        wlp2s0     network        AR9485 Wireless Network Adapter
/0/100/1c.2                     bridge         NM10/ICH7 Family PCI Express Port 3
/0/100/1c.2/0        mmc0       bus            RTS5209 PCI Express Card Reader
/0/100/1d                       bus            NM10/ICH7 Family USB UHCI Controller #1
/0/100/1d/1          usb2       bus            UHCI Host Controller
/0/100/1d.1                     bus            NM10/ICH7 Family USB UHCI Controller #2
/0/100/1d.1/1        usb3       bus            UHCI Host Controller
/0/100/1d.2                     bus            NM10/ICH7 Family USB UHCI Controller #3
/0/100/1d.2/1        usb4       bus            UHCI Host Controller
/0/100/1d.3                     bus            NM10/ICH7 Family USB UHCI Controller #4
/0/100/1d.3/1        usb5       bus            UHCI Host Controller
/0/100/1d.3/1/2                 communication  Bluetooth USB Host Controller
/0/100/1d.7                     bus            NM10/ICH7 Family USB2 EHCI Controller
/0/100/1d.7/1        usb1       bus            EHCI Host Controller
/0/100/1d.7/1/3      input18    multimedia     WebCam: WebCam
/0/100/1e                       bridge         82801 Mobile PCI Bridge
/0/100/1f                       bridge         NM10 Family LPC Controller
/0/100/1f/0                     input          PnP device PNP0303
/0/100/1f/1                     generic        PnP device SYN1b20
/0/100/1f/2                     system         PnP device PNP0103
/0/100/1f/3                     system         PnP device PNP0c02
/0/100/1f/4                     system         PnP device PNP0c02
/0/100/1f/5                     system         PnP device PNP0b00
/0/100/1f/6                     system         PnP device PNP0c02
/0/100/1f.2          scsi0      storage        NM10/ICH7 Family SATA Controller [AHCI mode]
/0/100/1f.2/0.0.0    /dev/sda   disk           120GB SanDisk SSD PLUS
/0/100/1f.2/0.0.0/1  /dev/sda1  volume         1023KiB BIOS Boot partition
/0/100/1f.2/0.0.0/2  /dev/sda2  volume         2GiB EXT4 volume
/0/100/1f.2/0.0.0/3  /dev/sda3  volume         109GiB EFI partition
/0/100/1f.3                     bus            NM10/ICH7 Family SMBus Controller
/1                   input0     input          Power Button
/2                   input1     input          Sleep Button
/3                   input10    input          SynPS/2 Synaptics TouchPad
/4                   input13    input          Video Bus
/5                   input14    input          Acer WMI hotkeys
/6                   input2     input          Lid Switch
/7                   input3     input          Power Button
/8                   input4     input          AT Translated Set 2 keyboard
```

## CPU

```CMD
informações sobre o processador.
lscpu
Retorno:
Architecture:           x86_64
  CPU op-mode(s):       32-bit, 64-bit
  Address sizes:        36 bits physical, 48 bits virtual
  Byte Order:           Little Endian
CPU(s):                 4
  On-line CPU(s) list:  0-3
Vendor ID:              GenuineIntel
  Model name:           Intel(R) Atom(TM) CPU N2600   @ 1.60GHz
    CPU family:         6
    Model:              54
    Thread(s) per core: 2
    Core(s) per socket: 2
    Socket(s):          1
    Stepping:           1
    CPU max MHz:        1600,0000
    CPU min MHz:        600,0000
    BogoMIPS:           3191.92
    Flags:              fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx lm
                         constant_tsc arch_perfmon pebs bts rep_good nopl nonstop_tsc cpuid aperfmperf pni dtes64 monitor ds_cpl est tm2 ssse3 cx16 xtpr pdcm
                         movbe lahf_lm dtherm arat
Caches (sum of all):
  L1d:                  48 KiB (2 instances)
  L1i:                  64 KiB (2 instances)
  L2:                   1 MiB (2 instances)
NUMA:
  NUMA node(s):         1
  NUMA node0 CPU(s):    0-3
Vulnerabilities:
  Gather data sampling: Not affected
  Itlb multihit:        Not affected
  L1tf:                 Not affected
  Mds:                  Not affected
  Meltdown:             Not affected
  Mmio stale data:      Not affected
  Retbleed:             Not affected
  Spec store bypass:    Not affected
  Spectre v1:           Not affected
  Spectre v2:           Not affected
  Srbds:                Not affected
  Tsx async abort:      Not affected
```

## PCI

```CMD
Dispositivos PCI podem ser listados com
lspci
Retorno:
00:00.0 Host bridge: Intel Corporation Atom Processor D2xxx/N2xxx DRAM Controller (rev 04)
00:02.0 VGA compatible controller: Intel Corporation Atom Processor D2xxx/N2xxx Integrated Graphics Controller (rev 0b)
00:1b.0 Audio device: Intel Corporation NM10/ICH7 Family High Definition Audio Controller (rev 02)
00:1c.0 PCI bridge: Intel Corporation NM10/ICH7 Family PCI Express Port 1 (rev 02)
00:1c.1 PCI bridge: Intel Corporation NM10/ICH7 Family PCI Express Port 2 (rev 02)
00:1c.2 PCI bridge: Intel Corporation NM10/ICH7 Family PCI Express Port 3 (rev 02)
00:1d.0 USB controller: Intel Corporation NM10/ICH7 Family USB UHCI Controller #1 (rev 02)
00:1d.1 USB controller: Intel Corporation NM10/ICH7 Family USB UHCI Controller #2 (rev 02)
00:1d.2 USB controller: Intel Corporation NM10/ICH7 Family USB UHCI Controller #3 (rev 02)
00:1d.3 USB controller: Intel Corporation NM10/ICH7 Family USB UHCI Controller #4 (rev 02)
00:1d.7 USB controller: Intel Corporation NM10/ICH7 Family USB2 EHCI Controller (rev 02)
00:1e.0 PCI bridge: Intel Corporation 82801 Mobile PCI Bridge (rev e2)
00:1f.0 ISA bridge: Intel Corporation NM10 Family LPC Controller (rev 02)
00:1f.2 SATA controller: Intel Corporation NM10/ICH7 Family SATA Controller [AHCI mode] (rev 02)
00:1f.3 SMBus: Intel Corporation NM10/ICH7 Family SMBus Controller (rev 02)
01:00.0 Ethernet controller: Realtek Semiconductor Co., Ltd. RTL810xE PCI Express Fast Ethernet controller (rev 05)
02:00.0 Network controller: Qualcomm Atheros AR9485 Wireless Network Adapter (rev 01)
03:00.0 Unassigned class [ff00]: Realtek Semiconductor Co., Ltd. RTS5209 PCI Express Card Reader (rev 01)


```

## USB

```CMD
Dispositivos USB
lsusb
Retorno:
Bus 001 Device 002: ID 04f2:b367 Chicony Electronics Co., Ltd WebCam
Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
Bus 005 Device 003: ID 04ca:3004 Lite-On Technology Corp. Bluetooth USB Host Controller
Bus 005 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub
Bus 004 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub
Bus 003 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub
Bus 002 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub
```

## HD

```CMD
lista de dispositivos SATA e SCSI
sudo apt install lsscsi
lsscsi
Retorno:
[0:0:0:0]    disk    ATA      SanDisk SSD PLUS 00RL  /dev/sda

Listar discos
lsblk
Retorno:
NAME                      MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
loop1                       7:1    0  63,3M  1 loop /snap/core20/1852
loop2                       7:2    0  79,9M  1 loop /snap/lxd/22923
loop3                       7:3    0 111,9M  1 loop /snap/lxd/24322
loop4                       7:4    0  53,3M  1 loop /snap/snapd/19457
loop5                       7:5    0  40,8M  1 loop /snap/snapd/20092
loop6                       7:6    0  63,5M  1 loop /snap/core20/2015
sda                         8:0    0 111,8G  0 disk
├─sda1                      8:1    0     1M  0 part
├─sda2                      8:2    0     2G  0 part /boot
└─sda3                      8:3    0 109,8G  0 part
  └─ubuntu--vg-ubuntu--lv 253:0    0  54,9G  0 lvm  /
```

### S.M.A.R.T

```CMD
https://www.hardware.com.br/livros/ferramentas-linux/monitorando-saude-com-smart.html

sudo apt install smartmontools
sudo smartctl -s on /dev/sda

sudo smartctl -t short /dev/sda
sudo smartctl -l selftest /dev/sda
Retorno:
smartctl 7.2 2020-12-30 r5155 [x86_64-linux-5.15.0-83-generic] (local build)
Copyright (C) 2002-20, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF READ SMART DATA SECTION ===
SMART Self-test log structure revision number 1
Num  Test_Description    Status                  Remaining  LifeTime(hours)  LBA_of_first_error
# 1  Short offline       Completed without error       00%      1214         -

sudo smartctl -H /dev/sda
Retorno:
smartctl 7.2 2020-12-30 r5155 [x86_64-linux-5.15.0-83-generic] (local build)
Copyright (C) 2002-20, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF READ SMART DATA SECTION ===
SMART overall-health self-assessment test result: PASSED

https://grafana.com/grafana/dashboards/10530-s-m-a-r-t-disk-monitoring-for-prometheus-dashboard/
https://grafana.com/grafana/dashboards/10531-s-m-a-r-t-disk-monitoring-for-prometheus-errorboard/
https://grafana.com/grafana/dashboards/13654-s-m-a-r-t-dashboard/

sudo smartctl -i /dev/sda
sudo smartctl -a /dev/sda

```

### Serial

```CMD
sudo hdparm -i /dev/sda
```

### Informações

```CMD
sudo lshw -C disk
```

### Relatorio

```CMD
sudo apt install smartmontools
sudo smartctl -s on /dev/sda
sudo hdparm -i /dev/sda
sudo smartctl -t [short|long] /dev/sda
sudo smartctl -H /dev/sda
sudo smartctl -i /dev/sda
sudo smartctl -a /dev/sda
```
