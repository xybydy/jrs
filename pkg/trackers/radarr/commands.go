package radarr

import (
	"fmt"
	"io/ioutil"
	"log"
)

func (r *Radarr) CmdGetStatus() {
	req, _ := r.GetOngoingCommands()
	resp, _ := r.Client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)

}
