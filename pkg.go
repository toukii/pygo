package main

//go:generate export PYTHONPATH=.
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sbinet/go-python"
	"github.com/toukii/goutils"
)

var (
	module *python.PyObject
)

type resp [][]float64

func init() {
	log.Printf("evn:%s\n", os.Getenv("PYTHONPATH"))
	err := python.Initialize()
	log.Printf("python.Initialize()...")
	if err != nil {
		log.Printf("python.Initialize(), err:%+v", err)
		panic(err.Error())
	}
}

func ToPyTuple(vs ...float64) *python.PyObject {
	args := python.PyTuple_New(len(vs))
	for i, v := range vs {
		python.PyTuple_SetItem(args, i, python.PyFloat_FromDouble(v))
	}
	return args
}

func ToPyListV2(input [][]float64) *python.PyObject {
	args := python.PyTuple_New(len(input))
	for i, it := range input {
		subargs := python.PyList_New(0)
		for _, jt := range it {
			python.PyList_Append(subargs, python.PyFloat_FromDouble(jt))
		}
		python.PyTuple_SetItem(args, i, subargs)
	}
	return args
}

func ToPyList(input [][]float64) *python.PyObject {
	args := python.PyList_New(0)
	for _, it := range input {
		subargs := python.PyList_New(0)
		for _, jt := range it {
			python.PyList_Append(subargs, python.PyFloat_FromDouble(jt))
		}
		python.PyList_Append(args, subargs)
	}
	return args
}

func ToPyDictV2(vs map[string]int) *python.PyObject {
	args := python.PyDict_New()
	for k, v := range vs {
		python.PyDict_SetItem(
			args,
			python.PyString_FromString(k),
			python.PyInt_FromLong(v),
		)
	}
	return args
}

func ToPyDict(vs ...float64) *python.PyObject {
	args := python.PyDict_New()
	for i, v := range vs {
		python.PyDict_SetItem(
			args,
			python.PyString_FromString(strconv.FormatInt(int64(i), 10)),
			python.PyFloat_FromDouble(v),
		)
	}
	return args
}

func ToGoSlice(out string) []string {
	s1 := strings.Split(out, "),")
	if len(s1) <= 0 {
		return nil
	}
	return strings.Split(strings.Trim(s1[0], "(("), ", ")
}

func AtoFs(strs []string) []float64 {
	ret := make([]float64, 0, len(strs))
	for _, str := range strs {
		f, err := strconv.ParseFloat(str, 10)
		if err != nil {
			log.Printf("parse %s, err:%+v", str, err)
			continue
		}
		ret = append(ret, f)
	}
	return ret
}

func Init(m string) *python.PyObject {
	log.Printf("ImportModule:%s", m)
	module = python.PyImport_ImportModule(m)
	if module == nil {
		log.Printf("could not import '%s'", m)
	}
	return module
}

func GoPyFuncV2(funcname string, args [][]float64, params map[string]int) ([][]float64, error) {
	fname := module.GetAttrString(funcname)
	if fname == nil {
		err := fmt.Errorf("could not getattr(%s, '%s')", funcname, funcname)
		log.Printf("%+v", err)
		return nil, err
	}
	log.Printf("GoPyFuncV2, %s", funcname)

	pyargs := ToPyListV2(args)
	pyparams := ToPyDictV2(params)
	log.Printf("fname:%+v", *fname)
	log.Printf("pyargs:%+v", pyargs)
	log.Printf("pyparams:%+v", pyparams)

	out := fname.Call(pyargs, pyparams)
	log.Printf("out:%+v", out)

	var r resp
	err := json.Unmarshal(goutils.ToByte(out.Bytes().String()), &r)
	if err != nil {
		log.Printf("err:%+v", err)
		return nil, err
	}
	log.Printf("resp:%+v", r)
	return [][]float64(r), nil
}

func GoPyFunc(funcname string, args ...float64) []float64 {
	fname := module.GetAttrString(funcname)
	if fname == nil {
		log.Printf("could not getattr(%s, '%s')\n", funcname, funcname)
	}
	log.Printf("GoPyFunc, %s", funcname)

	pyargs := ToPyTuple(args...)
	pyparams := ToPyDict(args...)
	log.Printf("fname:%+v", *fname)
	log.Printf("pyargs:%+v", pyargs)
	log.Printf("pyparams:%+v", pyparams)

	out := fname.Call(pyargs, pyparams)
	log.Printf("out:%+v", out)

	strs := ToGoSlice(out.Bytes().String())
	return AtoFs(strs)
}
