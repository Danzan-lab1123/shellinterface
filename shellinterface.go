package remoteplug

import (
        "fmt"
        "log"
        "context"
        "os"
        "strings"
        "strconv"
        "net"
        "net/http"
        "text/template"
        "bytes"
        "sync"
)

// Config the plugin configuration.
type Config struct {
        Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
        return &Config{
                Headers: make(map[string]string),
        }
}

// Demo a Demo plugin.
type Demo struct {
        next     	http.Handler
        headers  	map[string]string
        name    	string
        template	*template.Template
        serviceName    	string
	serviceMax	string
	serviceMin	string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
        if len(config.Headers) == 0 {
                return nil, fmt.Errorf("headers cannot be empty")
        }
        return &Demo{
                headers:  	config.Headers,
                next:     	next,
                name:     	name,
                template: 	template.New("demo").Delims("[[", "]]"),
                serviceName:    "",
		serviceMax:	"",
		serviceMin:	"",
        }, nil
}


func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
        for key, value := range a.headers {
		if (key == "serviceName"){a.serviceName = value}
		if (key == "serviceMax"){a.serviceMax = value}
		if (key == "serviceMin"){a.serviceMin = value}
                tmpl, err := a.template.Parse(value)
                if err != nil {
                        http.Error(rw, err.Error(), http.StatusInternalServerError)
                        return
                }

                writer := &bytes.Buffer{}
                err = tmpl.Execute(writer, req)
                if err != nil {
                        http.Error(rw, err.Error(), http.StatusInternalServerError)
                        return
                }
                req.Header.Set(key, writer.String())
        }
	a.createScaleCommand(a.serviceName, a.serviceMax, a.serviceMin)
        a.next.ServeHTTP(rw, req)
}

func (a *Demo) createScaleCommand(nameService string, maximalServices string, minimalServices string){
    f, err := os.Create("scaleRequest.sh")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    createCmdRequest := []string{"sh ./dockerScaler.sh",nameService,maximalServices,minimalServices}
    cmdRequest := strings.Join(createCmdRequest," ")
    words := []string{"#!bin/bash", cmdRequest}
    for _, word := range words {
        _, err := f.WriteString(word + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}
