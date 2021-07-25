#!bin/bash

protoc calculator/calculatorpb/calculator.proto --go-grpc_out=./calculator/ --go_out=./calculator/ 