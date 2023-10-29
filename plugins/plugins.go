package main

import (
	"github.com/lukexlook/hugo-customify/plugins/functions"
	"github.com/lukexlook/hugo-customify/plugins/interfaces"
)

var f = &functions.Functions{}

func Range(cb func(m any, aliases []string, examples [][2]string)) {
	for _, reg := range []interfaces.Register{
		f.CustomLower_,
	} {
		reg(cb)
	}
}
