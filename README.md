go-logstash
===========

A golang package that sends newline delimited data via tcp to logstash.

This fork has been updated from the source repo, inorder to have a persistent TCP connection. Several other changes were done mainly to simplify the code and remove superflous code.

Example
-------
```go
import(
	"github.com/nitantsoni/go-logstash"
	"fmt"
)

func main() {
	host := net.TCPAddr {
        	IP:net.ParseIP("127.0.0.1"),
        	Port:9000,
	}
	l := &logstash.Logstash{Host:host,Connection:nil,}
	_, err := l.Connect()
	if err != nil {
		fmt.Println(err)
	}
	//Writeln takes []byte arrays instead of strings
	err = l.Writeln([]byte("{ 'foo' : 'bar' }"))
	if err != nil {
		fmt.Println(err)
	}
}
```
