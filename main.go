// A simple Quake 2 master server implementation.
//
// Quake 2 game servers with the `public` cvar set to 1 will report to any
// configured master servers. This allows players to find servers and see
// basic data about each server (current map, some settings, players, scores
// pings, etc). Game servers can specify a particular master server with a
// command like: `setmaster <x.x.x.x[:yyyyy]>`. Port UDP/27000 is the default.
//
// This implemtation includes an optional HTTP server to handle API requests
// for querying the current state/data. Enable with the `-api` flag.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/packetflinger/libq2/master"
	"google.golang.org/protobuf/encoding/prototext"

	pb "github.com/packetflinger/q2master/proto"
)

var (
	listenPort = flag.Int("port", 27900, "Port to listen on")
	listenIP   = flag.String("addr", "[::]", "IP address to listen on")
	foreground = flag.Bool("fg", false, "Log to stdout instead of file")
	logfile    = flag.String("logfile", "master.log", "The filename to use for the log")
	api        = flag.Bool("api", false, "Whether or not to enable the web API")
	apiPort    = flag.Int("apiport", 3333, "TCP port for web requests")
	apiIP      = flag.String("apiaddr", "[::]", "The IP address to listen on for web requests")
	config     = flag.String("config", ".config/q2master/config", "")
)

func main() {
	flag.Parse()
	c, err := loadConfig(*config)
	applyConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	if !*foreground {
		fp, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
		log.SetOutput(io.Writer(fp))
	}

	log.Println("*** Quake 2 Master Server")
	log.Printf("*** (c) 2022-%d Packetflinger Industries\n", time.Now().Year())
	log.Println("*** https://github.com/packetflinger/q2master")

	m := master.NewMaster()
	m.Address = *listenIP
	m.Port = *listenPort
	m.ApiEnabled = *api
	m.ApiIP = *apiIP
	m.ApiPort = *apiPort
	m.Run()
}

func loadConfig(f string) (*pb.MasterConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("loadConfig userhomedir: %v", err)
	}
	data, err := os.ReadFile(path.Join(home, f))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, fmt.Errorf("loadConfig readfile: %v", err)
		}
	}
	var config pb.MasterConfig
	err = prototext.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("loadConfig unmarshal: %v", err)
	}
	return &config, nil
}

func applyConfig(c *pb.MasterConfig) {
	if c == nil {
		return
	}
	if c.GetPort() != 0 {
		*listenPort = int(c.GetPort())
	}
	if c.GetAddr() != "" {
		*listenIP = c.GetAddr()
	}
	if c.GetForeground() {
		*foreground = c.GetForeground()
	}
	if c.GetLogfile() != "" {
		*logfile = c.GetLogfile()
	}
	if c.GetApi() {
		*api = c.GetApi()
	}
	if c.GetApiPort() != 0 {
		*apiPort = int(c.GetApiPort())
	}
	if c.GetApiAddr() != "" {
		*apiIP = c.GetApiAddr()
	}
}
