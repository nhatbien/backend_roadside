package request

type SaveRescueUnitRequest struct {
	Name     string  `json:"name,omitempty" validate:"required"`
	Phone    string  `json:"phone,omitempty" validate:"required"`
	Email    string  `json:"email,omitempty" validate:"required"`
	Address  string  `json:"address,omitempty" validate:"required"`
	Password string  `json:"password,omitempty" validate:"required"`
	Lat      float64 `json:"lat,omitempty"`
	Lng      float64 `json:"lng,omitempty"`
	Type     string  `json:"type,omitempty" validate:"required"`
	Status   int     `json:"status,omitempty" validate:"required"`
}

type UpdateLocationResuceUnitRequest struct {
	Lat float64 `json:"lat,omitempty" validate:"required"`
	Lng float64 `json:"lng,omitempty" validate:"required"`
}

type UpdateRescueUnitRequest struct {
	Name     string  `json:"name,omitempty" validate:"required"`
	Phone    string  `json:"phone,omitempty" validate:"required"`
	Password string  `json:"password,omitempty"`
	Email    string  `json:"email,omitempty" validate:"required"`
	Address  string  `json:"address,omitempty" validate:"required"`
	Lat      float64 `json:"lat,omitempty" validate:"required"`
	Lng      float64 `json:"lng,omitempty" validate:"required"`
	Type     int     `json:"type,omitempty" validate:"required"`
	Status   int     `json:"status,omitempty" validate:"required"`
}

type GetRescueUnitRequest struct {
	Id string `json:"id,omitempty" validate:"required"`
}

type GetRescueUnitsRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type GetRescueUnitsByLocationAndTypeRequest struct {
	Lat     float64 `json:"lat,omitempty" validate:"required"`
	Lng     float64 `json:"lng,omitempty" validate:"required"`
	Type    int     `json:"type,omitempty" validate:"required"`
	Offset  int     `json:"offset,omitempty"`
	Limit   int     `json:"limit,omitempty"`
	Radius  int     `json:"radius,omitempty"`
	Keyword string  `json:"keyword,omitempty"`
}

type GetRescueUnitsByLocationRequest struct {
	Lat     float64 `json:"lat,omitempty" validate:"required"`
	Lng     float64 `json:"lng,omitempty" validate:"required"`
	Offset  int     `json:"offset,omitempty"`
	Limit   int     `json:"limit,omitempty"`
	Radius  int     `json:"radius,omitempty"`
	Keyword string  `json:"keyword,omitempty"`
}

type GetRescueUnitsByTypeRequest struct {
	Type    int    `json:"type,omitempty" validate:"required"`
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Keyword string `json:"keyword,omitempty"`
}

type DeleteRescueUnitRequest struct {
	Id string `json:"id,omitempty" validate:"required"`
}
