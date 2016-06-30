package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/eandre/lunar"
	"golang.org/x/tools/go/loader"
)

type Lib struct {
	Path   string   `json:"path"`
	TOC    []string `json:"toc"`
	Embeds []string `json:"embeds"`
}

type Addon struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	SavedVariables []string `json:"saved_variables"`
	Assets         string   `json:"assets"`
	ImportPath     string   `json:"import_path"`
	Libs           []Lib    `json:"libs"`
}

func main() {
	output := flag.String("output", "./output", "Output directory")
	flag.Parse()

	addon, err := parseAddon(flag.Args())
	if err != nil {
		log.Fatalln("Could not parse addon:", err)
	}

	var conf loader.Config
	_, err = conf.FromArgs([]string{addon.ImportPath}, false)
	if err != nil {
		log.Fatalln("Could not load packages:", err)
	}

	prog, err := conf.Load()
	if err != nil {
		log.Fatalln("Could not parse packages:", err)
	}

	if err := os.MkdirAll(*output, 0755); err != nil {
		log.Fatalln("Could not create directory:", err)
	}

	parser := lunar.NewParser(prog)
	pkgName := filepath.Base(*output)
	var filenames []string

	var infos []*loader.PackageInfo
	for _, i := range prog.AllPackages {
		infos = append(infos, i)
	}
	pkgs := sortPkgs(infos)
	for _, pkg := range pkgs {
		if parser.IsTransientPkg(pkg.Pkg) {
			continue
		}

		path := pkg.Pkg.Path()
		pkgPath := filepath.Join(*output, path)
		if err := os.MkdirAll(pkgPath, 0755); err != nil {
			log.Fatalln("Could not create directory:", err)
		}

		for _, f := range pkg.Files {
			basePath := strings.TrimSuffix(prog.Fset.File(f.Pos()).Name(), ".go")
			fname := filepath.Base(basePath) + ".lua"
			filenames = append(filenames, filepath.Join(path, fname))

			out, err := os.Create(filepath.Join(pkgPath, fname))
			if err != nil {
				log.Fatalln("Could not create file:", err)
			}
			defer out.Close()

			// Check if a lua file already exists
			fi, err := os.Stat(basePath + ".lua")
			if err == nil && !fi.IsDir() {
				bytes, err := ioutil.ReadFile(basePath + ".lua")
				if err != nil {
					log.Fatalf("Could not read file %s: %v", fname, err)
				}
				if _, err := out.Write(bytes); err != nil {
					log.Fatalf("Could not write file %s: %v", fname, err)
				}
				continue
			}

			if err := parser.ParseNode(out, f); err != nil {
				log.Fatalf("Could not parse file %s: %v", fname, err)
			}
		}
	}

	// Copy libs
	if len(addon.Libs) > 0 {
		tocPaths, err := CopyLibs(addon.Libs, *output)
		if err != nil {
			log.Fatalf("Could not copy lualibs: %v", err)
		}
		filenames = append(tocPaths, filenames...)
	}

	// Copy assets
	if addon.Assets != "" {
		if _, err := CopyDir(addon.Assets, filepath.Join(*output, "assets"), nil); err != nil {
			log.Fatalf("Could not copy assets: %v", err)
		}
	}

	if err := writePackage(prog, addon, *output, pkgName, filenames); err != nil {
		log.Fatalf("Could not write package: %v", err)
	}
}

func writePackage(prog *loader.Program, addon *Addon, root, pkgName string, filenames []string) error {
	prelude, err := os.Create(filepath.Join(root, "_prelude.lua"))
	if err != nil {
		return err
	}
	defer prelude.Close()
	if _, err := lunar.WriteBuiltins(prelude); err != nil {
		return err
	}

	postlude, err := os.Create(filepath.Join(root, "_postlude.lua"))
	if err != nil {
		return err
	}
	if _, err := postlude.Write([]byte("lunar_go_builtins.run_inits()")); err != nil {
		return err
	}

	toc, err := os.Create(filepath.Join(root, pkgName+".toc"))
	if err != nil {
		return err
	}
	defer toc.Close()
	var tw tocWriter
	tw.AddKey("Interface", "60200")
	tw.AddKey("Title", addon.Title)
	tw.AddKey("SavedVariables", strings.Join(addon.SavedVariables, ", "))
	tw.AddFile("_prelude.lua")
	for _, fn := range filenames {
		tw.AddFile(fn)
	}
	tw.AddFile("_postlude.lua")
	return tw.Output(toc)
}

func parseAddon(args []string) (*Addon, error) {
	if len(args) == 0 {
		return nil, errors.New("no addons.json path specified")
	}

	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		return nil, err
	}
	var addon Addon
	err = json.Unmarshal(data, &addon)
	return &addon, err
}
