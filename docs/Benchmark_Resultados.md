# Benchmark_Resultados

## P6200

```CMD
Comando
7z b -mmt1

7-Zip [64] 16.02 : Copyright (c) 1999-2016 Igor Pavlov : 2016-05-21
p7zip Version 16.02 (locale=pt_PT.UTF-8,Utf16=on,HugeFiles=on,64 bits,1 CPU Intel(R) Pentium(R) CPU        P6200  @ 2.13GHz (20655),ASM)

Intel(R) Pentium(R) CPU        P6200  @ 2.13GHz (20655)
CPU Freq: - - - - - - - - -

RAM size:    3731 MB,  # CPU hardware threads:   1
RAM usage:    435 MB,  # Benchmark threads:      1

                       Compressing  |                  Decompressing
Dict     Speed Usage    R/U Rating  |      Speed Usage    R/U Rating
         KiB/s     %   MIPS   MIPS  |      KiB/s     %   MIPS   MIPS

22:       2162    99   2114   2104  |      25064    99   2153   2140
23:       2055    99   2105   2094  |      24484    98   2152   2119
24:       1987    99   2149   2137  |      24358    99   2150   2138
25:       1899    99   2185   2169  |      23916   100   2139   2129
----------------------------------  | ------------------------------
Avr:              99   2138   2126  |               99   2149   2132
Tot:              99   2144   2129

```

```CMD
Comando
7z b

7-Zip [64] 16.02 : Copyright (c) 1999-2016 Igor Pavlov : 2016-05-21
p7zip Version 16.02 (locale=pt_PT.UTF-8,Utf16=on,HugeFiles=on,64 bits,1 CPU Intel(R) Pentium(R) CPU        P6200  @ 2.13GHz (20655),ASM)

Intel(R) Pentium(R) CPU        P6200  @ 2.13GHz (20655)
CPU Freq: - - - - 128000000 - - - -

RAM size:    3731 MB,  # CPU hardware threads:   1
RAM usage:    435 MB,  # Benchmark threads:      1

                       Compressing  |                  Decompressing
Dict     Speed Usage    R/U Rating  |      Speed Usage    R/U Rating
         KiB/s     %   MIPS   MIPS  |      KiB/s     %   MIPS   MIPS

22:       2160    99   2112   2101  |      24821    98   2152   2119
23:       2051   100   2100   2090  |      24736    99   2154   2141
24:       1977    99   2139   2127  |      24359    99   2150   2139
25:       1886    99   2178   2154  |      23890    99   2138   2126
----------------------------------  | ------------------------------
Avr:              99   2132   2118  |               99   2149   2131
Tot:              99   2140   2125

```

```CMD
Comando:
sysbench cpu --threads=2 run

Running the test with following options:
Number of threads: 2
Initializing random number generator from current time


Prime numbers limit: 10000

Initializing worker threads...

Threads started!

CPU speed:
    events per second:   619.04

General statistics:
    total time:                          10.0016s
    total number of events:              6193

Latency (ms):
         min:                                    1.61
         avg:                                    3.23
         max:                                   12.43
         95th percentile:                        5.57
         sum:                                19997.96

Threads fairness:
    events (avg/stddev):           3096.5000/0.50
    execution time (avg/stddev):   9.9990/0.00
```

```CMD
Comando
sysbench memory run

Running the test with following options:
Number of threads: 1
Initializing random number generator from current time


Running memory speed test with the following options:
  block size: 1KiB
  total size: 102400MiB
  operation: write
  scope: global

Initializing worker threads...

Threads started!

Total operations: 32001079 (3199242.25 per second)

31251.05 MiB transferred (3124.26 MiB/sec)


General statistics:
    total time:                          10.0001s
    total number of events:              32001079

Latency (ms):
         min:                                    0.00
         avg:                                    0.00
         max:                                    5.52
         95th percentile:                        0.00
         sum:                                 5097.28

Threads fairness:
    events (avg/stddev):           32001079.0000/0.00
    execution time (avg/stddev):   5.0973/0.00
```

