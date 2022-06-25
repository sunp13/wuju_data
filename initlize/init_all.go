package initlize

func InitAll() error {
	if err := initLog(); err != nil {
		return err
	}

	if err := initDB(); err != nil {
		return err
	}
	return nil
}
