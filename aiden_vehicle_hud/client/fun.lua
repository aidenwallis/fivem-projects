-- This will make your car lights, horn, and colour flash faster the quicker you drive
Citizen.CreateThread(function()
    local lights = false
    local maxThreshold = 1000
    local minThreshold = 100
    local diffableAmount = maxThreshold - minThreshold

    while true do
        Citizen.Wait(0)

        local ped = PlayerPedId()
        local wait = maxThreshold

        if IsPedInAnyVehicle(ped) then
            local vehicle = GetVehiclePedIsUsing(ped)
            local colors = GetVehicleColours(vehicle)
            local newColor = math.random(0, 159)
            local speed = math.floor(GetEntitySpeed(vehicle) * 2.236936)
            wait = math.max(minThreshold, maxThreshold - (diffableAmount * (math.min(100, speed) / 100)))
            SetVehicleColours(vehicle, newColor, newColor)
            SoundVehicleHornThisFrame(vehicle)

            local lightState = 1
            if lights then
                lightState = 2
            end
            SetVehicleLights(vehicle, lightState)

            lights = not lights
        end

        Citizen.Wait(wait)
    end
end)
