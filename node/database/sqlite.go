package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"path"

	"github.com/adrg/xdg"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

// Query executes a SQL query on the specified SQLite database.
// It supports both SELECT queries (returns rows as JSON) and non-SELECT queries (logs affected rows).
func Query(databaseName string, query string, ctx context.Context) ([]byte, error) {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", path.Join(xdg.DataHome, databaseName+".sqlite"))
	if err != nil {
		zap.L().Error("Failed to open database", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("failed to open database: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			zap.L().Error("Failed to close database connection", zap.String("database", databaseName), zap.Error(closeErr))
		}
	}()

	// Begin a transaction
	txn, err := db.Begin()
	if err != nil {
		zap.L().Error("Failed to begin transaction", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			zap.L().Error("Failed to rollback transaction", zap.String("database", databaseName), zap.Error(rollbackErr))
		}
	}()

	// Execute the query
	rows, err := txn.QueryContext(ctx, query)

	// Retrieve column names from the result set
	columns, err := rows.Columns()
	if err != nil {
		zap.L().Error("Failed to retrieve columns", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("failed to retrieve columns: %w", err)
	}

	// Prepare a slice to hold the result rows
	var resultRows []map[string]interface{}

	// Iterate over the rows and process each one
	for rows.Next() {
		// Prepare slices to hold column values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		// Assign pointers to the values slice
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the current row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			zap.L().Error("Failed to scan row", zap.String("database", databaseName), zap.Error(err))
			return []byte(""), fmt.Errorf("failed to scan row: %w", err)
		}

		// Map the column names to their corresponding values
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			// Convert []byte to string for readability
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}

			rowMap[col] = v
		}

		// Append the row to the result set
		resultRows = append(resultRows, rowMap)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		zap.L().Error("Error during row iteration", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("error during row iteration: %w", err)
	}

	// Commit the transaction
	if err := txn.Commit(); err != nil {
		zap.L().Error("Failed to commit transaction", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Convert the result rows to JSON
	jsonData, err := json.Marshal(resultRows)
	if err != nil {
		zap.L().Error("Failed to marshal rows to JSON", zap.String("database", databaseName), zap.Error(err))
		return []byte(""), fmt.Errorf("failed to marshal rows to JSON: %w", err)
	}

	zap.L().Info("SELECT query executed successfully", zap.String("database", databaseName), zap.String("query", query), zap.Int("rows_returned", len(resultRows)))
	return jsonData, nil // Return the JSON result
}

// Exec executes a non-SELECT SQL query (e.g., INSERT, UPDATE, DELETE) on the specified SQLite database.
// It returns the number of rows affected.
func Exec(databaseName string, query string, ctx context.Context) (int64, error) {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", path.Join(xdg.DataHome, databaseName+".sqlite"))
	if err != nil {
		zap.L().Error("Failed to open database", zap.String("database", databaseName), zap.Error(err))
		return 0, fmt.Errorf("failed to open database: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			zap.L().Error("Failed to close database connection", zap.String("database", databaseName), zap.Error(closeErr))
		}
	}()

	// Begin a transaction
	txn, err := db.Begin()
	if err != nil {
		zap.L().Error("Failed to begin transaction", zap.String("database", databaseName), zap.Error(err))
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			zap.L().Error("Failed to rollback transaction", zap.String("database", databaseName), zap.Error(rollbackErr))
		}
	}()

	// Execute the query
	result, err := txn.ExecContext(ctx, query)
	if err != nil {
		zap.L().Error("Failed to execute query", zap.String("database", databaseName), zap.String("query", query), zap.Error(err))
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	// Commit the transaction
	if err := txn.Commit(); err != nil {
		zap.L().Error("Failed to commit transaction", zap.String("database", databaseName), zap.Error(err))
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		zap.L().Error("Failed to retrieve rows affected", zap.String("database", databaseName), zap.Error(err))
		return 0, fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	zap.L().Info("Non-SELECT query executed successfully", zap.String("database", databaseName), zap.String("query", query), zap.Int64("rows_affected", rowsAffected))
	return rowsAffected, nil // Return the number of rows affected
}
