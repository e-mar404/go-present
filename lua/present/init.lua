local M = {}

M.present = function ()
  local info = debug.getinfo(1, "S")
  local current_path = info.source:sub(2)
  -- i want to change lua/present/init.lua with main.go
  local server_path = current_path:gsub("lua/present/init.lua", "main.go")

  print "starting presentation..."

  local cmd = {"go", "run", server_path}
  local _job_id = vim.fn.jobstart(cmd, {
    on_stdout = function (id, data, event)
      print(id, vim.inspect(data), event)
    end,
    stdout_buffered = true,
  })
end

return M
