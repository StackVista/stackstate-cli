package asp

func AskQuestions(questions []InteractiveValue) error {
	for _, prompt := range questions {
		if err := prompt.Set(); err != nil {
			return err
		}
	}

	return nil
}

func AskQuestion(q InteractiveValue) error {
	return q.Set()
}
