# About

kafkacli is a simple Apache Kafka client for debugging.

features
 - support SASL/SCRAM authentication mechanism
 - pure golang library 
 - cross-platform supports MS-Windows, macOS and Linux


build binary for Linux on win10

    build-linux.bat    
    
    
build binary for Linux on linux/macOS

    make -f Makefile
    
    
build binary for win10

    build.bat
    
    
usage

    ./kafkacli.bin -b 127.0.0.1:9092 -t mytopic -o first
    ./kafkacli.bin -b 127.0.0.1:9092 -t mytopic -g mygrp -u myuser -p mypass
    
