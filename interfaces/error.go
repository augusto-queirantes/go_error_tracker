package interfaces

type Error struct {
    Name string `json:"name"`
    StackTrace string `json:"stack_trace"`
}
