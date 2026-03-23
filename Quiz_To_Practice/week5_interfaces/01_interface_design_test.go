package week5_interfaces

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestStorageInterface(t *testing.T) {
	t.Skip("Remove this line after implementing Storage interface")

	testStorage := func(t *testing.T, storage Storage, name string) {
		// Test Save
		err := storage.Save("key1", []byte("value1"))
		assert.NoError(t, err, "%s: Save should succeed", name)

		// Test Load
		data, err := storage.Load("key1")
		assert.NoError(t, err, "%s: Load should succeed", name)
		assert.Equal(t, []byte("value1"), data, "%s: Loaded data should match", name)

		// Test Exists
		exists := storage.Exists("key1")
		assert.True(t, exists, "%s: Key should exist", name)

		// Test List
		keys, err := storage.List()
		assert.NoError(t, err, "%s: List should succeed", name)
		assert.Contains(t, keys, "key1", "%s: List should contain key", name)

		// Test Delete
		err = storage.Delete("key1")
		assert.NoError(t, err, "%s: Delete should succeed", name)

		exists = storage.Exists("key1")
		assert.False(t, exists, "%s: Key should not exist after delete", name)
	}

	// Test different implementations
	// memStorage := &MemoryStorage{}
	// testStorage(t, memStorage, "MemoryStorage")

	// fileStorage := &FileStorage{baseDir: "/tmp/test"}
	// testStorage(t, fileStorage, "FileStorage")

	// Test CacheStorage with TTL
	// underlying := &MemoryStorage{}
	// cacheStorage := &CacheStorage{
	// 	underlying: underlying,
	// 	ttl: 100 * time.Millisecond,
	// }
	// cacheStorage.Save("temp", []byte("temporary"))
	// time.Sleep(200 * time.Millisecond)
	// exists := cacheStorage.Exists("temp")
	// assert.False(t, exists, "Cached item should expire")
}

func TestDatabaseInterface(t *testing.T) {
	t.Skip("Remove this line after implementing Database interface")

	testDatabase := func(t *testing.T, db Database, name string) {
		// Test Connect
		err := db.Connect("test://localhost/testdb")
		assert.NoError(t, err, "%s: Connect should succeed", name)

		// Test Execute
		affected, err := db.Execute("INSERT INTO users (name) VALUES (?)", "John")
		assert.NoError(t, err, "%s: Execute should succeed", name)
		assert.Greater(t, affected, int64(0), "%s: Should affect rows", name)

		// Test Query
		results, err := db.Query("SELECT * FROM users WHERE name = ?", "John")
		assert.NoError(t, err, "%s: Query should succeed", name)
		assert.NotEmpty(t, results, "%s: Should return results", name)

		// Test Transaction
		err = db.Transaction(func(tx Transaction) error {
			_, err := tx.Execute("UPDATE users SET name = ? WHERE name = ?", "Jane", "John")
			return err
		})
		assert.NoError(t, err, "%s: Transaction should succeed", name)

		// Test Close
		err = db.Close()
		assert.NoError(t, err, "%s: Close should succeed", name)
	}

	// Test different implementations
	// mockDB := &MockDatabase{}
	// testDatabase(t, mockDB, "MockDatabase")

	// sqliteDB := &SQLiteDatabase{}
	// testDatabase(t, sqliteDB, "SQLiteDatabase")
}

func TestLoggerInterface(t *testing.T) {
	t.Skip("Remove this line after implementing Logger interface")

	testLogger := func(t *testing.T, logger Logger, name string) {
		// Set level to Info
		logger.SetLevel(InfoLevel)

		// Debug should not be logged (below level)
		logger.Debug("Debug message", "key", "value")

		// Info and above should be logged
		logger.Info("Info message", "user", "john")
		logger.Warn("Warning message", "temperature", 95)
		logger.Error("Error message", "error", "connection failed")

		// Note: Fatal would typically exit, so we skip it in tests
	}

	// Test different implementations
	// consoleLogger := &ConsoleLogger{}
	// testLogger(t, consoleLogger, "ConsoleLogger")

	// var buf bytes.Buffer
	// fileLogger := &FileLogger{writer: &buf}
	// testLogger(t, fileLogger, "FileLogger")

	// Test MultiLogger
	// multiLogger := &MultiLogger{
	// 	loggers: []Logger{consoleLogger, fileLogger},
	// }
	// testLogger(t, multiLogger, "MultiLogger")
}

func TestSerializerInterface(t *testing.T) {
	t.Skip("Remove this line after implementing Serializer interface")

	type TestData struct {
		Name  string
		Age   int
		Email string
	}

	testData := TestData{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	testSerializer := func(t *testing.T, serializer Serializer, name string) {
		// Test Serialize
		data, err := serializer.Serialize(testData)
		assert.NoError(t, err, "%s: Serialize should succeed", name)
		assert.NotEmpty(t, data, "%s: Serialized data should not be empty", name)

		// Test Deserialize
		var result TestData
		err = serializer.Deserialize(data, &result)
		assert.NoError(t, err, "%s: Deserialize should succeed", name)
		assert.Equal(t, testData, result, "%s: Deserialized data should match", name)

		// Test GetContentType
		contentType := serializer.GetContentType()
		assert.NotEmpty(t, contentType, "%s: Content type should not be empty", name)
	}

	// Test different serializers
	// testSerializer(t, &JSONSerializer{}, "JSONSerializer")
	// testSerializer(t, &XMLSerializer{}, "XMLSerializer")
	// testSerializer(t, &BinarySerializer{}, "BinarySerializer")

	// Test factory function
	// jsonSer := GetSerializer("json")
	// assert.NotNil(t, jsonSer)
	// assert.IsType(t, &JSONSerializer{}, jsonSer)
}

func TestPaymentProcessor(t *testing.T) {
	t.Skip("Remove this line after implementing PaymentProcessor")

	testProcessor := func(t *testing.T, processor PaymentProcessor, details PaymentDetails, name string) {
		// Test ValidatePaymentMethod
		err := processor.ValidatePaymentMethod(details)
		assert.NoError(t, err, "%s: Validation should succeed", name)

		// Test ProcessPayment
		transaction, err := processor.ProcessPayment(100.00, "USD", details)
		assert.NoError(t, err, "%s: Payment processing should succeed", name)
		assert.NotEmpty(t, transaction.ID, "%s: Transaction should have ID", name)
		assert.Equal(t, 100.00, transaction.Amount, "%s: Amount should match", name)

		// Test GetTransactionStatus
		status, err := processor.GetTransactionStatus(transaction.ID)
		assert.NoError(t, err, "%s: Get status should succeed", name)
		assert.Equal(t, Completed, status, "%s: Status should be completed", name)

		// Test RefundPayment
		err = processor.RefundPayment(transaction.ID, 50.00)
		assert.NoError(t, err, "%s: Refund should succeed", name)
	}

	// Test different payment processors
	// ccDetails := &CreditCardDetails{
	// 	Number: "4111111111111111",
	// 	CVV:    "123",
	// 	Expiry: "12/25",
	// }
	// ccProcessor := &CreditCardProcessor{}
	// testProcessor(t, ccProcessor, ccDetails, "CreditCard")

	// ppDetails := &PayPalDetails{
	// 	Email:    "test@example.com",
	// 	Password: "password",
	// }
	// ppProcessor := &PayPalProcessor{}
	// testProcessor(t, ppProcessor, ppDetails, "PayPal")
}