local moving = false
local started = false
local inVehicle = false
local initialCoords = nil
local active = false
local recording = false
local blip = nil
local radialBlip = nil

function addBlip()
	blip = AddBlipForCoord(initialCoords.x, initialCoords.y, initialCoords.z)
	radialBlip = AddBlipForRadius(initialCoords.x, initialCoords.y, initialCoords.z, 20.0)
	SetBlipSprite(blip, 38)
	SetBlipDisplay(blip, 2)
	SetBlipScale(blip, 1.0)
	SetBlipColour(blip, 69)
	SetBlipColour(radialBlip, 69)
	SetBlipAlpha(radialBlip, 128)
	BeginTextCommandSetBlipName("STRING")
	AddTextComponentString("Complete Benchmark")
	EndTextCommandSetBlipName(blip)
end


function cleanupBlip()
	if blip ~= nil then
		RemoveBlip(blip)
	end
	if radialBlip ~= nil then
		RemoveBlip(radialBlip)
	end
	blip = nil
	radialBlip = nil
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
		if distance > 20 then
			active = true
		end
		return
	end

	if distance > 20 then
		return
	end

	-- arrived
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

RegisterKeyMapping('+benchmark', 'Benchmark Vehicle', 'keyboard', 'i')
