package main

import (
	"blixenkrone/spirii/internal/chargers"
	"blixenkrone/spirii/server/http"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type eventbus interface {
	produce(ctx context.Context, event any) error
	consume(ctx context.Context) error
	start(context.Context, chan os.Signal) // just for example
}

type httpServer interface {
	ListenAndServe() error
	ShutDown(context.Context) error
}

type memoryReadWriter[T any] interface {
	Read(ctx context.Context, id string) (T, error)
	Write(ctx context.Context, v T) error
}

type app struct {
	server   httpServer
	db       memoryReadWriter[chargers.MeterReading]
	consumer eventbus
}

type fakeConsumer struct {
	interval time.Duration
	data     chan chargers.MeterReading
}

func (f fakeConsumer) produce(ctx context.Context, data any) error {
	t := time.NewTicker(time.Second * f.interval)
	count := 0
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case t := <-t.C:
			count++
			fmt.Println("reading", t.String())
			f.data <- chargers.MeterReading{
				Timestamp:       t,
				MeterID:         strconv.Itoa(count),
				ConsumerID:      uuid.NewString(),
				MeterReadingVal: rand.Intn(10) + 1,
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (f fakeConsumer) consume(ctx context.Context) error {
	for v := range f.data {
		spew.Dump(v)
	}
	return nil
}

func (f fakeConsumer) start(ctx context.Context, done chan os.Signal) {
	go func() {
		f.produce(ctx, done)
	}()
	go func() {
		f.consume(ctx)
	}()
	fmt.Println("started consumer")
	<-done
	fmt.Println("stopped consumer")
}

func (a app) Start() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		cancel()
		log.Println("teardown complete")
	}()

	go func() {
		log.Printf("started server")
		if err := a.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	a.consumer.start(ctx, done)
	<-done

	if err := a.server.ShutDown(ctx); err != nil {
		panic(err)
	}
	log.Println("gracefully shutdown")

	return nil
}

func main() {
	l := logrus.New()

	dataCh := make(chan chargers.MeterReading)
	f := fakeConsumer{interval: 1, data: dataCh}

	fooDB := chargers.NewChargersDB()
	srv := http.NewServer(l, ":8080", fooDB)

	app := app{
		server:   srv,
		db:       fooDB,
		consumer: f,
	}

	if err := app.Start(); err != nil {
		l.Fatalln(err)
		os.Exit(1)
	}
}
