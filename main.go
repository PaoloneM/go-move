package main
 
import (
	"log"
	"os"
)
 
func main() {
	log.Debug("start")
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("end")
}