/*--------------------------------------------------------------*/
/*							Main Module							*/
/*							Jake Mussler						*/
/*							20 Nov 2021							*/
/*							REV1_00								*/
/*--------------------------------------------------------------*/

package file

import (
	"encoding/json"
	"fmt"
	"os"

	webserv "github.com/Urvirith/go_fractal/src/server"
)

/* TO DO - ADD WEBSITE AND CERTIFICATE LOCATION COMPARED TO CALL SOURCE */
type Config struct {
	Webserv webserv.WebData
}

func Setup() {
	var config Config

	fmt.Println("Configuration File Setup")
	fmt.Println("Please Enter Server Name:")
	fmt.Scanln(&config.Webserv.ServerName)

	js, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		fmt.Println("Failed To Serialize Config File")
		return
	}

	err = os.WriteFile("./config.json", js, 0644)
	if err != nil {
		fmt.Println("Failed To Write Config File")
		return
	}
}

func Confget() (Config, error) {
	var config Config
	bytes, err := os.ReadFile("./config.json")

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
