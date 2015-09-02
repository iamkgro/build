package main

import "fmt"
import "path/filepath"
import "os"
import "log"

func main() {

	// Get list of current directory recursively
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	directories := []string{}

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		directories = append(directories, path)
		// fmt.Printf("%s with %d bytes\n", path, f.Size())
		return nil
	})

	fmt.Print(directories)

	// Check for dependencies directory existence and build list of missing dependencies
	/*
	  lair/
	    |_deps/
	      |_caddy/
	      |_meteor/
	      |_node/
	      |_mongodb/
	      |_lair-api/
	      |_lair-app/
	    |_db/
	      |_mongodb/
	    |_lair
	*/

	// Download missing dependencies

	// Untar dependencies

	// Remove tar files

	// Start up the app

}

/*
Lair start in directory, detect if these things exist, if they dont, download
the meteor tarball, node, mongodb, api-server, caddy, and the lair app itself

Directory structure:

lair/
  |_deps/
    |_caddy/
    |_meteor/
    |_node/
    |_mongodb/
    |_lair-api/
    |_lair-app/
  |_db/
    |_mongodb/
  |_lair

Lair app will be on github or something in tarball

Lair api-server will be on github or something in tarball

Mongodb https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1404-3.0.6.tgz


lair, i dont see directories exist, download them
extract them
clean up

Ask for info, set env variables, launch it
Default config and be configurable itself

come with yaml file with some defaults

Configure mongodb for oplog, make config file for mongod

Google how to deploy with oplog parsing

If config is just strings with default config that works too, write on startup if dont exist

caddyFile := `
thign from boop {
}
`
*/
