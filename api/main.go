package main

import (
	"embed"
	"fmt"
	"io/fs"
	"regexp"
	"strings"

	_ "github.com/joho/godotenv/autoload" // auto load env
)

//go:embed all:dist
var webBuild embed.FS

func main() {
	// db, _ := database.NewDB("test.db")

	// for _, instance := range db.GetAutoStartInstances() {
	// 	go instance.Run()
	// }

	// ser := server.New(db)

	// /* init frontend -- start */
	dir, err := fs.Sub(webBuild, "dist")
	if err != nil {
		panic(err)
	}

	// ser.AddFileHandler(dir)
	// /* init frontend -- end */

	// ser.Start()
	fmt.Println(getDynamicRoutes(dir).HasIs("instances/awa/1"))
}

var checkDynamicRoute = regexp.MustCompile(`/\[[^/]*\]`)

type T map[string]T

// get dynamic routes from embed files
// check path is match `/\[[^/]*\]` ( for next.js export path format )
func getDynamicRoutes(dir fs.FS) T {
	dPaths := &T{}

	fs.WalkDir(dir, ".", func(path string, file fs.DirEntry, _ error) (err error) {
		if file.IsDir() {
			if strings.HasPrefix(path, "/") {
				path = "/" + path
			}
			if checkDynamicRoute.MatchString(path) {
				var t T
				for i, p := range strings.Split(path, "/") {
					if i == 0 {
						(*dPaths)[p] = T{}
						t = (*dPaths)[p]
					} else {
						t[p] = T{}
					}
				}
			}
		}

		return // return nil
	})

	return *dPaths
}

func (s T) HasIs(path string) bool {
	var t T

	paths := strings.Split(path, "/")
	for i, p := range paths {
		if i == 0 {
			if _, ok := s[p]; ok {
				t = s[p]
				continue
			} else {
				break
			}
		}
		if _, ok := t[p]; ok {
			t = t[p]

			if i == len(paths)-1 {
				return true
			}
		}
		for key := range t {
			fmt.Println(key)
			if strings.HasPrefix(key, "[") && strings.HasSuffix(key, "]") {
				if i == len(paths)-1 {
					return true
				}
				continue
			}
			return false
		}
	}

	return false
}
