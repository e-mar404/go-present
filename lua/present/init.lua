local M = {}

M.present = function ()
  print "starting presentation..."

  print "calling go server to start"

  local _job_id = vim.fn.jobstart({"go", "run", "main.go"}, {
    on_stdout = function (id, data, event)
      print(id, vim.inspect(data), event)
    end,
    stdout_buffered = true,
  })

end

return M
