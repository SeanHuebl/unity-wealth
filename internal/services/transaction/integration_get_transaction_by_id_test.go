package transaction

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/seanhuebl/unity-wealth/internal/database"
	"github.com/stretchr/testify/require"
)

func TestIntegrationGetTransactionByID(t *testing.T) {
	tests := []struct {
		name                string
		userID              uuid.UUID
		txnID               uuid.UUID
		req                 NewTransactionRequest
		txErr               error
		expectedTxErrSubstr string
	}{
		{
			name:   "successful retrieval",
			userID: uuid.New(),
			txnID:  uuid.New(),
			req: NewTransactionRequest{
				Date:             "2025-02-25",
				Merchant:         "costco",
				Amount:           197.25,
				DetailedCategory: 40,
			},
			txErr:               nil,
			expectedTxErrSubstr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			db, err := sql.Open("sqlite3", ":memory:")
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("PRAGMA foreign_keys = ON")
			require.NoError(t, err)

			CreateTestingSchema(t, db)

			transactionalQ := database.NewRealTransactionalQuerier(database.New(db))
			txQ := database.NewRealTransactionQuerier(transactionalQ)
			userQ := database.NewRealUserQuerier(transactionalQ)

			seedGetTxByIDTestData(t, db, userQ, tc.userID, txQ, tc.txnID)

			svc := NewTransactionService(txQ)

			tx, err := svc.GetTransactionByID(ctx, tc.userID.String(), tc.txnID.String())
			require.NoError(t, err)
			require.NotNil(t, tx)

			expectedTx := &Transaction{
				ID:               tc.txnID.String(),
				UserID:           tc.userID.String(),
				Date:             tc.req.Date,
				Merchant:         tc.req.Merchant,
				Amount:           tc.req.Amount,
				DetailedCategory: tc.req.DetailedCategory,
			}

			if diff := cmp.Diff(tx, expectedTx); diff != "" {
				t.Errorf("transaction mismatch (-want +got)\n%s", diff)
			}
		})

	}
}

func seedGetTxByIDTestData(t *testing.T, db *sql.DB, userQ database.UserQuerier, userID uuid.UUID, txQ database.TransactionQuerier, txID uuid.UUID) {
	SeedTestUser(t, userQ, userID)
	SeedTestCategories(t, db)
	SeedTestTransaction(t, txQ, userID, txID, &NewTransactionRequest{
		Date:             "2025-02-25",
		Merchant:         "costco",
		Amount:           197.25,
		DetailedCategory: 40,
	})
}
