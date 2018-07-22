package jwt_test

import (
	"testing"

	"github.com/err0r500/go-realworld-clean/implem/jwt.authHandler"
	"github.com/stretchr/testify/assert"
)

func TestUserToken_happyCase(t *testing.T) {
	testUserID := "userID"

	tH := jwt.NewTokenHandler("theJWTsalt")
	token, err := tH.GenUserToken(testUserID)
	assert.NoError(t, err)

	userID, err := tH.GetUserName(token)
	assert.NoError(t, err)
	assert.Equal(t, testUserID, userID)
}

func TestUserToken_GenToken_fails(t *testing.T) {
	tH := jwt.NewTokenHandler("theJWTsalt")
	token, err := tH.GenUserToken("")
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestUserToken_GetUserID_fails(t *testing.T) {
	tH := jwt.NewTokenHandler("theJWTsalt")
	token, err := tH.GenUserToken("userID")
	assert.NoError(t, err)

	t.Run("otherSalt", func(t *testing.T) {
		tH2 := jwt.NewTokenHandler("otherSalt")
		userID, err := tH2.GetUserName(token)
		assert.Error(t, err)
		assert.Equal(t, "", userID)
	})
}

//
//func TestTokenGenType(t *testing.T) {
//	tG := jwt.tokenHandler{i.Salt}
//	var tests = []struct {
//		test     d.Admin
//		expected bool
//	}{
//		{d.Admin{ID: 123}, true},
//		{d.Admin{}, false},
//	}
//	for k, test := range tests {
//		token, err := tG.GenToken(d.Admin(test.test))
//
//		if test.expected == false && err == nil {
//			t.Errorf("GenToken() returned no error on case #%d", k)
//		} else if test.expected == true && err != nil {
//			t.Errorf("GenToken() returned an unexpected error on case #%d", k)
//		} else if test.expected == true && err == nil {
//			// try to Parse valid signed Claims
//			// test if the sign key is actually checked
//			_, err := i.GetAdminClaims(token, []byte("thiskeyShouldntPass"))
//			if err == nil {
//				t.Errorf("GetAdminClaims should fail on #%d", k)
//			}
//			admin, err := i.GetAdminClaims(token, i.Salt)
//			if err != nil {
//				t.Errorf("GetAdminClaims should pass on #%d", k)
//			}
//
//			// test if the claims haven't been altered
//			if admin.AdminID != test.test.Name {
//				t.Errorf("the JWT altered the data on case#%d", k)
//			}
//		}
//	}
//}
//func TestHostTokenGenType(t *testing.T) {
//	tG := i.tokenHandler{i.Salt}
//	var tests = []struct {
//		test     i.HostProduct
//		expected bool
//	}{
//		{i.HostProduct{UserID: 316, LBCUserID: 2791, LBCProductID: 108}, true},
//	}
//	for k, test := range tests {
//		token, err := tG.GenToken(test.test)
//		log.Print(token)
//		if !test.expected && err == nil {
//			t.Errorf("GenToken() returned no error on case #%d", k)
//		} else if test.expected && err != nil {
//			t.Errorf("GenToken() returned an unexpected error on case #%d : %s", k, err)
//		} else if test.expected && err == nil {
//			// try to Parse valid signed Claims
//			// test if the sign key is actually checked
//			_, err := i.GetHostClaims(token, []byte("thiskeyShouldntPass"))
//			if err == nil {
//				t.Errorf("GetAdminClaims should fail on #%d", k)
//			}
//			hC, err := i.GetHostClaims(token, i.Salt)
//			if err != nil {
//				t.Errorf("GetAdminClaims should pass on #%d", k)
//			}
//
//			// test if the claims haven't been altered
//			if hC.UserID != test.test.UserID {
//				t.Errorf("the JWT altered the data on case#%d", k)
//			}
//		}
//	}
//}
//func TestTokenSignatureCheck(t *testing.T) {
//	otherSecret := []byte("thisKeyShouldntPass")
//	badtG := i.tokenHandler{otherSecret}
//	token, err := badtG.GenToken(d.Admin{ID: 123})
//	if err != nil {
//		t.Errorf("GenToken() failed to generate AdminToken")
//	}
//	_, err = i.GetAdminClaims(token, otherSecret)
//	if err != nil {
//		t.Errorf("GetAdminClaims() should be able to return Admin with this key")
//	}
//
//	_, err = i.GetAdminClaims(token, i.Salt)
//	if err == nil {
//		t.Errorf("GetAdminClaims() should failed : wrong secret")
//	}
//}
//
//func TestTokenExpiracyCheck(t *testing.T) {
//	exptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, i.NewAdminClaims(1, -2*time.Hour))
//	exptokenString, _ := exptoken.SignedString(i.Salt)
//	_, err := i.GetAdminClaims(exptokenString, i.Salt)
//	if err == nil {
//		t.Errorf("GetAdminClaims() should failed : expired token")
//	}
//}
