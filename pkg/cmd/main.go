package cmd

import (
	"fmt"
	"github.com/facebookgo/symwalk"
	"github.com/fkautz/gitbom-go"
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmd"
	"github.com/rwxrob/cmdbox/util"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var Cmd = &bonzai.Cmd{
	Name:      `gitbom`,
	Summary:   `gitbom`,
	Usage:     `[gitbom|h|help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 gitbom-go contributors`,
	License:   `Apache-2`,
	Commands:  []*bonzai.Cmd{cmd.Help, artifactTree},

	Description: `
		The foo commands do foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Otherwise, just
		follow the same guidelines as for Go documentation. Note that the
		x.Call Method here is omitted since the main work is delegated to
		the subcommands in the command tree. The help command, however, is
		the default because it is first. `,

	// no Call since has Commands, if had Call would only call if
	// commands didn't match

	Call: func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(util.Emph("**NAME**", 0, -1) + `
       gitbom (v0.0.1) - Generate gitboms from files

` + util.Emph("**USAGE**", 0, 01) + `
       gitbom [files]
       gitbom [file] bom [input-files]

       gitbom will create a .bom/ directory in the current working
       directory and store generated gitboms in .bom/

` + util.Emph("**LEGAL**", 0, 01) + `
       gitbom (v0.0.1) Copyright 2022 gitbom-go contributors
       SPDX-License-Identifier: Apache-2.0
`)
			return nil
		}

		if len(args) > 2 && args[1] == "bom" {
			gb := gitbom.NewGitBom()
			// generate artifact tree
			for i := 2; i < len(args); i++ {
				if err := addPathToGitbom(gb, args[i]); err != nil {
					return err
				}
			}

			// generate target gitbom with artifact tree
			if err := writeObject(".bom", gb); err != nil {
				return err
			}

			gb2 := gitbom.NewGitBom()
			info, err := os.Stat(args[0])
			if err != nil {
				return err
			}
			if err = addFileToGitbom(args[0], info, gb2, gb); err != nil {
				return err
			}

			if err := writeObject(".bom", gb2); err != nil {
				return err
			}

			fmt.Println(gb2.Identity())
		} else {
			gb := gitbom.NewGitBom()
			for i := 0; i < len(args); i++ {
				if err := addPathToGitbom(gb, args[i]); err != nil {
					log.Println(err)
					return err
				}
			}

			// generate target gitbom with artifact tree
			if err := writeObject(".bom", gb); err != nil {
				log.Println(err)
				return err
			}

			fmt.Println(gb.Identity())
		}

		return nil
	},
}

var artifactTree = &bonzai.Cmd{
	Name: "artifact-tree",
	Call: nil,
}

func writeObject(prefix string, gb gitbom.ArtifactTree) error {
	objectDir := path.Join(prefix, "object", gb.Identity()[:2])
	objectPath := path.Join(objectDir, gb.Identity()[2:])
	if err := os.MkdirAll(objectDir, 0755); err != nil {
		log.Println(err)
		return err
	}
	if err := ioutil.WriteFile(objectPath, []byte(gb.String()), 0644); err != nil {
		return err
	}
	return nil
}

func addPathToGitbom(gb gitbom.ArtifactTree, fileName string) error {
	err := symwalk.Walk(fileName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			err2 := addFileToGitbom(path, info, gb, nil)
			if err2 != nil {
				return err2
			}
		}
		return nil
	})
	return err
}

func addFileToGitbom(path string, info os.FileInfo, gb gitbom.ArtifactTree, identifier gitbom.Identifier) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("error closing %s: %s", path, err)
		}
	}(f)

	if err := gb.AddSha1ReferenceFromReader(f, identifier, info.Size()); err != nil {
		return err
	}
	return nil
}
