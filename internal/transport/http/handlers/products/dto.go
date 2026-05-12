package products

type createProductReq struct {
	Name          string `json:"name"`
	MeasurementId int    `json:"measurementId"`
	ParentId      *int   `json:"parentId"`
}
