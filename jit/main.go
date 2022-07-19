package main

import (
	"fmt"
	"log"

	zserio "github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/internal/model"
)

func main() {
	m, err := model.FromFilesystem("./example/schema/contacts.zs")
	if err != nil {
		fmt.Printf("parse schema: %v", err)
	}

	mm := &M{Model: m}

	bytes, err := zserio.Marshal(mm)
	if err != nil {
		log.Fatalf("error serializing address: %w", err)
	}

	fmt.Print(bytes)
}
