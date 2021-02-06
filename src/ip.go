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
	"bufio"
)

// The date is updated automatically by a user emacs function named insert-timestamp.
const (
        version = "System info:<br>ipgo.go compiled on this date:::Sat Feb  6 17:34:05 2021"

)

/////////////////////  Functions defining the IPGO language hosted on HTML and CSS.  ////////////////////////////////////////////////////////////

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
func div_close () {
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

////////////////////////  logging system  ////////////////////////////////////////////////////////

var (
        Trace   *log.Logger
        Info    *log.Logger
        Warning *log.Logger
        Error   *log.Logger
)

func log_system_init(info_file string, error_file string) {

        errorHandle, _ := os.OpenFile(error_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        infoHandle, _ := os.OpenFile(info_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)


        Info = log.New(infoHandle,
                "INFO: ",
                log.Ldate|log.Ltime|log.Lshortfile)


        Error = log.New(errorHandle,
                "ERROR: ",
                log.Ldate|log.Ltime|log.Lshortfile)

}


///////////////////////  Operating System interoperations  ///////////////////////////////////////////////////

func system_command(command ...string) int{
        prog,err:= exec.LookPath(command[0])
        //      fmt.Println(prog," ",command[1])
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

/////////////////////  utilities  //////////////////////////////////////////////

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


////////////////////  Configuration file operations  ////////////////////////////////////////////////////

// Read a text file into memory.
// Return a slice of the lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func get_configuration_parameter(p string) string{
	m := make(map[string]string)
	lines, err := readLines("../conf/ipgo.conf")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
       for _, line := range lines {                                                                             
                if len (line) == 0{                                                                              
                        continue}                                                                                
                if line[0] == 35{                                                                                
                        continue}                                                                                
	       s := strings.Fields(line)
		m[s[0]] = s[1]
	}
	return m[p]
}


type Configuration struct {
        Website_url string
        Website_directory string
}

func read_configuration_file (filename string, setting string) string {
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
	if setting == "website_directory" {
		return configuration.Website_directory
	}
	if setting == "website_url" {
		return configuration.Website_url
	}
	return("error")
}

/////////////////////////// html support  ///////////////////////////////////////////////////

func response_header(){
        fmt.Println("Content-type: text/html")
        fmt.Println("")
}

func create_menu(website_directory string) {
        div()
        fmt.Println ("A reflection upon your ip address.")
        ul()
        li("<a href=" + website_url + "/bin/ip.cgi?host> host</a>")
        li("<a href=" + website_url + "/bin/ip.cgi?whois> whois</a>")
        li("<a href=" + website_url + "/bin/ip.cgi?env> env</a>")
        li("<a href=" + website_url + "/bin/ip.cgi?version> version</a>")
        li("<a href=" + website_url + "/bin/ip.cgi?help> help</a>")
        ulclose()
        div_close()
}

//////////////////////////// MAIN /////////////////////////////////////////////////


// Read in HTML files, create HTML elements. execute external programs.

func main() {

        var website_url string       = get_configuration_parameter("website_url")
        var website_directory string = get_configuration_parameter("website_directory")
	//	if (website_url == "" | website_directory == "")
	
        // var website_directory string = read_configuration_file("../conf/ipgo_config.json", "website_directory")
        // var website_url string = read_configuration_file("../conf/ipgo_config.json", "website_url")

        if  website_directory != "" {
                log_system_init("info.log", "error.log")
                response_header()
                cat("../html/ip.html")
                create_menu(website_url)
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
                //      div_off()
                //      flexbox_off()
                //      fmt.Println("QUERY_STRING:", os.Getenv("QUERY_STRING"), "<br>")
                // for _, e := range os.Environ() {
                //      //        pair := strings.SplitN(e, "=", 2)
                //      //        p(pair[0],"=",pair[1])
                //      fmt.Println(e,"<br>")
                // }
                if(strings.EqualFold(os.Getenv("QUERY_STRING"),"version")){
                        fmt.Println(version,"<br>")
                }
                if(strings.EqualFold(os.Getenv("QUERY_STRING"),"host")){
                        //      flexbox()
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

                        //              flexbox()
                        floatbox()
                        div33_float_left()
                        fmt.Print("&nbsp")
                        div_off()
                        div66_float_left()
                        //float_right()
                        system_command("host",os.Getenv("REMOTE_ADDR"))
                        //              div_off()
                        floatbox_off()
                        //float_off()
                        //              flexbox_off()
                }
                if(strings.EqualFold(os.Getenv("QUERY_STRING"),"whois")){
                        //      flexbox()
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

                        //              flexbox()
                        // floatbox()
                        // div33_float_left()
                        // fmt.Print("&nbsp")
                        // div_off()
                        // div66_float_left()
                        //float_right()
                        fmt.Println("<pre>")
                        system_command("whois",os.Getenv("REMOTE_ADDR"))
                        fmt.Println("</pre>")
                        //              div_off()
                        //              floatbox_off()
                        //float_off()
                        //              flexbox_off()
                }
                if(strings.EqualFold(os.Getenv("QUERY_STRING"),"env")){
                        fmt.Println("The environment function is disabled for security reasons.<br>")
                        // fmt.Println("<pre>")
                        // fmt.Println("environment\n\n")
                        // for _, pair := range os.Environ() {
                        //      fmt.Println(pair)}
                        // fmt.Println("</pre>")
                }

                if(strings.EqualFold(os.Getenv("QUERY_STRING"),"help")){
                        fmt.Println("NSA // MasterTools provides:<br>")
                        fmt.Println("ip.cgi?help<br>")
                        fmt.Println("ip.cgi?version<br>")
                        fmt.Println("ip.cgi?host<br>")
                        fmt.Println("ip.cgi?whois<br>")
                        fmt.Println("ip.cgi?env<br>")
                }
                

                fmt.Println("</div>")
                fmt.Println("</body>")
                fmt.Println("</html>")
        }
}
