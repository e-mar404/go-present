local uv = vim.uv
local rpc_port = 8081
local host = "127.0.0.1"
local client = uv.new_tcp()

client:connect(host, rpc_port, function (err)
  if err then
    vim.print "starting go server"

    local info = debug.getinfo(1, "S")
    local current_path = info.source:sub(2)
    local server_path = current_path:gsub("plugin/init.lua", "main.go")

    local on_exit = function(code, signal)
      print("exit code: ", code)
      print("exit signal: ", signal)
    end

    local handle, pid = uv.spawn("go", {
      args = {"run", server_path}
    }, on_exit)

    print("server handle: ", handle, ", pid: ", pid)
  else
    vim.print "server is already started"
  end
end)