```CMD
Comando
sysbench fileio --file-test-mode=seqwr run

sysbench 1.0.18 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 1
Initializing random number generator from current time


Extra file open flags: (none)
128 files, 16MiB each
2GiB total file size
Block size 16KiB
Periodic FSYNC enabled, calling fsync() each 100 requests.
Calling fsync() at the end of test, Enabled.
Using synchronous I/O mode
Doing sequential write (creation) test
Initializing worker threads...

Threads started!


File operations:
    reads/s:                      0.00
    writes/s:                     790.74
    fsyncs/s:                     1012.94

Throughput:
    read, MiB/s:                  0.00
    written, MiB/s:               12.36

General statistics:
    total time:                          10.1146s
    total number of events:              18120

Latency (ms):
         min:                                    0.02
         avg:                                    0.55
         max:                                  355.74
         95th percentile:                        0.33
         sum:                                10056.14

Threads fairness:
    events (avg/stddev):           18120.0000/0.00
    execution time (avg/stddev):   10.0561/0.00
```

```CMD
Programa
s-tui

Monitor: 931 MHz
Estresse: 2127 MHz

```

```CMD
Comando
fio --name=seqread --readwrite=read --direct=1 --ioengine=libaio --bs=1M --size=2000M

seqread: (g=0): rw=read, bs=(R) 1024KiB-1024KiB, (W) 1024KiB-1024KiB, (T) 1024KiB-1024KiB, ioengine=libaio, iodepth=1
fio-3.16
Starting 1 process
seqread: Laying out IO file (1 file / 2000MiB)
Jobs: 1 (f=1): [R(1)][100.0%][r=70.1MiB/s][r=70 IOPS][eta 00m:00s]
seqread: (groupid=0, jobs=1): err= 0: pid=48103: Fri Jan 14 01:26:38 2022
  read: IOPS=73, BW=73.2MiB/s (76.7MB/s)(2000MiB/27331msec)
    slat (usec): min=65, max=819, avg=151.85, stdev=17.53
    clat (msec): min=3, max=147, avg=13.50, stdev= 3.64
     lat (msec): min=4, max=147, avg=13.66, stdev= 3.64
    clat percentiles (msec):
     |  1.00th=[   11],  5.00th=[   13], 10.00th=[   13], 20.00th=[   13],
     | 30.00th=[   13], 40.00th=[   13], 50.00th=[   15], 60.00th=[   15],
     | 70.00th=[   15], 80.00th=[   15], 90.00th=[   15], 95.00th=[   16],
     | 99.00th=[   16], 99.50th=[   17], 99.90th=[   73], 99.95th=[  148],
     | 99.99th=[  148]
   bw (  KiB/s): min=45056, max=83968, per=99.86%, avg=74825.28, stdev=6912.41, samples=54
   iops        : min=   44, max=   82, avg=73.06, stdev= 6.76, samples=54
  lat (msec)   : 4=0.10%, 10=0.10%, 20=99.55%, 50=0.15%, 100=0.05%
  lat (msec)   : 250=0.05%
  cpu          : usr=0.10%, sys=1.26%, ctx=2040, majf=0, minf=268
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=2000,0,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=73.2MiB/s (76.7MB/s), 73.2MiB/s-73.2MiB/s (76.7MB/s-76.7MB/s), io=2000MiB (2097MB), run=27331-27331msec

Disk stats (read/write):
    dm-0: ios=1999/45, merge=0/0, ticks=27020/824, in_queue=27844, util=99.58%, aggrios=4003/40, aggrmerge=0/5, aggrticks=52769/683, aggrin_queue=45100, aggrutil=99.53%
  sda: ios=4003/40, merge=0/5, ticks=52769/683, in_queue=45100, util=99.53%

```

