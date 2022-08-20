local moving = false
local started = false
local inVehicle = false
local initialCoords = nil
local active = false
local recording = false
local blip = nil
local radialBlip = nil
local checkpoint = nil
local distanceThreshold = 20

function addBlip()
	blip = AddBlipForCoord(initialCoords.x, initialCoords.y, initialCoords.z)
	SetBlipSprite(blip, 38)
	SetBlipDisplay(blip, 2)
	SetBlipScale(blip, 1.0)
	SetBlipColour(blip, 69)
	BeginTextCommandSetBlipName("STRING")
	AddTextComponentString("Complete Benchmark")
	EndTextCommandSetBlipName(blip)

	radialBlip = AddBlipForRadius(initialCoords.x, initialCoords.y, initialCoords.z, distanceThreshold + 0.0)
	SetBlipColour(radialBlip, 69)
	SetBlipAlpha(radialBlip, 128)

	checkpoint = CreateCheckpoint(
		4,
		initialCoords.x, -- pos1.x
		initialCoords.y, -- pos1.y
		initialCoords.z, -- pos1.z
		initialCoords.x, -- pos2.x
		initialCoords.y, -- pos2.y
		initialCoords.z, -- pos2.z
		(distanceThreshold * 2) + 0.0,            -- radius
		120,             -- red
		255,             -- green
		120,             -- blue
		80,              -- opacity
		0                -- reserved
	)
end

function cleanupBlip()
	if blip ~= nil then
		RemoveBlip(blip)
		blip = nil
	end
	if radialBlip ~= nil then
		RemoveBlip(radialBlip)
		radialBlip = nil
	end
	if checkpoint ~= nil then
		DeleteCheckpoint(checkpoint)
		checkpoint = nil
	end
end

function uiShow(show)
	SendNUIMessage({ type = "ui-show", show = show })

	if not show then
		cleanupBlip()
	end
end


function resetState()
	cleanupBlip()
	moving = false
	recording = false
	active = false
	initialCoords = nil
end

function checkVehicleStatus(ped)
	if IsPlayerDead(PlayerId()) then
		started = false
		uiShow(false)
		resetState()
		return
	end

	local currentlyInVehicle = IsPedInAnyVehicle(ped, false)

	if currentlyInVehicle then
		SendNUIMessage({ type = "update-speed", speed = GetEntitySpeed(GetVehiclePedIsUsing(ped)) })
	end

	if currentlyInVehicle == inVehicle then
		return
	end

	inVehicle = currentlyInVehicle
	resetState()
end

-- weather hack
Citizen.CreateThread(function()
	while true do
		SetWeatherTypeNowPersist("EXTRASUNNY")
		Citizen.Wait(10000)
	end
end)

function mainTick(ped)
	if not started then
		return
	end

	if initialCoords == nil then
		initialCoords = GetEntityCoords(ped, false)
		addBlip()
		return
	end

	local current = GetEntityCoords(ped, false)
	local distance = #(initialCoords - current)

	if distance > 0.1 and recording == false then
		recording = true
		SendNUIMessage({ type = "recording" })
		return
	end

	if not active then
		if distance > distanceThreshold then
			active = true
		end
		return
	end

	if distance > distanceThreshold then
		return
	end

	started = false
	resetState()
	SendNUIMessage({ type = "finished" })
end

Citizen.CreateThread(function()
	local ped = PlayerPedId()

	while true do
		Citizen.Wait(50)

		checkVehicleStatus(ped)
		mainTick(ped)
	end
end)

RegisterCommand('+benchmark', function()
	started = not started
	cleanupBlip()
	uiShow(started)
	resetState()
end, false)

RegisterNUICallback('setDistance', function(data, cb)
	distanceThreshold = data.distance
	cleanupBlip()
	addBlip()

	cb({ ok = true })
end)

RegisterNUICallback('saveDistance', function(_, cb)
	SetNuiFocus(false, false)
	cb({ ok = true })
end)

RegisterKeyMapping('+benchmark', 'Benchmark Vehicle', 'keyboard', 'i')

RegisterCommand('+benchmarkfocus', function()
	if recording == true then
		return
	end

	SendNUIMessage({ type = "adjust-distance", distance = distanceThreshold })
	SetNuiFocus(true, true)
end, false)

RegisterKeyMapping('+benchmarkfocus', 'Change benchmark distance', 'keyboard', 'o')