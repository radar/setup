package ruby

func Run() error {
	steps := []func() error{
		checkRuby,
		checkBundler,
		checkDependencies,
	}

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}