```CMD
Comando
fio --name=seqwrite --readwrite=write --direct=1 --ioengine=libaio --bs=1M --size=2000M

seqwrite: (g=0): rw=write, bs=(R) 1024KiB-1024KiB, (W) 1024KiB-1024KiB, (T) 1024KiB-1024KiB, ioengine=libaio, iodepth=1
fio-3.16
Starting 1 process
seqwrite: Laying out IO file (1 file / 2000MiB)
Jobs: 1 (f=1): [W(1)][100.0%][w=80.1MiB/s][w=80 IOPS][eta 00m:00s]
seqwrite: (groupid=0, jobs=1): err= 0: pid=48155: Fri Jan 14 01:28:41 2022
  write: IOPS=71, BW=71.0MiB/s (75.5MB/s)(2000MiB/27791msec); 0 zone resets
    slat (usec): min=103, max=887, avg=222.87, stdev=28.96
    clat (msec): min=4, max=175, avg=13.66, stdev= 5.65
     lat (msec): min=4, max=176, avg=13.89, stdev= 5.67
    clat percentiles (msec):
     |  1.00th=[   10],  5.00th=[   13], 10.00th=[   13], 20.00th=[   13],
     | 30.00th=[   13], 40.00th=[   13], 50.00th=[   13], 60.00th=[   15],
     | 70.00th=[   15], 80.00th=[   15], 90.00th=[   15], 95.00th=[   17],
     | 99.00th=[   18], 99.50th=[   32], 99.90th=[   89], 99.95th=[  176],
     | 99.99th=[  176]
   bw (  KiB/s): min=55296, max=83968, per=100.00%, avg=73719.44, stdev=7817.59, samples=55
   iops        : min=   54, max=   82, avg=71.95, stdev= 7.61, samples=55
  lat (msec)   : 10=1.20%, 20=97.90%, 50=0.40%, 100=0.45%, 250=0.05%
  cpu          : usr=0.72%, sys=1.03%, ctx=2072, majf=0, minf=11
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=0,2000,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
  WRITE: bw=71.0MiB/s (75.5MB/s), 71.0MiB/s-71.0MiB/s (75.5MB/s-75.5MB/s), io=2000MiB (2097MB), run=27791-27791msec

Disk stats (read/write):
    dm-0: ios=0/2073, merge=0/0, ticks=0/28352, in_queue=28352, util=99.65%, aggrios=0/2050, aggrmerge=0/35, aggrticks=0/28516, aggrin_queue=26076, aggrutil=99.57%
  sda: ios=0/2050, merge=0/35, ticks=0/28516, in_queue=26076, util=99.57%
```

