#!/bin/bash

glide install
gin --path src --bin main-bin --port 8080 run main.go
