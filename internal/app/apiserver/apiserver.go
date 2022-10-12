package apiserver

import (
	"log"
	"math/rand"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Kirillznkv/nloAPI/pkg/api"
)

//type logData struct {
//	mean, std, freq float64
//}

type ImplementedNloServer struct {
	pb.UnimplementedNloServer
}

func (t *ImplementedNloServer) Do(r *pb.Request, srv pb.Nlo_DoServer) error {
	//var ld logData
	//ld.mean = rand.Float64()*21 - 10
	//ld.std = rand.Float64()*1.2 + 0.3
	var res pb.Response
	res.SessionId = uuid.New().String()
	dist := distuv.Normal{
		Mu:    rand.Float64()*21 - 10,
		Sigma: rand.Float64()*1.2 + 0.3,
	}
	for {
		res.Frequency = dist.Rand()
		res.Timestamp = timestamppb.Now()
		//ld.freq = res.Frequency
		if err := srv.Send(&res); err != nil {
			return nil
		}
		log.Printf("m: %f s: %f\t\t->\t%f\n", dist.Mu, dist.Sigma, res.Frequency)
		//t.myLog(ld)
	}
	return nil
}

//func (t *ImplementedNloServer) myLog(ld logData) {
//	fmt.Printf("m: %f s: %f\t\t->\t%f\n", ld.mean, ld.std, ld.freq)
//}