```CMD
Comando
fio --name=randrw --readwrite=randrw --direct=1 --ioengine=libaio --bs=4k --size=200M --group_reporting --numjobs=8

randrw: (g=0): rw=randrw, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=libaio, iodepth=1
...
fio-3.16
Starting 8 processes
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
Jobs: 8 (f=8): [m(8)][1.8%][r=284KiB/s,w=264KiB/s][r=71,w=66 IOPS][eta 03h:41m:1                                             Jobs: 8 (f=8): [m(8)][1.8%][r=304KiB/s,w=256KiB/s][r=76,w=64 IOPS][eta 03h:40m:3                                             Jobs: 8 (f=8): [m(8)][1.8%][r=404KiB/s,w=384KiB/s][r=101,w=96 IOPS][eta 03h:41m:                                             Jobs: 8 (f=8): [m(8)][1.9%][r=352KiB/s,w=376KiB/s][r=88,w=94 IOPS][eta 03h:41m:0                                             Jobs: 8 (f=8): [m(8)][1.9%][r=384KiB/s,w=480KiB/s][r=96,w=120 IOPS][eta 03h:41m:                                             Jobs: 8 (f=8): [m(8)][1.9%][r=376KiB/s,w=420KiB/s][r=94,w=105 IOPS][eta 03h:41m:                                             Jobs: 1 (f=1): [m(1),_(7)][99.9%][r=316KiB/s,w=284KiB/s][r=79,w=71 IOPS][eta 00m:04s]
randrw: (groupid=0, jobs=8): err= 0: pid=48207: Fri Jan 14 02:18:13 2022
  read: IOPS=71, BW=286KiB/s (293kB/s)(800MiB/2868284msec)
    slat (usec): min=13, max=4058, avg=43.37, stdev=11.57
    clat (usec): min=149, max=2594.2k, avg=79113.24, stdev=172103.82
     lat (usec): min=173, max=2594.3k, avg=79157.73, stdev=172103.75
    clat percentiles (msec):
     |  1.00th=[    4],  5.00th=[    7], 10.00th=[   10], 20.00th=[   14],
     | 30.00th=[   18], 40.00th=[   23], 50.00th=[   31], 60.00th=[   43],
     | 70.00th=[   62], 80.00th=[   97], 90.00th=[  176], 95.00th=[  288],
     | 99.00th=[  709], 99.50th=[ 1250], 99.90th=[ 2039], 99.95th=[ 2106],
     | 99.99th=[ 2232]
   bw (  KiB/s): min=   52, max= 1823, per=100.00%, avg=437.23, stdev=36.29, samples=30150
   iops        : min=    8, max=  455, avg=108.39, stdev= 9.07, samples=30150
  write: IOPS=71, BW=285KiB/s (292kB/s)(800MiB/2868284msec); 0 zone resets
    slat (usec): min=11, max=34100, avg=45.74, stdev=91.20
    clat (usec): min=124, max=903079, avg=4774.17, stdev=32832.01
     lat (usec): min=206, max=903123, avg=4821.01, stdev=32836.50
    clat percentiles (usec):
     |  1.00th=[   241],  5.00th=[   265], 10.00th=[   277], 20.00th=[   306],
     | 30.00th=[   355], 40.00th=[   445], 50.00th=[   619], 60.00th=[  1287],
     | 70.00th=[  1500], 80.00th=[  2737], 90.00th=[  2966], 95.00th=[  3163],
     | 99.00th=[125305], 99.50th=[270533], 99.90th=[471860], 99.95th=[574620],
     | 99.99th=[750781]
   bw (  KiB/s): min=   52, max= 2573, per=100.00%, avg=480.92, stdev=45.25, samples=27228
   iops        : min=    8, max=  641, avg=119.31, stdev=11.31, samples=27228
  lat (usec)   : 250=1.10%, 500=21.49%, 750=4.12%, 1000=0.60%
  lat (msec)   : 2=11.97%, 4=10.04%, 10=5.51%, 20=11.76%, 50=14.88%
  lat (msec)   : 100=8.20%, 250=7.02%, 500=2.35%, 750=0.52%, 1000=0.16%
  lat (msec)   : 2000=0.11%, >=2000=0.19%
  cpu          : usr=0.04%, sys=0.14%, ctx=411714, majf=0, minf=109
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=204901,204699,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=286KiB/s (293kB/s), 286KiB/s-286KiB/s (293kB/s-293kB/s), io=800MiB (839MB), run=2868284-2868284msec
  WRITE: bw=285KiB/s (292kB/s), 285KiB/s-285KiB/s (292kB/s-292kB/s), io=800MiB (838MB), run=2868284-2868284msec

Disk stats (read/write):
    dm-0: ios=204943/210157, merge=0/0, ticks=16275084/1293940, in_queue=17569024, util=100.00%, aggrios=204919/208764, aggrmerge=32/1408, aggrticks=16212996/1290116, aggrin_queue=16852128, aggrutil=100.00%
  sda: ios=204919/208764, merge=32/1408, ticks=16212996/1290116, in_queue=16852128, util=100.00%
```

## Atom_N455

```CMD
Comando
7z b -mmt1

7-Zip [64] 16.02 : Copyright (c) 1999-2016 Igor Pavlov : 2016-05-21
p7zip Version 16.02 (locale=en_US.UTF-8,Utf16=on,HugeFiles=on,64 bits,2 CPUs Intel(R) Atom(TM) CPU N455   @ 1.66GHz (106CA),ASM)

         Intel(R) Atom(TM) CPU N455   @ 1.66GHz (106CA)
CPU Freq: 21333333 16000000 21333333 32000000 64000000 128000000 256000000 341333333 1024000000

RAM size:    1971 MB,  # CPU hardware threads:   2
RAM usage:    435 MB,  # Benchmark threads:      1

                       Compressing  |                  Decompressing
Dict     Speed Usage    R/U Rating  |      Speed Usage    R/U Rating
         KiB/s     %   MIPS   MIPS  |      KiB/s     %   MIPS   MIPS

22:        755   100    737    735  |      11674   100    999    997
23:        704   100    720    718  |      11531   100   1001    998
24:        676   100    729    727  |      11370   100   1001    998
25:        653   100    748    746  |      11168   100    997    994
----------------------------------  | ------------------------------
Avr:             100    734    731  |              100    999    997
Tot:             100    866    864

```

