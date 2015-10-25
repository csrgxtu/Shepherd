package confs

type Config struct {
  Host string `json:"host"`
  Port int32 `json:"port"`
  UserName string `json:"username"`
  Password string `json:"password"`
}

func ReadConfig() Config {
  var config Config

  config.Host = "127.0.0.1"
  config.Port = 27017

  return config
}
