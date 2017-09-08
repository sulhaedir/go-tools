package main

import "go-tools/lib/request"
import "fmt"

func main() {
	var req = request.NewRequest()
	req.Method = "GET"
	req.URL = "https://www.tokopedia.com/alphamulia/nillkin-frosted-hard-case-samsung-galaxy-s8-black?trkid=f%3DCa0000L000P0W0S0Sh00Co0Po0Fr0Cb0_src%3Dpopular_page%3D1_ob%3D23_q%3Dsamsung+s8_catid%3D69_po%3D6"
	req.Headers = map[string]string{
		"user-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0",
	}
	res, body, err := req.Exec()
	if err != nil {
		fmt.Printf(err.Error())
	} else if res.StatusCode != 200 {
		fmt.Printf("%d", res.StatusCode)
	}
	fmt.Println(string(body))

}
