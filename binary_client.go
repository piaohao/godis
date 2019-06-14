package godis

type BinaryClient struct {
	*Connection
	//Host              string
	//Port              int
	//ConnectionTimeout int
	//SoTimeout         int
	Password string
	Db       int
	//IsInMulti         bool
	//IsInWatch         bool
	//Ssl               bool
}

func NewBinaryClient(options RedisOptions) *BinaryClient {
	db := 0
	if options.Db != 0 {
		db = options.Db
	}
	client := &BinaryClient{
		//Host:              options.Host,
		//Port:              options.Port,
		//ConnectionTimeout: options.ConnectionTimeout,
		//SoTimeout:         options.SoTimeout,
		Password: options.Password,
		Db:       db,
		//IsInMulti:         options.IsInMulti,
		//IsInWatch:         options.IsInWatch,
		//Ssl:               options.Ssl,
	}
	client.Connection = NewConnection(options.Host, options.Port, options.ConnectionTimeout, options.SoTimeout, options.Ssl)
	return client
}

func (c *BinaryClient) Connect() error {
	err := c.Connection.Connect()
	if err != nil {
		return err
	}
	if c.Password != "" {
		err = c.Auth(c.Password)
		if err != nil {
			return err
		}
		_, err = c.getStatusCodeReply()
		if err != nil {
			return err
		}
	}
	if c.Db > 0 {
		err = c.Select(c.Db)
		if err != nil {
			return err
		}
		_, err = c.getStatusCodeReply()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *BinaryClient) Auth(password string) error {
	c.Password = password
	return c.SendCommand(CMD_AUTH, []byte(password))
}

func (c *BinaryClient) Select(index int) error {
	return c.SendCommand(CMD_SELECT, IntToByteArray(index))
}

func (c *BinaryClient) Set(key, value []byte) error {
	return c.SendCommand(CMD_SET, key, value)
}

func (c *BinaryClient) Get(key []byte) error {
	return c.SendCommand(CMD_GET, key)
}

func (c *BinaryClient) Del(key []byte) error {
	return c.SendCommand(CMD_DEL, key)
}

func (c *BinaryClient) DelBatch(keys ...[]byte) error {
	return c.SendCommand(CMD_DEL, keys...)
}
