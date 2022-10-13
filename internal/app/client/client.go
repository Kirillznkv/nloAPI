package client

import (
	"context"
	"github.com/Kirillznkv/nloAPI/internal/pkg/model"
	"github.com/Kirillznkv/nloAPI/internal/pkg/store"
	pb "github.com/Kirillznkv/nloAPI/pkg/api"
	"gonum.org/v1/gonum/stat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

type Client struct {
	store *store.Store
}

func New(config *store.Config) *Client {
	return &Client{
		store: store.New(config),
	}
}

func (s *Client) Start() error {
	defer s.store.Close()
	if err := s.configureStore(); err != nil {
		return err
	}

	conn := connectToServ("127.0.0.1:8080")
	defer conn.Close()

	s.findAnomaly(conn)

	return nil
}

func (s *Client) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return s.store.MigrateUP()
}

func connectToServ(addr string) *grpc.ClientConn {
	opts := make([]grpc.DialOption, 0, 2)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}

func (s *Client) findAnomaly(conn *grpc.ClientConn) {
	forecastArr := make([]float64, 1000)
	var mean, std float64
	ch := make(chan *pb.Response)
	i := 0
	go execDo(ch, conn)
	for {
		data := <-ch
		if i < 1000 {
			forecastArr[i] = data.Frequency
			mean, std = stat.MeanStdDev(forecastArr[:i+1], nil)
			log.Printf("%d: (mean:%f std:%f)", i, mean, std)
			log.Printf("%s %f %s\n", data.SessionId, data.Frequency, data.Timestamp.AsTime())
			if i == 1000 {
				forecastArr = nil
			}
			i++
		} else if data.Frequency < (mean-std) || data.Frequency > (mean+std) {
			if err := s.store.Anomaly().Create(&model.Anomaly{
				SessionId: data.SessionId,
				Frequency: data.Frequency,
				Timestamp: data.Timestamp.AsTime(),
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func execDo(ch chan *pb.Response, conn *grpc.ClientConn) {
	defer close(ch)
	stream := getDoStream(conn)
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ch <- data
	}
}

func getDoStream(conn *grpc.ClientConn) pb.Nlo_DoClient {
	c := pb.NewNloClient(conn)
	stream, err := c.Do(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("%v.Do(_) = _, %v", c, err)
	}
	return stream
}
