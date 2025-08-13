local M = {}

M.present = function ()
  local info = debug.getinfo(1, "S")
  local current_path = info.source:sub(2)
  local server_path = current_path:gsub("lua/present/init.lua", "main.go")

  local cmd = {"go", "run", server_path}
  local job_id = vim.fn.jobstart(cmd, {
    on_stdout = function (id, data, event)
      print(vim.inspect(data), event)
    end,
    rpc = true, -- on_stdout gets disabled if this is on
  })

  local result = vim.fn.rpcrequest(job_id, "Present.NextSlide")

  print("result from sending rpc request to Presnt.NextSlide: ", result)
end

return M
