fx_version 'cerulean'
game 'gta5'

name "aiden_react_trpc_template"
description "Displays vehicle speed, gear, and rpm."
author "Aiden"
version "0.0.1"

client_script "client/main.lua"

-- When in prod
ui_page "dist/client/index.html"

files {
	"dist/client/index.html",
	"dist/client/main.js"
}
