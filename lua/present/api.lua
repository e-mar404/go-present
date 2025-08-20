local config = require "present.config"

local M = {}

M.rpc_status = function ()
  local running

  config.rpc.client:connect(config.rpc.host, config.rpc.port, function (err)
    running = err
    print(err)
  end)

  return running
end

return M
