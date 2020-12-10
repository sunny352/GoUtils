package goutils

func Go(action func()) {
	go func() {
		defer func() {
			if err := recover(); nil != err {
				return
			}
		}()
		action()
	}()
}

func GoWithLog(action func(), logger func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); nil != err {
				logger(err)
				return
			}
		}()
		action()
	}()
}

func GoLoop(action func() error) {
	go func() {
		defer func() {
			if err := recover(); nil != err {
				return
			}
		}()
		for {
			err := action()
			if nil != err {
				break
			}
		}
	}()
}

func GoLoopWithLog(action func() error, logger func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); nil != err {
				logger(err)
				return
			}
		}()
		for {
			err := action()
			if nil != err {
				logger(err)
				break
			}
		}
	}()
}

func GoInfinityLoop(action func() error) {
	go func() {
		for {
			err := func() error {
				defer func() {
					if err := recover(); nil != err {
						return
					}
				}()
				for {
					err := action()
					if nil != err {
						return err
					}
				}
			}()
			if nil != err {
				break
			}
		}
	}()
}

func GoInfinityLoopWithLog(action func() error, logger func(err interface{})) {
	go func() {
		for {
			err := func() error {
				defer func() {
					if err := recover(); nil != err {
						logger(err)
						return
					}
				}()
				for {
					err := action()
					if nil != err {
						logger(err)
						return err
					}
				}
			}()
			if nil != err {
				break
			}
		}
	}()
}
