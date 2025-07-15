package authcommandrequest

type SaveAccountRequest struct {
	Token       string `json:"token" binding:"required"`                         // token đăng kí thông tin người dùng
	Email       string `json:"email" binding:"required,email"`                   // Email bắt buộc và phải đúng định dạng
	Password    string `json:"password" binding:"required,min=8"`                // Password tối thiểu 8 ký tự
	ConfirmPass string `json:"confirm_pass" binding:"required,eqfield=Password"` // Xác nhận password phải giống Password
	Name        string `json:"name" binding:"required"`                          // Tên người dùng
	Phone       string `json:"phone" binding:"omitempty,e164"`                   // Số điện thoại (optional, theo chuẩn E.164)
	Gender      int8   `json:"gender" binding:"omitempty,oneof=0 1 2"`           // Giới tính: 0=male, 1=female, 2=other
	Birthday    string `json:"birthday" binding:"omitempty"`                     // Ngày sinh (optional), định dạng yyyy-mm-dd
}