```CMD
Comando
7z b

7-Zip [64] 16.02 : Copyright (c) 1999-2016 Igor Pavlov : 2016-05-21
p7zip Version 16.02 (locale=en_US.UTF-8,Utf16=on,HugeFiles=on,64 bits,2 CPUs Intel(R) Atom(TM) CPU N455   @ 1.66GHz (106CA),ASM)

         Intel(R) Atom(TM) CPU N455   @ 1.66GHz (106CA)
CPU Freq: 21333333 32000000 32000000 21333333 64000000 128000000 256000000 512000000 1024000000

RAM size:    1971 MB,  # CPU hardware threads:   2
RAM usage:    441 MB,  # Benchmark threads:      2

                       Compressing  |                  Decompressing
Dict     Speed Usage    R/U Rating  |      Speed Usage    R/U Rating
         KiB/s     %   MIPS   MIPS  |      KiB/s     %   MIPS   MIPS

22:       1056   161    640   1028  |      18110   197    786   1546
23:       1015   163    634   1035  |      17915   197    788   1551
24:       1006   168    645   1083  |      17727   197    790   1556
25:        986   171    658   1126  |      17414   197    788   1550
----------------------------------  | ------------------------------
Avr:             166    644   1068  |              197    788   1551
Tot:             181    716   1309

```

```CMD
Comando
sysbench cpu --threads=2 run

Running the test with following options:
Number of threads: 2
Initializing random number generator from current time


Prime numbers limit: 10000

Initializing worker threads...

Threads started!

CPU speed:
    events per second:   123.16

General statistics:
    total time:                          10.0159s
    total number of events:              1234

Latency (ms):
         min:                                   12.32
         avg:                                   16.23
         max:                                   38.11
         95th percentile:                       16.12
         sum:                                20022.21

Threads fairness:
    events (avg/stddev):           617.0000/4.00
    execution time (avg/stddev):   10.0111/0.00

```

```CMD
Comando
sysbench memory run

sysbench 1.0.18 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 1
Initializing random number generator from current time


Running memory speed test with the following options:
  block size: 1KiB
  total size: 102400MiB
  operation: write
  scope: global

Initializing worker threads...

Threads started!

Total operations: 1639357 (163877.31 per second)

1600.93 MiB transferred (160.04 MiB/sec)


General statistics:
    total time:                          10.0002s
    total number of events:              1639357

Latency (ms):
         min:                                    0.00
         avg:                                    0.00
         max:                                    0.51
         95th percentile:                        0.00
         sum:                                 3599.56

Threads fairness:
    events (avg/stddev):           1639357.0000/0.00
    execution time (avg/stddev):   3.5996/0.00
```

```CMD
Comando
sysbench fileio --file-test-mode=seqwr run

sysbench 1.0.18 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 1
Initializing random number generator from current time


Extra file open flags: (none)
128 files, 16MiB each
2GiB total file size
Block size 16KiB
Periodic FSYNC enabled, calling fsync() each 100 requests.
Calling fsync() at the end of test, Enabled.
Using synchronous I/O mode
Doing sequential write (creation) test
Initializing worker threads...

Threads started!


File operations:
    reads/s:                      0.00
    writes/s:                     656.19
    fsyncs/s:                     851.76

Throughput:
    read, MiB/s:                  0.00
    written, MiB/s:               10.25

General statistics:
    total time:                          10.0547s
    total number of events:              15039

Latency (ms):
         min:                                    0.09
         avg:                                    0.66
         max:                                  143.78
         95th percentile:                        0.47
         sum:                                 9893.21

Threads fairness:
    events (avg/stddev):           15039.0000/0.00
    execution time (avg/stddev):   9.8932/0.00
```

```CMD
Programa
s-tui

Monitor: 997 MHz
Estresse: 1662 MHz

```

