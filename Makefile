proto:
	protoc --go_out=./pb --go_opt=paths=source_relative deposit_money.proto
	protoc --go_out=./pb --go_opt=paths=source_relative threshold.proto
	protoc --go_out=./pb --go_opt=paths=source_relative wallet.proto