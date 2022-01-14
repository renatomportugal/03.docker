# Benchmark

## sysbench

```CMD
sudo apt install sysbench

sysbench cpu --threads=2 run
```

## s-tui

```CMD
sudo apt install s-tui stress

s-tui

```

## 7-Zip

```CMD
sudo apt install p7zip-full

To run single threaded benchmark, use the command below:
7z b -mmt1

Multi-threaded benchmarking can be run by using the following command:
7z b

Ver resultados de PCs em https://www.7-cpu.com/

```

## Fio

```CMD
sudo apt update && sudo apt install fio -y

f you get a message that fio cannot be found, it means you don’t have the universe repository enabled. You can enable it with apt install software-properties-common && add-apt-repository universe and then repeat the command above to install fio.


You can run a sequential read test with:

fio --name=seqread --readwrite=read --direct=1 --ioengine=libaio --bs=1M --size=2000M

Get more accurate results:

fio --name=seqread --readwrite=read --direct=1 --ioengine=libaio --bs=1M --size=8000M

To test for sequential write speed, run:

fio --name=seqwrite --readwrite=write --direct=1 --ioengine=libaio --bs=1M --size=2000M

To test how your cloud storage performs under the most stressful conditions, run this test:

fio --name=randrw --readwrite=randrw --direct=1 --ioengine=libaio --bs=4k --size=200M --group_reporting --numjobs=8

Just as above, increase --size if the test finishes too fast. In this case, bandwidth is less important – consider it secondary.
```

## Speedtest-cli

```CMD
sudo apt install speedtest-cli

speedtest

You can manually list servers in your country with a command like:
speedtest --list | grep -i germany

Select the number from the list, and pass it to the next command like the following line of code.
speedtest --server 4462
```
