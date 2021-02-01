// Web service utilizing a hosted language on HTML and CSS.


package main

import(
	"fmt"
	"os"
	"log"
	"io"
	"os/exec"
	"strings"
	"encoding/json"
)

// The date is updated automatically by a user emacs function named insert-timestamp.
const (
	version = "System info:<br>ipgo.go compiled on this date:::Sun Jan 31 19:25:28 2021"

)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func log_system_init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	//	errorHandle io.Writer
	file string) {

	errorHandle, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	Trace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

type Configuration struct {
    website_directory string
}

func read_configuration_file (filename string) string {
	file, e := os.Open(filename)
	if e != nil {
		response_header()
		cat("../html/error.html")
		return ("error")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration.website_directory


}

func response_header(){
	fmt.Println("Content-type: text/html")
	fmt.Println("")
}



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
func div () {
	fmt.Println("<div>")
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

func ul () {
	fmt.Println("<ul>")
}

func ulclose () {
	fmt.Println("</ul>")
}

func li (s string) {
	fmt.Println("<li>" + s + "</li>")
}

////////////////////////////////////////////////////////////////////////

func create_menu() {
	var host string = "https://nsa.international/"
	div()
	ul()
	li("<a href=" + host + "bin/ip.cgi?host> host</a>")
	li("<a href=" + host + "bin/ip.cgi?whois> whois</a>")
	li("<a href=" + host + "bin/ip.cgi?env> env</a>")
	li("<a href=" + host + "bin/ip.cgi?version> version</a>")
	li("<a href=" + host + "bin/ip.cgi?help> help</a>")
	ulclose()
	div_off()
}

// Read in HTML files, create HTML elements. execute external programs.

func main() {

	log_system_init(os.Stdout, os.Stdout, os.Stdout, "error.log")
	if read_configuration_file("../conf/ipgo_config.json") == "error" {
		Error.Println("can't read the ../conf/ipgo_config.json file")
		goto early_exit
	}
	response_header()
	cat("../html/ip.html")
	create_menu()
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
		fmt.Println("The version function is disabled for security reasons.<br>")
		//		fmt.Println(version,"<br>")
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
	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"env")){
		fmt.Println("<pre>")
		fmt.Println("environment\n\n")
		for _, pair := range os.Environ() {
			fmt.Println(pair)}
		fmt.Println("</pre>")
	}

	if(strings.EqualFold(os.Getenv("QUERY_STRING"),"help")){
		fmt.Println("NSA // MasterTools provides:<br>")
		fmt.Println("ip.cgi?help<br>")
		fmt.Println("ip.cgi?version<br>")
		fmt.Println("ip.cgi?host<br>")
		fmt.Println("ip.cgi?whois<br>")
		fmt.Println("ip.cgi?env<br>")
	}
	

early_exit:
	fmt.Println("</div>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
