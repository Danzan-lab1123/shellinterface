//daniel:prototyp der middleware die als plugin für das ablegen von skripten innerhalb des treafikserver
package shellinterface

import (
        "fmt"
        "log"
        "context"
        "os"
        "strings"
        "net/http"
        "text/template"
        "bytes"
)

// daniel: übernahme der werte des docker-compose.yml
type Config struct {
        Headers map[string]string `json:"headers,omitempty"`
}

// daniel: weist bei keiner übergabe von argumenten innerhalb der konfigurationen den variable standartwerte zu
func CreateConfig() *Config {
        return &Config{
                Headers: make(map[string]string),
        }
}

// daniel: spezifikation des shellinterface "objectes" 
type Shellinterface struct {
        next     	http.Handler
        headers  	map[string]string
        name    	string
        template	*template.Template
        serviceName string
		serviceMax	string
		serviceMin	string
}

// daniel: instanziierung des shellinterface "objectes"
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
        if len(config.Headers) == 0 {
                return nil, fmt.Errorf("headers cannot be empty")
        }
        return &Shellinterface{
                headers:  	config.Headers,
                next:     	next,
                name:     	name,
                template: 	template.New("shellinterface").Delims("[[", "]]"),
                serviceName:    "",
		serviceMax:	"",
		serviceMin:	"",
        }, nil
}

//daniel: beantworte eine anfrage and nutze die methode createScaleCommand()
func (a *Shellinterface) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
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
		createScaleCommand(a.serviceName, a.serviceMax, a.serviceMin)
        a.next.ServeHTTP(rw, req)
}

//daniel: wird diese methode aufgrufen werden die übergebenen konfigurationen genutzt um einen auf den services angepasste bash-datei im traefik-server abzulegen
func createScaleCommand(nameService string, maximalServices string, minimalServices string){
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
