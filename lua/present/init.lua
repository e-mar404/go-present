local uv = vim.uv

local M = {}

M.present = function ()
  local rpc_port = 8081
  local host = "127.0.0.1"
  local client = uv.new_tcp()

  -- should first check if the server is running rpc server 
  -- if it isnt then it should do all the things bellow
  client:connect(host, rpc_port, function (err)
    --check if connection was successful 
  end)

  -- start server
  -- local info = debug.getinfo(1, "S")
  -- local current_path = info.source:sub(2)
  -- local server_path = current_path:gsub("lua/present/init.lua", "main.go")
  --
  -- local on_exit = function(obj)
  --   print(obj.code)
  --   print(obj.signal)
  --   print(obj.stdout)
  --   print(obj.stderr)
  -- end
  -- local _obj = vim.system({"go", "run", server_path}, {}, on_exit):wait()

  -- will use defer_fn to periodically check after it gets turned and if it times out then close the entire program and exit with an error
  vim.defer_fn(function ()
  end, 5000)

  local chan = vim.fn.sockconnect("tcp", "127.0.0.1:8081", {
    rpc = true,
  })
  local result = vim.fn.rpcrequest(chan, "Present.NextSlide")
  vim.fn.chanclose(chan)
  print("result from sending rpc request to Presnt.NextSlide: ", vim.inspect(result))
end

return M
