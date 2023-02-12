package contract

//go:generate docker run --rm -v ${PWD}:/root ethereum/solc:0.8.10 --overwrite --bin --optimize --abi /root/Events.sol -o /root/build
//go:generate abigen --abi=./build/Events.abi --pkg contract --out=output.go
