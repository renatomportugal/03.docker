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
Programa
s-tui

Monitor: 931 MHz
Estresse: 2127 MHz

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
Programa
s-tui

Monitor: 997 MHz
Estresse: 1662 MHz

```
