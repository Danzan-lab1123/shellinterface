package shellinterface

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

type Config struct {
        Headers map[string]string `json:"headers,omitempty"`
}

func CreateConfig() *Config {
        return &Config{
                Headers: make(map[string]string),
        }
}

type Demo struct {
        next            http.Handler
        headers         map[string]string
        name            string
        template        *template.Template
        serviceName     string
        serviceMax      string
        serviceMin      string
        requestCounterCap string
        serviceTimerCap string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
        if len(config.Headers) == 0 {
                return nil, fmt.Errorf("headers cannot be empty")
        }
        return &Demo{
                headers:        config.Headers,
                next:           next,
                name:           name,
                template:       template.New("demo").Delims("[[", "]]"),
                serviceName:    "",
                serviceMax:     "",
                serviceMin:     "",
                requestCounterCap: "",
                serviceTimerCap: "",
        }, nil
}


func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
        for key, value := range a.headers {
                if (key == "serviceName"){a.serviceName = value}
                if (key == "serviceMax"){a.serviceMax = value}
                if (key == "serviceMin"){a.serviceMin = value}
                if (key == "requestCounterCap"){a.requestCounterCap = value}
                if (key == "serviceTimerCap"){a.serviceTimerCap = value}
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
        createScaleCommand(a.serviceName, a.serviceMax, a.serviceMin, a.requestCounterCap, a.serviceTimerCap)
        a.next.ServeHTTP(rw, req)
}

func createScaleCommand(nameService string, maximalServices string, minimalServices string, requestCapServices string, timerCapServices string){
    f, err := os.Create("scaleRequest.sh")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    createCmdRequest := []string{"sh ./dockerScaler.sh",nameService,maximalServices,minimalServices,timerCapServices,requestCapServices}
    cmdRequest := strings.Join(createCmdRequest," ")
    words := []string{"#!bin/bash", cmdRequest}
    for _, word := range words {
        _, err := f.WriteString(word + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}