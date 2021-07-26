package onepassword

type eventsService struct {
	client Client
}

func (es *eventsService) Introspect() (*IntrospectionResponse, error) {
	var rcvr IntrospectionResponse
	if err := es.client.Do(
		"/api/auth/introspect",
		"GET",
		nil,
		&rcvr,
	); err != nil {
		return nil, err
	}

	return &rcvr, nil
}
func (es *eventsService) GetItemUsages(resetCursor *ResetCursor, currCursor string) (*ItemUsageResponse, error) {
	var rcvr ItemUsageResponse
	if err := es.client.Do(
		"api/v1/itemusages",
		"POST",
		&CursorOptions{
			ResetCursor: resetCursor,
			CurrCursor:  currCursor,
		},
		&rcvr,
	); err != nil {
		return nil, err
	}

	return &rcvr, nil
}
func (es *eventsService) GetSignInAttempts(resetCursor *ResetCursor, currCursor string) (*SignInAttemptsResponse, error) {
	var rcvr SignInAttemptsResponse
	if err := es.client.Do(
		"api/v1/signinattempts",
		"POST",
		&CursorOptions{
			ResetCursor: resetCursor,
			CurrCursor:  currCursor,
		},
		&rcvr,
	); err != nil {
		return nil, err
	}

	return &rcvr, nil
}
