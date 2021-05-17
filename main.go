package main


import(
	etherscan "github.com/nanmu42/etherscan-api"
	"fmt"
)

	const APIKEY string = "22B45ENZRGIBCCKXNYUXKQCE1GJVRMEUIH"


func main(){
	client := etherscan.New(, APIKEY)
	fmt.Println("Client active:", client)

	balance, err := client.AccountBalance("0x926d76bF9EEc5db5E5baA00043015D4367B6BDEE")
	if err != nil{
		panic(err)
	}
	fmt.Println("Balance:", balance)

	price, err := client.EtherLatestPrice()
	if err != nil{
		panic(err)
	}
	fmt.Println("Price:", price)
}