```CMD
Comando
fio --name=seqread --readwrite=read --direct=1 --ioengine=libaio --bs=1M --size=2000M

seqread: (g=0): rw=read, bs=(R) 1024KiB-1024KiB, (W) 1024KiB-1024KiB, (T) 1024KiB-1024KiB, ioengine=libaio, iodepth=1
fio-3.16
Starting 1 process
seqread: Laying out IO file (1 file / 2000MiB)
Jobs: 1 (f=1): [R(1)][100.0%][r=69.0MiB/s][r=69 IOPS][eta 00m:00s]
seqread: (groupid=0, jobs=1): err= 0: pid=34167: Fri Jan 14 01:27:39 2022
  read: IOPS=65, BW=65.9MiB/s (69.1MB/s)(2000MiB/30356msec)
    slat (usec): min=194, max=2873, avg=320.46, stdev=69.62
    clat (usec): min=8494, max=92203, avg=14828.22, stdev=5443.24
     lat (usec): min=8905, max=92477, avg=15154.48, stdev=5442.15
    clat percentiles (usec):
     |  1.00th=[13435],  5.00th=[13435], 10.00th=[13566], 20.00th=[13566],
     | 30.00th=[13566], 40.00th=[13698], 50.00th=[13698], 60.00th=[14484],
     | 70.00th=[14615], 80.00th=[14615], 90.00th=[15401], 95.00th=[16057],
     | 99.00th=[38011], 99.50th=[58983], 99.90th=[90702], 99.95th=[91751],
     | 99.99th=[91751]
   bw (  KiB/s): min=28672, max=73728, per=99.96%, avg=67436.40, stdev=7569.78, samples=60
   iops        : min=   28, max=   72, avg=65.78, stdev= 7.41, samples=60
  lat (msec)   : 10=0.05%, 20=97.75%, 50=1.60%, 100=0.60%
  cpu          : usr=0.29%, sys=2.41%, ctx=2109, majf=0, minf=269
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=2000,0,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=65.9MiB/s (69.1MB/s), 65.9MiB/s-65.9MiB/s (69.1MB/s-69.1MB/s), io=2000MiB (2097MB), run=30356-30356msec

Disk stats (read/write):
    dm-0: ios=2033/16, merge=0/0, ticks=30332/296, in_queue=30628, util=99.66%, aggrios=2039/10, aggrmerge=0/6, aggrticks=30595/284, aggrin_queue=26836, aggrutil=99.50%
  sda: ios=2039/10, merge=0/6, ticks=30595/284, in_queue=26836, util=99.50%
```

```CMD
Comando
fio --name=seqwrite --readwrite=write --direct=1 --ioengine=libaio --bs=1M --size=2000M

seqwrite: (g=0): rw=write, bs=(R) 1024KiB-1024KiB, (W) 1024KiB-1024KiB, (T) 1024KiB-1024KiB, ioengine=libaio, iodepth=1
fio-3.16
Starting 1 process
seqwrite: Laying out IO file (1 file / 2000MiB)
Jobs: 1 (f=1): [W(1)][100.0%][w=64.1MiB/s][w=64 IOPS][eta 00m:00s]
seqwrite: (groupid=0, jobs=1): err= 0: pid=34182: Fri Jan 14 01:28:50 2022
  write: IOPS=63, BW=63.5MiB/s (66.6MB/s)(2000MiB/31475msec); 0 zone resets
    slat (usec): min=465, max=11681, avg=739.61, stdev=276.51
    clat (msec): min=8, max=186, avg=14.96, stdev= 9.53
     lat (msec): min=9, max=198, avg=15.71, stdev= 9.66
    clat percentiles (msec):
     |  1.00th=[   10],  5.00th=[   13], 10.00th=[   14], 20.00th=[   14],
     | 30.00th=[   14], 40.00th=[   14], 50.00th=[   14], 60.00th=[   15],
     | 70.00th=[   15], 80.00th=[   15], 90.00th=[   16], 95.00th=[   16],
     | 99.00th=[   56], 99.50th=[   90], 99.90th=[  174], 99.95th=[  188],
     | 99.99th=[  188]
   bw (  KiB/s): min=40960, max=73728, per=99.88%, avg=64990.63, stdev=7705.52, samples=62
   iops        : min=   40, max=   72, avg=63.35, stdev= 7.53, samples=62
  lat (msec)   : 10=2.10%, 20=94.25%, 50=2.50%, 100=0.90%, 250=0.25%
  cpu          : usr=1.60%, sys=3.66%, ctx=2274, majf=0, minf=11
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=0,2000,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
  WRITE: bw=63.5MiB/s (66.6MB/s), 63.5MiB/s-63.5MiB/s (66.6MB/s-66.6MB/s), io=2000MiB (2097MB), run=31475-31475msec

Disk stats (read/write):
    dm-0: ios=1/2065, merge=0/0, ticks=120/30700, in_queue=30820, util=99.64%, aggrios=1/4027, aggrmerge=0/48, aggrticks=121/50120, aggrin_queue=43764, aggrutil=99.56%
  sda: ios=1/4027, merge=0/48, ticks=121/50120, in_queue=43764, util=99.56%
```

