package kafka

type MockConsumer struct{}

func (m *MockConsumer) Poll() *Records {
	records := &Records{}
	records.Items = append(records.Items, "i am kafka mock consumer")
	return records
}
