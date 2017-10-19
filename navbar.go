package main

import "github.com/LaughingCabbage/kevingentile.com/web"

var nav = web.Nav{
	Buttons: []web.NavButton{
		{Name: "Home",
			Link: "../index.html"},
		{Name: "GitHub",
			Link: "https://github.com/laughingcabbage/",
		},
		{Name: "Contact",
			Link:  "#contact",
			Class: "scrolly",
		},
	},
}
