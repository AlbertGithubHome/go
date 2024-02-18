#!/bin/bash

protoc --proto_path=srcproto --go_out=.. examplemsg.proto