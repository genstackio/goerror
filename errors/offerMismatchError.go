package errors

type OfferMismatchError struct {
	ExpectedOffer string
	ActualOffer   string
}

func (err OfferMismatchError) Error() string {
	return "Offer mismatch error (expected: " + err.ExpectedOffer + ", actual: " + err.ActualOffer + ")"
}

func (err OfferMismatchError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["expectedOffer"] = err.ExpectedOffer
	m["actualOffer"] = err.ActualOffer
	return m
}

func (err OfferMismatchError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "offer_mismatch",
		Message:    err.Error(),
		Code:       12401,
		StatusCode: 400,
		Data:       err.GetData(),
	}
}
