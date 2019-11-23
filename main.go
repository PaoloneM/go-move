package main
 
import (
	"log"
	"os"
)
 
func main() {
	log.Println("start")
	oldLocation := "/dir1/test.txt"
	newLocation := "/dir2/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("end")
}