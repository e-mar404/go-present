local uv = vim.uv
local config = require "present.config"

local M = {}

-- worry about how it starts later for now i will run locally
-- M.rpc_start = function ()
--   vim.print "rpc server not running, starting go server"
--
--   local info = debug.getinfo(1, "S")
--   local current_path = info.source:sub(2)
--   local server_path = current_path:gsub("plugin/init.lua", "main.go")
--
--   local on_exit = function(code, signal)
--     print("exit code: ", code)
--     print("exit signal: ", signal)
--   end
--
--   local handle, pid = uv.spawn("go", {
--     args = {"run", server_path}
--   }, on_exit)
--
--   print("server handle: ", handle, ", pid: ", pid)
-- end

M.rpc_call = function (cmd, ...)
  local addr = config.rpc.host .. ":" .. config.rpc.port
  local chan = vim.fn.sockconnect("tcp", addr, {
    rpc = true,
  })

  local result = vim.fn.rpcrequest(chan, cmd, vim.fn.expand("%:p:h"))
  vim.fn.chanclose(chan)

  return result
end

return M
