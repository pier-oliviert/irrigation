package config

import(
	"github.com/tsuru/config"
  "fmt"
  "log"
  "os"
  "strconv"
)

type Config struct {
  Port      uint16
  Path      string
  Database  map[string]*string
  Valves    []uint8
}

func Init(path string) (c *Config) {
  err := config.ReadConfigFile(fmt.Sprintf("%v/config.yml", path))
  if err != nil {
    log.Fatalln(`An error occured reading your configuration files.
    Is it possible you have not initialized irrigation yet?
    irrigation -initialize -path='/srv/http/irrigation'`)
  }

  c = &Config{}
  c.port()
  c.database()
  c.valves()
  c.Path = path
  return c
}

func AskForValue(option interface{}, msg string) {
  fmt.Println(msg)
  fmt.Scanln(option)
}

func (c *Config) Update() {
  updateConfigInternal(c)
  file := fmt.Sprintf("%v/config.yml", c.Path)
  err := os.Remove(file)
  if err != nil {
    log.Fatalln(fmt.Sprintf(`An error occured while updating the configuration file
    %v`,err))
  }
  err = config.WriteConfigFile(file, 0744)
  if err != nil {
    log.Fatalln(fmt.Sprintf(`An error occured while updating the configuration file
    %v`,err))
  }
}


// Private

func (c *Config) port() {
  value, err := config.GetInt("port")

  if err != nil {
    log.Println(fmt.Printf(
`Couldn't fetch port from the configuration file.
Irrigation will be launched on port 7777 until you fix this.
It should be set in %v
port: 7777
%v
`, c.Path, err))
    c.Port = 7777
  } else {
    c.Port = uint16(value)
  }
}

func (c *Config) database() {
  c.Database = map[string]*string{}
  name, err := config.GetString("database:name")

	if err != nil {
		log.Fatalln(`Configuration file is missing database information
    A correct configuration file would include the following in config.yml:
    database:
      name: 'databasename'
      user: 'mysqlUser'
     password: 'password'`)
	}

  c.Database["name"] = &name

  user, err := config.GetString("database:user")
  if err != nil {
    log.Fatalln(`An error happened when fetching the database configuration`)
  }
  c.Database["user"] = &user

  password, err := config.GetString("database:password")
  if err != nil {
    log.Fatalln(`An error happened when fetching the database configuration`)
  }
  c.Database["password"] = &password
}

func (c *Config) valves() {
	list, err := config.GetList("valves")
	if err != nil {
    log.Fatalln(`There was an issue trying to retrieve the valve list in your configuration file`)
	}

  valves := []uint8{}
  for _, value := range list {
    valve, err := strconv.ParseInt(value, 10, 8)
    if err != nil {
      log.Fatalln(fmt.Sprintf(`An error occured while parsing the valve numbers
      %v`, err))
    }
    valves = append(valves, uint8(valve))
  }
  c.Valves = valves
}

func updateConfigInternal(c *Config) {
  config.Set("port", c.Port)
  config.Set("database:name", *c.Database["name"])
  config.Set("database:user", *c.Database["user"])
  config.Set("database:password", *c.Database["password"])
}
