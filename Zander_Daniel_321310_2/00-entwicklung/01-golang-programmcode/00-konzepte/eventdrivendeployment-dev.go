package eventdrivendeployment

import(
	"context"
	"fmt"
	"net"
	"sync"
   	"log"
   	"os"
   	"os/exec"
   	"strings"
   	"strconv"
	
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
    	"github.com/docker/docker/client"

	"github.com/traefik/traefik/v3/pkg/middlewares"
	"github.com/traefik/traefik/v3/pkg/tcp"
)


const dockerClient := context.Background()
const dockerContext err := client.NewEnvClient()
        if err != nil {
                panic(err)
                fmt.Println("No Client available")
        }
const defaultService = "w_mondial"

var globalservicename string
var globalcounter int
globalcounter = 0

type Config struct{
	serviceName string
}

func CreateConfig() *Config{
	return &Config{
	serviceName: defaultService,}
}

type eventDrivenDeployment struct {
	name           string
	next           tcp.Handler
	serviceName    string
}

func New(ctx context.Context, next tcp.Handler, config *Config, name string) (tcp.Handler, error) {
	logger := middlewares.GetLogger(ctx, name, typeName)
	logger.Debug().Msg("Creating middleware")

	return &eventDrivenDeployment{
		name:           name,
		next:           next,
		serviceName:    config.serviceName,

	}, nil
}


func (edd *eventDrivenDeployment) ServeTCP(conn tcp.WriteCloser) {
	logger := middlewares.GetLogger(context.Background(), edd.name, edd.serviceName)

	ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		logger.Error().Err(err).Msg("Cannot parse IP from remote addr")
		conn.Close()
		return
	}

	serviceID := edd.serviceByName(edd.serviceName)
	globalservicename = edd.serviceName

	edd.increment(serviceID)
	defer edd.decrement(serviceID)
	edd.next.ServeTCP(conn)
}


//daniel:erstelle container
func (edd *eventDrivenDeployment) increment(serviceID string) {
	edd.mu.Lock()
	defer i.mu.Unlock()
	globalcounter += 1

	edd.createScaleCommand()
	edd.executeScaleCommand()

	return nil
}

//daniel:lösche container
func (edd *eventDrivenDeployment) decrement(ip string) {
	edd.mu.Lock()
	defer edd.mu.Unlock()
        globalcounter -= 1

	edd.createScaleCommand()
	edd.executeScaleCommand()

}

//daniel:finde passende id zum servicenamen
func (edd *eventDrivenDeployment) serviceByName(serviceName string) {
	var serviceReturnID string
	services, err := cli.ServiceList(ctx, types.ServiceListOptions{})
        if err != nil {
                panic(err)
                fmt.Println("No services running")
        }

      	for _, service := range services {
       	if service.Spec.Name == "w_mondial"{ serviceReturnID := service.ID}
       	}

	return serviceReturnID 
}

//daniel:erstelle ein bashskript zum ausführen eines docker service scale up befehls
func (edd *eventDrivenDeployment) createScaleCommand(){
var servicename string
  f, err := os.Create("dockerscale.sh")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    servicename = globalservicename
    containercounter := strconv.Itoa(globalcounter)
    createArgumentForParameterThree := []string{"docker service scale ",servicename,"=",containecounter}  
    ArgumentForParameterThree := strings.Join(createArgumentForParameterThree,"")

    words := []string{"#!bin/bash", ArgumentForParameterThree}
    for _, word := range words {
        _, err := f.WriteString(word + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}

//daniel:führe den erstellten Bashscript aus und lösche ihn darauf hin
func (edd *eventDrivenDeployment) executeScaleCommand(){

//daniel:gebe skript berechtigung
        chmod := exec.Command("chmod", "u+x", "./dockerscale.sh")
        cerr := cmd.Run()
        if cerr != nill {log.Fatalf("no permission given", cerr)}
		
//daniel:führe skript aus
	sh := exec.Command("sh", "./dockerscale.sh")
	sherr := sh.Run()
	if sherr != nil {log.Fatalf("no script executed", sherr)}

//daniel:entferne skript
	rm := exec.Command("rm", "./dockerscale.sh")
	rmerr := rm.Run()
	if rmerr != nil {log.Fatalf("no skript deleted", rmerr)}

	fmt.Println("service paramter changed")
}

