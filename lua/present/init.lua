local api = require "present.api"

local M = {}

M.present = function ()
  local result = api.rpc_call("Present.SetPath", {})

  print("result from sending rpc request to Present.SetPath: ", vim.inspect(result))
end

return M
