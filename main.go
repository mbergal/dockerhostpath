package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"os"
	"strings"
)

func main()  {
	args := os.Args[1:]

	containerPath := ""
	verbose := false
	switch len(args) {
	case 0:
		os.Exit(0);
	case 1:
		containerPath = args[0]
	case 2:
		containerPath = args[0]
		if args[1] == "-v" {
			verbose = true
		}
	}

	if ( verbose ) {
		fmt.Printf( "containerPath:  %s\n", containerPath )
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	hostname, _ := os.Hostname()
	if ( verbose ) {
		fmt.Printf( "hostname:  %s\n", hostname )
	}

	r, err := cli.ContainerInspect(ctx, hostname)
	if err != nil {
		panic(err)
	}
	for _, mount := range r.Mounts {
		if ( verbose ) {
			fmt.Printf( "mount:  %s->%s\n", mount.Source, mount.Destination )
		}
		if ( strings.HasPrefix(containerPath,mount.Destination) ) {
			fmt.Print(strings.Replace(containerPath, mount.Destination, mount.Source, 1));
			os.Exit(0)
		}
	}

	fmt.Println(containerPath)
	os.Exit(0);
}

