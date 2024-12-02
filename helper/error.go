package helper

// import "fmt"

func ErrorPanic(err error) error {
	if err != nil {
		// panic(fmt.Sprintf("Panic Error: %d", err))
		return err
	}
	return nil
}