```CMD
Comando
fio --name=randrw --readwrite=randrw --direct=1 --ioengine=libaio --bs=4k --size=200M --group_reporting --numjobs=8

randrw: (g=0): rw=randrw, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=libaio, iodepth=1
...
fio-3.16
Starting 8 processes
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
randrw: Laying out IO file (1 file / 200MiB)
Jobs: 1 (f=1): [m(1),_(7)][99.6%][r=496KiB/s,w=404KiB/s][r=124,w=101 IOPS][eta 00m:09s]
randrw: (groupid=0, jobs=8): err= 0: pid=34195: Fri Jan 14 02:05:21 2022
  read: IOPS=98, BW=392KiB/s (402kB/s)(800MiB/2089804msec)
    slat (usec): min=51, max=48758, avg=118.03, stdev=225.61
    clat (usec): min=166, max=1364.2k, avg=70404.16, stdev=77005.00
     lat (usec): min=248, max=1364.3k, avg=70527.34, stdev=77013.32
    clat percentiles (msec):
     |  1.00th=[    4],  5.00th=[    8], 10.00th=[   11], 20.00th=[   16],
     | 30.00th=[   23], 40.00th=[   33], 50.00th=[   45], 60.00th=[   60],
     | 70.00th=[   81], 80.00th=[  111], 90.00th=[  163], 95.00th=[  220],
     | 99.00th=[  384], 99.50th=[  439], 99.90th=[  550], 99.95th=[  617],
     | 99.99th=[  802]
   bw (  KiB/s): min=   56, max= 2147, per=100.00%, avg=408.16, stdev=26.63, samples=32011
   iops        : min=    8, max=  535, avg=101.20, stdev= 6.65, samples=32011
  write: IOPS=97, BW=392KiB/s (401kB/s)(800MiB/2089804msec); 0 zone resets
    slat (usec): min=55, max=31413, avg=133.70, stdev=248.01
    clat (usec): min=15, max=1522.2k, avg=7866.37, stdev=34640.23
     lat (usec): min=455, max=1522.3k, avg=8005.31, stdev=34673.84
    clat percentiles (usec):
     |  1.00th=[   465],  5.00th=[   594], 10.00th=[   717], 20.00th=[  1303],
     | 30.00th=[  2114], 40.00th=[  2376], 50.00th=[  2671], 60.00th=[  3130],
     | 70.00th=[  3818], 80.00th=[  5997], 90.00th=[ 10683], 95.00th=[ 17171],
     | 99.00th=[ 98042], 99.50th=[350225], 99.90th=[446694], 99.95th=[467665],
     | 99.99th=[541066]
   bw (  KiB/s): min=   56, max= 2484, per=100.00%, avg=427.81, stdev=35.90, samples=30507
   iops        : min=    8, max=  620, avg=106.12, stdev= 8.97, samples=30507
  lat (usec)   : 20=0.01%, 250=0.01%, 500=1.29%, 750=4.36%, 1000=2.07%
  lat (msec)   : 2=6.03%, 4=22.66%, 10=12.81%, 20=11.99%, 50=15.17%
  lat (msec)   : 100=11.58%, 250=9.87%, 500=2.08%, 750=0.09%, 1000=0.01%
  lat (msec)   : 2000=0.01%
  cpu          : usr=0.12%, sys=0.42%, ctx=415550, majf=0, minf=103
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=204901,204699,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=392KiB/s (402kB/s), 392KiB/s-392KiB/s (402kB/s-402kB/s), io=800MiB (839MB), run=2089804-2089804msec
  WRITE: bw=392KiB/s (401kB/s), 392KiB/s-392KiB/s (401kB/s-401kB/s), io=800MiB (838MB), run=2089804-2089804msec

Disk stats (read/write):
    dm-0: ios=204912/206776, merge=0/0, ticks=14433048/1826364, in_queue=16259412, util=100.00%, aggrios=204922/205961, aggrmerge=0/830, aggrticks=14439486/1824801, aggrin_queue=15432904, aggrutil=100.00%
  sda: ios=204922/205961, merge=0/830, ticks=14439486/1824801, in_queue=15432904, util=100.00%
```
