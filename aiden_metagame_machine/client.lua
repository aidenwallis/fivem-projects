RegisterKeyMapping("+focus:nui", "Focus NUI", "keyboard", "f7")

local focused = false

RegisterCommand("+focus:nui", function()
    focused = not focused
    print(focused)
    SetNuiFocus(focused, focused)
end)

RegisterNUICallback("UnfocusNUI", function(data, cb)
    focused = false
    SetNuiFocus(false, false)
    cb({ ok = true })
end)
