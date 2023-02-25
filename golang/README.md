brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
vi ~/.bash_profile
export GOPATH=$HOME/go
export PATH=$JAVA_HOME/bin:$M2_HOME/bin:$GOPATH/bin:$PATH;
