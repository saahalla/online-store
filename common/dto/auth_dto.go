package dto

type (
	UserDB struct {
		ID         int    `db:"id" goqu:"skipinsert"`
		Username   string `db:"username"`
		Password   string `db:"password"`
		Email      string `db:"email"`
		Phone      string `db:"phone"`
		UserRoleID int    `db:"user_role_id"`
		DefaultDate
	}

	UserDataResp struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Role     string `json:"role"`
	}
)

// req resp
type (

	// register
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	RegisterResponse struct {
		DefaultResponse
	}

	// login
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		DefaultResponse
		JwtToken string `json:"jwt_token"`
	}

	//
	ClaimJWTData struct {
		UserID   int
		Username string
		Exp      int
	}
)
