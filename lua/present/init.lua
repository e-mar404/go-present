local M = {}

M.present = function ()
  vim.notify "starting presentation..."

  vim.notify "starting go server"

  local job_id = vim.fn.jobstart("go run main.go", {
    on_stdout = function(something, data)
      print(vim.inspect(something), vim.inspect(data))
    end
  })
  print(job_id)
end

return M
