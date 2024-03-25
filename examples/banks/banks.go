package examples

import (
	"fmt"

	"github.com/JVMoreiraD/go-brasilapi/cmd/libs/banks"
)

func main() {
	// A função abaixo retorna todos os bancos nacionais
	b, _ := banks.GetAllBanks()

	fmt.Println(b)

}
