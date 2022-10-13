package apiserver

import (
	"log"
	"math/rand"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Kirillznkv/nloAPI/pkg/api"
)

type ImplementedNloServer struct {
	pb.UnimplementedNloServer
}

func (t *ImplementedNloServer) Do(r *pb.Request, srv pb.Nlo_DoServer) error {
	var res pb.Response
	res.SessionId = uuid.New().String()
	dist := distuv.Normal{
		Mu:    rand.Float64()*21 - 10,
		Sigma: rand.Float64()*1.2 + 0.3,
	}
	for {
		res.Frequency = dist.Rand()
		res.Timestamp = timestamppb.Now()
		if err := srv.Send(&res); err != nil {
			return nil
		}
		log.Printf("m: %f s: %f\t\t->\t%f\n", dist.Mu, dist.Sigma, res.Frequency)
	}
	return nil
}
