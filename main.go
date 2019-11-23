package main
 
import (
	"log"
	"os"
)
 
func main() {
	log.Println("start")
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("end")
}