local config = require "present.config"

local M = {}

M.present = function ()
  local addr = config.rpc.host .. ":" .. config.rpc.port
  local chan = vim.fn.sockconnect("tcp", addr, {
    rpc = true,
  })

  local result = vim.fn.rpcrequest(chan, "Present.SetPath", vim.fn.expand("%:p:h"))
  vim.fn.chanclose(chan)
  print("result from sending rpc request to Present.SetPath: ", vim.inspect(result))
end

return M
