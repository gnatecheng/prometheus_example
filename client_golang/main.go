// Create: 2019/04/29 19:06:00 Change: 2019/04/30 16:04:05
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	api "github.com/prometheus/client_golang/api/prometheus"
	"github.com/prometheus/common/model"
	"golang.org/x/net/context"
	"log"
	"time"
)

func displayScalar(v *model.Scalar) {
	fmt.Printf("%s %s\n", v.Timestamp.String(), v.Value.String())
}

func displayVector(v model.Vector) {
	for _, i := range v {
		fmt.Printf("%s %s %s\n", i.Timestamp.String(), i.Metric.String(), i.Value.String())
	}
}

func displayMatrix(v model.Matrix) {
	for _, i := range v {
		fmt.Printf("%s\n", i.Metric.String())
		for _, j := range i.Values {
			fmt.Printf("\t%s %s\n", j.Timestamp.String(), j.Value.String())
		}
	}
}

func displayString(v *model.String) {
	fmt.Printf("%s %s\n", v.Timestamp.String(), v.Value)
}

func main() {
	config := api.Config{
		Address: "http://10.10.114.206:9090",
	}
	client, err := api.New(config)
	if err != nil {
		log.Fatal(err)
	}
	query_client := api.NewQueryAPI(client)

	ctx := context.TODO()
	value, err := query_client.Query(ctx, "kube_pod_info{node=\"t-k8s-node-10-187\"}", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	switch value.Type() {
	case model.ValNone:
		fmt.Println("None Type")
	case model.ValScalar:
		fmt.Println("Scalar Type")
		v, _ := value.(*model.Scalar)
		displayScalar(v)
	case model.ValVector:
		fmt.Println("Vector Type")
		v, _ := value.(model.Vector)
		displayVector(v)
	case model.ValMatrix:
		fmt.Println("Matrix Type")
		v, _ := value.(model.Matrix)
		displayMatrix(v)
	case model.ValString:
		fmt.Println("String Type")
		v, _ := value.(*model.String)
		displayString(v)
	default:
		fmt.Printf("Unknow Type")
	}
	return
}
