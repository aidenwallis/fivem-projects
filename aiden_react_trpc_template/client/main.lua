local currentToken = ""

function emitToken()
  SendNUIMessage({
    type = "new-token",
    token = currentToken
  })
end


Citizen.CreateThread(function()
  print("Waiting for token to be populated.")
  while currentToken == "" do
    currentToken = exports.aiden_auth:getSessionToken()
    Citizen.Wait(100)
  end

  print("Got token")
  emitToken()

  while true do
    local token = exports.aiden_auth:getSessionToken()

    if token ~= currentToken then
      currentToken = token
      emitToken()
    end

    Citizen.Wait(10000)
  end
end)