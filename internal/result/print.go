package result

import (
	"encoding/json"
	"fmt"
)

func (r *Result) PrintTargets() {
	for _, hh := range r.Hosts {
		if hh.Name != "" {
			fmt.Println(hh.Name)
			continue
		} else if len(hh.IPs) != 0 {
			// If it's just an IP, we'll only have one
			fmt.Println(hh.IPs[0])
			continue
		} else if hh.Subnet != "" {
			fmt.Println(hh.Subnet)
			continue
		} else {
			r.Logger.Debugf("unknown host: %v", hh)
		}
	}
}

func (r *Result) PrintJSON() {
	b, err := json.MarshalIndent(r.Hosts, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
