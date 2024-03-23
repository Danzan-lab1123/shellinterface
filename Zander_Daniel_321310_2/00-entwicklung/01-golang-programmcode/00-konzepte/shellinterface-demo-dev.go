//daniel: konzept der zu gestalltenden middleware für den treafikserver
//daniel:die middleware die als plugin ummgestetz wird legt ausführbare sh-dateien im treafikserver ab die regelmaeßig aufgerufenen werden
package bachelor-thesis-321310

//daniel: import von der "os" bibliothek, dass einen dateiablage möglich ist
import (
        "log"
        "context"
        "os"
        "strings"
        "strconv"
        "net"
        "net/http"
        "text/template"
        "bytes"
)

//daniel: erstellung von standartwerten
const(
        serviceName = stack_anwendung
        serviceCounter= 0
)

// daniel: übernimmt konfiguationen der docker-compose.yml 
type Config struct {
        [...]
        serviceName string
        serviceCounter int
}

// daniel: weist bei keiner übergabe von argumenten innerhalb der konfigurationen den variable standartwerte zu 
func CreateConfig() *Config {
        return &Config{
                [...]
                serviceName: defaultService,
                serviceCounter: defaultCounter,
        }
}

// daniel: teilauschnitt der spezifikation des shellinterface "objectes" 
type shellinterface struct {
        [...]
        serviceName    string
        serviceCounter int
}

// daniel: teilauschnitt der instanziierung des shellinterface "objectes"
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
        if len(config.Headers) == 0 {
                return nil, fmt.Errorf("headers cannot be empty")
        }
        return &Demo{
                [...]
                serviceName:    config.serviceName,
                serviceCounter: config.serviceCounter,
        }, nil
}

//daniel: beantworte eine anfrage and nutze die methode createScaleCommand()
func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
                [...]
                a.createScaleCommand(a.serviceName)
                a.serviceCounter++
       }
        a.next.ServeHTTP(rw, req)
}

//daniel: wird diese methode aufgrufen werden die übergebenen konfigurationen genutzt um einen auf den services angepasste bash-datei im traefik-server abzulegen
func (a *Demo) createScaleCommand(serviceName string){
//daniel: erstelle datei "dockerscaler.sh"
    f, err := os.Create("dockerscaler.sh")
    if err != nil {
        log.Fatal(err)
    }
//daniel:verschieb die schließung um in die datei zu schreiben
    defer f.Close()
        
//daniel: entnehme die countervar und convertieren diese in einen string
    var countername string
    counterstring = strconv.Itoa(a.serviceCounter)
        
//daniel: schreibe den docker befehlt docker scale ${servicename}=${servicecounter} in die datei
    createCommand := []string{"docker service scale ",a.serviceName,"=",counterstring}
    command := strings.Join(createCommand,"")
    words := []string{"#!bin/bash", command}
    for _, word := range words {
        _, err := f.WriteString(word + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}
