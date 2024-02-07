package auth

//import (
//	"dealls/core"
//	"dealls/core/v1/entity"
//	"dealls/core/v1/port/account"
//	"reflect"
//	"testing"
//)
//
//func TestGenerateToken(t *testing.T) {
//	type args struct {
//		claim                   entity.JwtClaim
//		accessTokenExpiredTime  int64
//		refreshTokenExpiredTime int64
//	}
//	tests := []struct {
//		name  string
//		args  args
//		want  *entity.Jwt
//		want1 *core.CustomError
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := generateToken(tt.args.claim, tt.args.accessTokenExpiredTime, tt.args.refreshTokenExpiredTime)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("generateToken() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("generateToken() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func TestNew(t *testing.T) {
//	type args struct {
//		accountRepository account.AccountRepository
//		cacheRepository   account.CacheRepository
//	}
//	tests := []struct {
//		name string
//		args args
//		want account.AuthUsecase
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewAuthUsecase(tt.args.accountRepository, tt.args.cacheRepository); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_accountServiceImpl_GetToken(t *testing.T) {
//	type fields struct {
//		accountRepository account.AccountRepository
//		cacheRepository   account.CacheRepository
//	}
//	type args struct {
//		ic       core.InternalContext
//		username string
//		password string
//		clientId string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *entity.Jwt
//		want1  *core.CustomError
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			service := &accountServiceImpl{
//				accountRepository: tt.fields.accountRepository,
//				cacheRepository:   tt.fields.cacheRepository,
//			}
//			got, got1 := service.Login(tt.args.ic, tt.args.username, tt.args.password, tt.args.clientId)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Login() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_accountServiceImpl_IsActiveToken(t *testing.T) {
//	type fields struct {
//		accountRepository account.AccountRepository
//		cacheRepository   account.CacheRepository
//	}
//	type args struct {
//		ic     core.InternalContext
//		userId string
//		token  string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *core.CustomError
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			service := &accountServiceImpl{
//				accountRepository: tt.fields.accountRepository,
//				cacheRepository:   tt.fields.cacheRepository,
//			}
//			if got := service.IsActiveToken(tt.args.ic, tt.args.userId, tt.args.token); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("IsActiveToken() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_accountServiceImpl_RefreshToken(t *testing.T) {
//	type fields struct {
//		accountRepository account.AccountRepository
//		cacheRepository   account.CacheRepository
//	}
//	type args struct {
//		ic     core.InternalContext
//		userId string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *entity.Jwt
//		want1  *core.CustomError
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			service := &accountServiceImpl{
//				accountRepository: tt.fields.accountRepository,
//				cacheRepository:   tt.fields.cacheRepository,
//			}
//			got, got1 := service.Relogin(tt.args.ic, tt.args.userId)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Relogin() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("Relogin() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_accountServiceImpl_RevokeToken(t *testing.T) {
//	type fields struct {
//		accountRepository account.AccountRepository
//		cacheRepository   account.CacheRepository
//	}
//	type args struct {
//		ic     core.InternalContext
//		userId string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *core.CustomError
//	}{
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			service := &accountServiceImpl{
//				accountRepository: tt.fields.accountRepository,
//				cacheRepository:   tt.fields.cacheRepository,
//			}
//			if got := service.Logout(tt.args.ic, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Logout() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
