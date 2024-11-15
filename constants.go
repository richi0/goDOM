package goDOM

var acceptedAttributes = []string{
	"accept",
	"accept-charset",
	"accesskey",
	"action",
	"alt",
	"async",
	"autocomplete",
	"autofocus",
	"autoplay",
	"charset",
	"checked",
	"cite",
	"content",
	"contenteditable",
	"controls",
	"coords",
	"data",
	"data-*",
	"datetime",
	"default",
	"defer",
	"dir",
	"dirname",
	"disabled",
	"download",
	"draggable",
	"enctype",
	"enterkeyhint",
	"for",
	"form",
	"formaction",
	"headers",
	"href",
	"hreflang",
	"http-equiv",
	"inert",
	"inputmode",
	"ismap",
	"kind",
	"label",
	"lang",
	"list",
	"loop",
	"low",
	"max",
	"maxlength",
	"media",
	"method",
	"min",
	"multiple",
	"muted",
	"name",
	"novalidate",
	"onabort",
	"onafterprint",
	"onbeforeprint",
	"onbeforeunload",
	"onblur",
	"oncanplay",
	"oncanplaythrough",
	"onchange",
	"onclick",
	"oncontextmenu",
	"oncopy",
	"oncuechange",
	"oncut",
	"ondblclick",
	"ondrag",
	"ondragend",
	"ondragenter",
	"ondragleave",
	"ondragover",
	"ondragstart",
	"ondrop",
	"ondurationchange",
	"onemptied",
	"onended",
	"onerror",
	"onfocus",
	"onhashchange",
	"oninput",
	"oninvalid",
	"onkeydown",
	"onkeypress",
	"onkeyup",
	"onload",
	"onloadeddata",
	"onloadedmetadata",
	"onloadstart",
	"onmousedown",
	"onmousemove",
	"onmouseout",
	"onmouseover",
	"onmouseup",
	"onmousewheel",
	"onoffline",
	"ononline",
	"onpagehide",
	"onpageshow",
	"onpaste",
	"onpause",
	"onplay",
	"onplaying",
	"onpopstate",
	"onprogress",
	"onratechange",
	"onreset",
	"onresize",
	"onscroll",
	"onsearch",
	"onseeked",
	"onseeking",
	"onselect",
	"onstalled",
	"onstorage",
	"onsubmit",
	"onsuspend",
	"ontimeupdate",
	"ontoggle",
	"onunload",
	"onvolumechange",
	"onwaiting",
	"onwheel",
	"open",
	"optimum",
	"pattern",
	"placeholder",
	"popover",
	"popovertarget",
	"popovertargetaction",
	"poster",
	"preload",
	"readonly",
	"rel",
	"required",
	"reversed",
	"sandbox",
	"scope",
	"selected",
	"shape",
	"span",
	"spellcheck",
	"src",
	"srcdoc",
	"srclang",
	"srcset",
	"start",
	"step",
	"tabindex",
	"target",
	"title",
	"translate",
	"type",
	"usemap",
	"value",
}