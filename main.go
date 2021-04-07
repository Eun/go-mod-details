package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	act "github.com/sethvargo/go-githubactions"
	"golang.org/x/mod/modfile"
)

func main() {
	file := act.GetInput("modfile")
	if file == "" {
		file = "go.mod"
	}
	f, err := parse(file)
	if err != nil {
		log.Fatalf("unable to parse %q: %s", file, err)
	}
	act.SetOutput("modfile", file)
	act.SetOutput("go_version", f.Go.Version)
	act.SetOutput("module", f.Module.Mod.Path)

	// ignore the Syntax field for following types
	for _, r := range []reflect.Type{
		reflect.TypeOf(modfile.Require{}),
		reflect.TypeOf(modfile.Exclude{}),
		reflect.TypeOf(modfile.Replace{}),
		reflect.TypeOf(modfile.Retract{}),
	} {
		jsoniter.RegisterFieldEncoder(r.String(), "Syntax", nilEncoder{})
	}

	s, err := encodeToJSON(f.Require)
	if err != nil {
		log.Fatalf("unable to encode require: %s", err)
	}
	act.SetOutput("require_json", s)

	s, err = encodeToJSON(f.Exclude)
	if err != nil {
		log.Fatalf("unable to encode exclude: %s", err)
	}
	act.SetOutput("exclude_json", s)

	s, err = encodeToJSON(f.Replace)
	if err != nil {
		log.Fatalf("unable to encode replace: %s", err)
	}
	act.SetOutput("replace_json", s)

	s, err = encodeToJSON(f.Retract)
	if err != nil {
		log.Fatalf("unable to encode retract: %s", err)
	}
	act.SetOutput("retract_json", s)
}

type nilEncoder struct{}

func (n nilEncoder) IsEmpty(ptr unsafe.Pointer) bool { return true }

func (n nilEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {}

func (n nilEncoder) IsEmbeddedPtrNil(ptr unsafe.Pointer) bool { return true }

func encodeToJSON(v interface{}) (string, error) {
	return jsoniter.ConfigFastest.MarshalToString(v)
}

func parse(file string) (*modfile.File, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return modfile.ParseLax(file, buf, nil)
}
