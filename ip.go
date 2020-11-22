package main

import(
	"fmt"
	"os"
	"log"
	"io"
	"os/exec"
	"strings"
)
const (
	version = "System info:<br>ip-code compiled on this date:::Sun Nov 22 04:02:40 2020"

)

func system_command(command ...string) int{
	prog,err:= exec.LookPath(command[0])
	fmt.Println(prog," ",command[1])
	if err != nil {
		fmt.Println("ERROR:",err)
		return 1
	}
	//fmt.Println("<pre>")
	cmd := &exec.Cmd {
		Path: prog,
			Args: []string{command[0],command[1]},
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR:",err)
	}
	//fmt.Println("</pre>")
	return 0
}
			
func cat(fname string) {

	fh, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = io.Copy(os.Stdout, fh)
	if err != nil {
		log.Fatal(err)
	}
}

func color (color string) {
	if color == "green" {
		fmt.Println("<span class=spanip>")
	}
}
		
func main() {
	//    os.Setenv("FOO", "1")
	fmt.Println("Content-type: text/html")
	fmt.Println("")
	cat("../html/ip.html")
	fmt.Print("REMOTE_ADDR:")
	color("green")
	fmt.Print(os.Getenv("REMOTE_ADDR"),"<br>")
	//	fmt.Println("QUERY_STRING:", os.Getenv("QUERY_STRING"), "<br>")
	// for _, e := range os.Environ() {
	// 	//        pair := strings.SplitN(e, "=", 2)
	// 	//        p(pair[0],"=",pair[1])
	// 	fmt.Println(e,"<br>")
	// }
	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"version")){
		fmt.Println(version,"<br>")
	}
	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"host")){
		fmt.Println("QUERY_STRING:", os.Getenv("QUERY_STRING"),"<br>")
		system_command("host",os.Getenv("REMOTE_ADDR"))
	}
	//	system_command("ls")
	//	fmt.Println("</p>")
	fmt.Println("</div>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
