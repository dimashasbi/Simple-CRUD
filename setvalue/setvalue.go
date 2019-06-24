package setvalue

import (
	"M-GateDBConfig/model"
	"encoding/json"
	"fmt"
)

// SetValue to input value to DB
func SetValue() {
	backsetting := &model.BackSettingsConfig{
		URLJavaMPAY: "localhost:27000",
		URLJavaSVA:  "localhost:27001",
	}
	res1B, _ := json.Marshal(backsetting)
	fmt.Println(string(res1B))
}
