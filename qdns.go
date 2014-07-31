// doc : http://golang.org/pkg/net/

package main

import (
	"flag"
	"fmt"
	"net"
)

var fldomain string

func init() {
	flag.StringVar(&fldomain, "name", "", "Domaine query")
}

func lip(domain string) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		//fmt.Printf("err:%s\n", err)
		return
	}

	fmt.Printf("----------------\n")
	fmt.Printf(" Information IP \n")
	fmt.Printf("----------------\n")

	for _, ip := range ips {
		fmt.Printf("\t%s\n", ip)
	}

}

func lcname(domain string) {
	cname, err := net.LookupCNAME(domain)

	if err != nil {
		//fmt.Printf("err:%s\n", err)
		return
	}

	fmt.Printf("-------------------\n")
	fmt.Printf(" Information CNAME \n")
	fmt.Printf("-------------------\n")

	fmt.Printf("\t%s\n", cname)
}

func lhosts(domain string) {
	hosts, err := net.LookupHost(domain)

	if err != nil {
		//fmt.Printf("err:%s\n", err)
		return
	}

	fmt.Printf("------------------\n")
	fmt.Printf(" Information HOST \n")
	fmt.Printf("------------------\n")

	for _, host := range hosts {
		fmt.Printf("\t%s\n", host)
	}
}

func lmx(domain string) {
	mxs, err := net.LookupMX(domain)

	if err != nil {
		//fmt.Printf("err:%s\n", err)
		return
	}

	fmt.Printf("-----------------------------------\n")
	fmt.Printf(" Information MX (Serveurs de mail) \n")
	fmt.Printf("-----------------------------------\n")

	for _, mx := range mxs {
		fmt.Printf("\t%d - %s\n", mx.Pref, mx.Host)
	}
}

func ltxt(domain string) {
	txts, err := net.LookupTXT(domain)

	if err != nil {
		//fmt.Printf("err:%s\n", err)
		return
	}

	fmt.Printf("-----------------\n")
	fmt.Printf(" Information TXT \n")
	fmt.Printf("-----------------\n")

	for _, txt := range txts {
		fmt.Printf("\t%s\n", txt)
	}

}

func main() {

	flag.Parse()

	if len(fldomain) == 0 {
		flag.Usage()
	}

	lip(fldomain)
	lmx(fldomain)
	lcname(fldomain)
	lhosts(fldomain)
	ltxt(fldomain)

}
