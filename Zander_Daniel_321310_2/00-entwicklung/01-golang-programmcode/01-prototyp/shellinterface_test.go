//daniel: test der erweiterung createScaleCommand(), des offiziellen plugin demo template von treafik
package bachelor-thesis-321310

//daniel: importierung der test klasse
import (
	"testing"
)


//daniel: teste createScaleCommand mit selbst bestimmtenen werte und prüfe auf funktionsfähigkeit
func TestDemo_createScaleCommand(t *testing.T) {
	type args struct {
		nameService     string
		maximalServices string
		minimalServices string
	}
	
//daniel: bestimmung der einzusetzenden werte
	tests := []struct {
		name   string
		args   args
	}{	
	// TODO: Add test cases.
	{
		name:	"test",
		args: args{
		nameService:     "stackname_servicename",
		maximalServices: "10",
		minimalServices: "5",
		},
	
	},
}

//daniel: test and auswertung der methode
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		createScaleCommand(tt.args.nameService, tt.args.maximalServices, tt.args.minimalServices)
		})
	}
}