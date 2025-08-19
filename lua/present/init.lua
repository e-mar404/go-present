local uv = vim.uv

local M = {}

M.present = function ()
  local rpc_port = 8081
  local host = "127.0.0.1"

  local chan = vim.fn.sockconnect("tcp", "127.0.0.1:8081", {
    rpc = true,
  })
  local result = vim.fn.rpcrequest(chan, "Present.SetPath", "/home/emar/plugins/present.nvim/md")
  vim.fn.chanclose(chan)
  print("result from sending rpc request to Present.SetPath: ", vim.inspect(result))
end

return M
