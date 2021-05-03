protoc --go_out=$GOPATH/src/ *.proto
#protoc  --go_opt=paths=source_relative   --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb