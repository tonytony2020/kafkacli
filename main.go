package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/giant-stone/go/gutil"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"

	_ "github.com/segmentio/kafka-go/gzip"
	_ "github.com/segmentio/kafka-go/lz4"
	_ "github.com/segmentio/kafka-go/snappy"
	_ "github.com/segmentio/kafka-go/zstd"
)

var (
	brokers    string
	topic      string
	groupid    string
	offset     string
	retryafter int
	username   string
	password   string
	minbytes   int
	maxbytes   int
)

func main() {
	flag.StringVar(&brokers, "b", "127.0.0.1:9092", "brokers")
	flag.StringVar(&topic, "t", "test", "topic")
	flag.StringVar(&groupid, "g", "", "groupid")
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "password")
	flag.IntVar(&retryafter, "r", 15, "retry after n seconds")
	flag.StringVar(&offset, "o", "last", fmt.Sprintf("offset first or last"))
	flag.IntVar(&minbytes, "m", 10e3, "Min number of bytes to fetch from kafka in each request.")
	flag.IntVar(&maxbytes, "n", 10e6, "Max number of bytes to fetch from kafka in each request.")
	flag.Parse()

	dialer := kafka.Dialer{
		Timeout:   5 * time.Second,
		DualStack: true,
	}

	if username != "" && password != "" {
		// TODO: custom mechanism with flag option
		sm, err := scram.Mechanism(scram.SHA512, username, password)
		gutil.ExitOnErr(err)
		dialer.SASLMechanism = sm
	}

	cfg := kafka.ReaderConfig{
		Brokers:  strings.Split(brokers, ","),
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   &dialer,
	}

	if groupid != "" {
		cfg.GroupID = groupid
	}

	r := kafka.NewReader(cfg)

	if groupid == "" {
		var err error
		if offset == "first" {
			err = r.SetOffset(kafka.FirstOffset)
		} else if offset == "last" {
			err = r.SetOffset(kafka.LastOffset)
		}
		gutil.ExitOnErr(err)
	}

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[error] ReadMessage %v", err)
			time.Sleep(time.Second * time.Duration(retryafter))
			continue
		}

		fmt.Printf("part=%d offset=%d %s=%s \n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

}
