local config = {}

config.rpc = {
  auto_start = true,
  port = 8081,
  host = "127.0.0.1",
  client = vim.uv.new_tcp()
}

return config
