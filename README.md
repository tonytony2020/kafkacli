# About

![Go](https://github.com/tonytony2020/kafkacli/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tonytony2020/kafkacli)](https://goreportcard.com/report/github.com/tonytony2020/kafkacli)
[![LICENSE](https://img.shields.io/github/license/tonytony2020/kafkacli.svg?style=flat-square)](https://github.com/tonytony2020/kafkacli/blob/master/LICENSE)


kafkacli is a simple Apache Kafka client for debugging.

features
 - support SASL/SCRAM authentication mechanism
 - pure golang library 
 - cross-platform supports MS-Windows, macOS and Linux


## Building

build binary for Linux on win10

    build-linux.bat    
    
    
build binary for Linux on linux/macOS

    make -f Makefile
    
    
build binary for win10

    build.bat
    


## Usage


    Usage of ./kafkacli.bin:
      -b string
            brokers (default "127.0.0.1:9092")
      -g string
            groupid
      -m int
            Min number of bytes to fetch from kafka in each request. (default 10000)
      -n int
            Max number of bytes to fetch from kafka in each request. (default 10000000)
      -o string
            offset first or last (default "last")
      -p string
            password
      -r int
            retry after n seconds (default 15)
      -t string
            topic (default "test")
      -u string
            username


    ./kafkacli.bin -b 127.0.0.1:9092 -t mytopic -o first
    ./kafkacli.bin -b 127.0.0.1:9092 -t mytopic -g mygrp -u myuser -p mypass
    
