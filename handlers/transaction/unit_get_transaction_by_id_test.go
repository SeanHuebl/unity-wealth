package transaction_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	htx "github.com/seanhuebl/unity-wealth/handlers/transaction"
	"github.com/seanhuebl/unity-wealth/internal/constants"
	"github.com/seanhuebl/unity-wealth/internal/database"
	dbmocks "github.com/seanhuebl/unity-wealth/internal/mocks/database"
	"github.com/seanhuebl/unity-wealth/internal/services/transaction"
	"github.com/seanhuebl/unity-wealth/internal/testfixtures"
	"github.com/seanhuebl/unity-wealth/internal/testhelpers"
	"github.com/seanhuebl/unity-wealth/internal/testmodels"
)

func TestGetTxByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []testmodels.GetTxTestCase{
		{
			BaseHTTPTestCase: testfixtures.NilUserID,
			TxID:             uuid.NewString(),
		},
		{
			BaseHTTPTestCase: testfixtures.InvalidUserID,
			TxID:             uuid.NewString(),
		},
		{
			BaseHTTPTestCase: testfixtures.InvalidTxID,
			TxID:             "",
		},
		{
			BaseHTTPTestCase: testmodels.BaseHTTPTestCase{

				Name:               "error getting tx",
				UserID:             uuid.New(),
				ExpectedError:      "unable to get transaction",
				ExpectedStatusCode: http.StatusInternalServerError,
			},
			TxID:  uuid.NewString(),
			TxErr: errors.New("error getting transaction"),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			mockTxQ := dbmocks.NewTransactionQuerier(t)
			w := httptest.NewRecorder()
			svc := transaction.NewTransactionService(mockTxQ)
			req := httptest.NewRequest("GET", fmt.Sprintf("/transactions/%v", tc.TxID), nil)

			dummyRow := database.GetUserTransactionByIDRow{
				ID:                 tc.TxID,
				UserID:             tc.UserID.String(),
				TransactionDate:    "2025-03-05",
				Merchant:           "costco",
				AmountCents:        12598,
				DetailedCategoryID: 40,
			}

			if tc.UserIDErr == nil && tc.TxID != "" {
				mockTxQ.On("GetUserTransactionByID", context.Background(), database.GetUserTransactionByIDParams{
					UserID: tc.UserID.String(),
					ID:     tc.TxID,
				}).Return(dummyRow, tc.TxErr)
			}

			h := htx.NewHandler(svc)

			router := gin.New()
			if tc.TxID == "" {
				c, _ := gin.CreateTestContext(w)
				c.Request = req
				c.Params = gin.Params{{Key: "id", Value: ""}}
				if tc.Name == "unauthorized: user ID not UUID" {
					c.Set(string(constants.UserIDKey), "userID")
				} else {
					c.Set(string(constants.UserIDKey), tc.UserID)
				}
				h.GetTransactionByID(c)
			} else {

				router.GET("/transactions/:id", func(c *gin.Context) {
					if tc.Name == "unauthorized: user ID not UUID" {
						c.Set(string(constants.UserIDKey), "userID")
					} else {
						c.Set(string(constants.UserIDKey), tc.UserID)
					}
					h.GetTransactionByID(c)
				})
				router.ServeHTTP(w, req)
			}

			actualResponse := testhelpers.ProcessResponse(w, t)
			testhelpers.CheckTxHTTPResponse(t, w, tc, actualResponse)
			mockTxQ.AssertExpectations(t)
		})
	}
}
