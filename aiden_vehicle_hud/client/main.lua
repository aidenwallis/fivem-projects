local isInVehicle = false
local currentHealth = 100
local currentGear = -1
local currentSpeed = 0
local currentRPM = 0
local currentGravity = maxGravity

Citizen.CreateThread(function()
	while true do
		Citizen.Wait(0)

		local ped = PlayerPedId()

		if not isInVehicle and not IsPlayerDead(PlayerId()) then
			if IsPedInAnyVehicle(ped, false) then
				-- suddenly appeared in a vehicle, possible teleport
				isInVehicle = true
				SendNUIMessage({
					type = "entered-vehicle"
				})
			end
		elseif isInVehicle then
			if not IsPedInAnyVehicle(ped, false) or IsPlayerDead(PlayerId()) then
				-- bye, vehicle
				isInVehicle = false
				SendNUIMessage({
					type = "left-vehicle"
				})
			else
				-- emit speed
				local vehicle = GetVehiclePedIsUsing(ped)
				local gear = GetVehicleCurrentGear(vehicle)
				local speed = math.floor(GetEntitySpeed(vehicle) * 2.236936)
				local health = math.floor((GetVehicleEngineHealth(vehicle) + GetVehicleBodyHealth(vehicle)) / 20)
				local rpm = math.floor(GetVehicleCurrentRpm(vehicle) * 100)
				local changed = false
				local diff = { type = "update-stats" }

				if currentGear ~= gear then
					changed = true
					currentGear = gear
					diff["gear"] = gear
				end

				if currentSpeed ~= speed then
					changed = true
					currentSpeed = speed
					diff["speed"] = speed
				end

				if currentRPM ~= rpm then
					changed = true
					currentRPM = rpm
					diff["rpm"] = rpm
				end

				if currentHealth ~= health then
					changed = true
					currentHealth = health
					diff["health"] = health
				end

				if changed then
					SendNUIMessage(diff)
				end
			end
		end
		Citizen.Wait(10)
	end
end)
