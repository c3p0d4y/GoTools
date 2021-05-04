package main 
import(
	"log"
	"net"
	"strconv"
	"time"
	"os"
)

var ipToScan = os.Args[1]
var minPort = 1
var maxPort = 1024 

func main() { 
	activeThreads := 0 
	donChannel := make(chan bool)

	for port := minPort; port <= maxPort; port++{
		go testTCPConnection(ipToScan, port, donChannel)
		activeThreads ++
	}
	// wait for all threads to finish 
	for activeThreads > 0 {
		<-donChannel
		activeThreads--
	}
}
func testTCPConnection(ip string, port int, donChannel chan bool){
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port),
	time.Second*10)

	if err == nil {
		log.Printf("Host %s has open port : %d\n", ip,port)

	}
	donChannel <- true 
}
