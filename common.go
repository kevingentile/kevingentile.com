package main

import "github.com/LaughingCabbage/kevingentile.com/web"

//type NavButton struct {
//	Name  string
//	Link  string
//	Class string
//}
var Nav = web.Nav{
	Buttons: []web.NavButton{
		{Name: "Home",
			Link: "../index.html"},
		{Name: "GitHub",
			Link: "https://github.com/laughingcabbage/",
		},
		{Name: "Contact",
			Link: "../contact.html",
		},
	},
}

//type ContactButton struct {
//	Label string
//	Link  string
//	Icon  string
//}
var Footer = web.ContactBar{
	Header: "Want to stay in touch?",
	Body:   "Connect on social media!",
	Contacts: []web.ContactButton{
		{Label: "Email",
			Icon: "fa-envelope",
			Link: "mailto:gentile_kevin94@hotmail.com",
		},
		{Label: "GitHub",
			Icon: "fa-github",
			Link: "https://www.github.com/laughingcabbage",
		},
		{Label: "Twitter",
			Icon: "fa-twitter",
			Link: "https://twitter.com/kevin_gentile",
		},
		{Label: "LinkedIn",
			Icon: "fa-linkedin",
			Link: "https://www.linkedin.com/in/gentilekevin/",
		},
	},
}

var ScriptsCommon = []web.Script{
	"../assets/js/jquery.min.js",
	"../assets/js/jquery.dropotron.min.js",
	"../assets/js/jquery.scrolly.min.js",
	"../assets/js/jquery.onvisible.min.js",
	"../assets/js/skel.min.js",
	"../assets/js/util.js",
	"../assets/js/main.js",
}
