package main

import (
	"context"
	"log"
	"strconv"

	unary "github.com/abhsinh2/study/golang/grpc/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}(cc)

	client := unary.NewPhoneClient(cc)

	log.Println("Starting to do a Unary Phone RPC...")

	GetServerContactName(client)
	GetServerContactNum(client)
	ListContacts(client)
}

func GetServerContactName(client unary.PhoneClient) {
	req := &unary.GetContactNameRequest{
		Number: "220123651",
	}

	res, err := client.GetContactName(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactName: ", res.Firstname+" "+res.Lastname)
}

func GetServerContactNum(client unary.PhoneClient) {
	req := &unary.GetContactNumRequest{
		Firstname: "Klay",
		Lastname:  "Thompson",
	}

	res, err := client.GetContactNum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}

	log.Println("Response from GetContactNum: ", strconv.FormatUint(res.Num, 10)+" "+res.Result)
}

func ListContacts(client unary.PhoneClient) {
	res, err := client.ListContacts(context.Background(), &unary.ListContactsRequest{})
	if err != nil {
		log.Fatalf("error while calling Unary Phone RPC: %v", err)
	}
	log.Println("Response from ListContacts: sum is ", res.Sum)
	log.Println("contacts list is ", res.Contacts)
}
