local uv = vim.uv

local M = {}

M.present = function ()
  local rpc_port = 8081
  local host = "127.0.0.1"

  -- will use defer_fn to periodically check after it gets turned and if it times out then close the entire program and exit with an error
  -- vim.defer_fn(function ()
    -- end, 5000)
    --
    -- local chan = vim.fn.sockconnect("tcp", "127.0.0.1:8081", {
      --   rpc = true,
      -- })
      -- local result = vim.fn.rpcrequest(chan, "Present.NextSlide")
      -- vim.fn.chanclose(chan)
      -- print("result from sending rpc request to Presnt.NextSlide: ", vim.inspect(result))
end

return M
