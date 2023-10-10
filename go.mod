module github.com/PretendoNetwork/luigis-mansion-2

go 1.19

require (
	github.com/PretendoNetwork/grpc-go v1.0.1
	github.com/PretendoNetwork/nex-go v1.0.41
	github.com/PretendoNetwork/nex-protocols-common-go v1.0.22
	github.com/PretendoNetwork/nex-protocols-go v1.0.54
	github.com/PretendoNetwork/plogger-go v1.0.4
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.17
	google.golang.org/grpc v1.56.1
)

require (
	github.com/fatih/color v1.15.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jwalton/go-supportscolor v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/superwhiskers/crunch/v3 v3.5.7 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace (
	github.com/PretendoNetwork/nex-go => ../../libraries/nex-go
	github.com/PretendoNetwork/nex-protocols-common-go => ../../libraries/nex-protocols-common-go
	github.com/PretendoNetwork/nex-protocols-go => ../../libraries/nex-protocols-go
)
