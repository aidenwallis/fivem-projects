fx_version 'cerulean'
game 'gta5'

name "aiden_vehicle_hud"
description "Displays vehicle speed, gear, and rpm."
author "Aiden"
version "0.0.1"

client_scripts {
	'client/main.lua',
    -- enable this if you want chaos
    -- 'client/fun.lua'
}

ui_page "dist/index.html"

files {
	"dist/index.html",
	"dist/index.js"
}
