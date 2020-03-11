/*Packages*/
package main

/*Importations*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*Functions*/
func main()  {
	/* Variables */
	var names [3]string /*Array type string*/
	/*Logic to capture data*/
	for i := 0; i < 3; i++ {/*Cicle for*/
		/*Statement for capture by keyspace*/
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please tell me your name? ")
		data, _ := reader.ReadString('\n')
		data = strings.Replace(data, "\n", "", -1)
		names[i] = data
	}
	/*Show the captured*/
	fmt.Println("The names register are: ", names)
}