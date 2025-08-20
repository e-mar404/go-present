local api = require "present.api"
local config = require "present.config"
local uv = vim.uv

local rpc_status = api.rpc_status()

-- print(vim.inspect(rpc_status))

if not config.rpc.auto_start then
  -- vim.print "rpc server is set to not autostart"
  return
end


-- if not rpc_status.running then
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
--   return
-- end
--
-- vim.print "rpc server is already started"
