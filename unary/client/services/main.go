package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/MarceloBaeza/grpc-test-worklist/client/models"
	pb "github.com/MarceloBaeza/grpc-test-worklist/proto"
)

type ServiceStruct struct {
	Client pb.WorklistServiceClient
}

var instace *ServiceStruct
var once sync.Once

func NewServiceStruct(c pb.WorklistServiceClient) *ServiceStruct {
	once.Do(func() {
		instace = &ServiceStruct{
			Client: c,
		}
	})

	return instace
}
func (ss *ServiceStruct) NewWork(w http.ResponseWriter, r *http.Request) {
	log.Println("New Work init")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "error get body " + err.Error()
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var newWork models.Work
	err = json.Unmarshal(body, &newWork)
	if err != nil {
		msg := "error to parse data " + err.Error()
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	_, err = ss.Client.CreateWork(context.Background(), &pb.RequestCreateWork{
		NewWork: &pb.Work{
			Id:         generateID(),
			Name:       newWork.Name,
			Priority:   pb.Work_WorkPriority(newWork.Priority),
			StatusWork: pb.Work_WorkStatus(0),
		},
	})
	if err != nil {
		msg := "error creating new work " + err.Error()
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func generateID() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Int())
}
