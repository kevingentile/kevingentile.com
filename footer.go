package main

import "github.com/LaughingCabbage/kevingentile.com/web"

//type ContactButton struct {
//	Label string
//	Link  string
//	Icon  string
//}
var footer = web.ContactBar{
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
