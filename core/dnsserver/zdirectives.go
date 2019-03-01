// generated by directives_generate.go; DO NOT EDIT

package dnsserver

// Directives are registered in the order they should be
// executed.
//
// Ordering is VERY important. Every plugin will
// feel the effects of all other plugin below
// (after) them during a request, but they must not
// care what plugin above them are doing.
var Directives = []string{
	"metadata",
	"tls",
	"reload",
	"nsid",
	"root",
	"bind",
	"debug",
	"trace",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"dnstap",
	"chaos",
	"loadbalance",
	"cache",
	"rewrite",
	"dnssec",
	"autopath",
	"template",
	"hosts",
	"route53",
	"federation",
	"k8s_external",
	"rio",
	"kubernetes",
	"file",
	"auto",
	"secondary",
	"etcd",
	"loop",
	"forward",
	"proxy",
	"erratic",
	"whoami",
	"on",
}
