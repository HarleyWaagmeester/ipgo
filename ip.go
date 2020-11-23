// Web service utilizing a hosted language on HTML and CSS.


package main

import(
	"fmt"
	"os"
	"log"
	"io"
	"os/exec"
	"strings"
)

// The date is updated automatically by emacs.
const (
	version = "System info:<br>ip-code compiled on this date:::Mon Nov 23 16:14:22 2020"

)

func system_command(command ...string) int{
	prog,err:= exec.LookPath(command[0])
	//	fmt.Println(prog," ",command[1])
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


// Functions defining the IPGO language hosted on HTML and CSS.


// use flex for floats
func flexbox () {
	fmt.Println("<div class='flexbox';>")
}
func flexbox_off () {
	fmt.Println("</div>")
}
func floatbox () {
	fmt.Println("<div class='floatbox';>")
}
func floatbox_off () {
	fmt.Println("</div>")
}
// span colors
func color (color string) {
	if color == "green" {
		fmt.Println("<span class=span_green>")
	}
	if color == "off" {
		fmt.Println("</span>")
	}
}
func div33 () {
	fmt.Println("<div class='div33'>")
}
func div66 () {
	fmt.Println("<div class='div66;'>")
}
func div33_float_left () {
	fmt.Println("<div class='div33_float_left'>")
}
func div66_float_left () {
	fmt.Println("<div class='div66_float_left;'>")
}
// Shorthand for span_float_...
func float_left () {
	fmt.Println("<span class='float-left'>")
}
func float_right () {
	fmt.Println("<span class='float-right'>")
}
func float_off () {
	fmt.Println("</span>")
}
func br () {
	fmt.Println("<br>")
}
func div_off () {
	fmt.Println("</div>")
}

// Read in HTML files, create HTML elements. execute external programs.

func main() {
	fmt.Println("Content-type: text/html")
	fmt.Println("")
	cat("../html/ip.html")
	flexbox()
	div33()
	float_left()
	fmt.Print("REMOTE_ADDR:")
	float_off()
	float_right()
	color("green")
	fmt.Print(os.Getenv("REMOTE_ADDR"))
	color("off")
	float_off()
	br()
	//	div_off()
	//	flexbox_off()
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
		//	flexbox()
		//div33()
		float_left()
		fmt.Print("QUERY_STRING:")
		float_off()
		float_right()
		color("green")
		fmt.Print(os.Getenv("QUERY_STRING"))
		color("off")
		float_off()
		div_off()
		flexbox_off()

		//		flexbox()
		floatbox()
		div33_float_left()
		fmt.Print("&nbsp")
		div_off()
		div66_float_left()
		//float_right()
		system_command("host",os.Getenv("REMOTE_ADDR"))
		//		div_off()
		floatbox_off()
		//float_off()
		//		flexbox_off()
	}
	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"whois")){
		//	flexbox()
		//div33()
		float_left()
		fmt.Print("QUERY_STRING:")
		float_off()
		float_right()
		color("green")
		fmt.Print(os.Getenv("QUERY_STRING"))
		color("off")
		float_off()
		div_off()
		flexbox_off()

		//		flexbox()
		// floatbox()
		// div33_float_left()
		// fmt.Print("&nbsp")
		// div_off()
		// div66_float_left()
		//float_right()
		fmt.Println("<pre>")
		system_command("whois",os.Getenv("REMOTE_ADDR"))
		fmt.Println("</pre>")
		//		div_off()
		//		floatbox_off()
		//float_off()
		//		flexbox_off()
	}
	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"help")){
		fmt.Println("NSA // MasterTools provides:<br>")
		fmt.Println("ip.cgi?help<br>")
		fmt.Println("ip.cgi?version<br>")
		fmt.Println("ip.cgi?host<br>")
		fmt.Println("ip.cgi?whois<br>")
	}
	

	fmt.Println("</div>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
