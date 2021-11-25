package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	todo "github.com/DymaSV/todo/proto/todo"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	srv := grpc.NewServer()
	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln("could not listen to :8888; %v", err)
	}
	log.Fatal(srv.Serve(l))
}

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

type taskServer struct{}
type length int64

var endianness = binary.LittleEndian

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return &todo.TaskList{}, fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	var tasks todo.TaskList
	for {
		if len(b) == 0 {
			return &tasks, nil
		} else if len(b) < sizeOfLength {
			return nil, fmt.Errorf("remaining odd %d bytes, what to do?", err)
		}

		var l length
		if err = binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return nil, fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]
		var task todo.Task
		if err := proto.Unmarshal(b[:l], &task); err == io.EOF {
			return nil, fmt.Errorf("could not read task: %v", err)
		}
		b = b[l:]
		tasks.Tasks = append(tasks.Tasks, &task)

		if task.Done {
			fmt.Printf("OK")
		} else {
			fmt.Printf("Oups...")
		}
		fmt.Printf(" %s\n", task.Text)
	}
}
