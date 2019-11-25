package main

import (
	"M-GateDBConfig/adapter"
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"
	"M-GateDBConfig/provider/fileconfig"
	"M-GateDBConfig/provider/postgres"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type (
	TestHandler struct {
		dbHandler *gorm.DB
	}

	TestFactory struct {
		dbconf  model.DBConfigurationModel
		eng     engine.EnginesFactory
		dbhand  engine.StorageFactory
		adapter adapter.Handler
	}
)

// real value
var (
	DBConf = model.DBConfigurationModel{
		Dbname:   "hasbidatabase",
		Password: "dimskii",
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "hasbi",
	}
)

// TestGetDBConfig test getting file config
func TestGetDBConfig(T *testing.T) {
	// Test if value JSON is not same

	// // create file config
	// EditDBconfValue(testValue)

	// test - get fileconfig
	db := fileconfig.GetDBConfig()

	checkdbconfvalue(DBConf, db, T)

	// return real value for Development
	fileconfig.CreateDBConfig()

}

// TestNewStorage for checking Table in DB exist or not.
func TestNewStorage(T *testing.T) {
	// 1. Test Table Existing or Not after Migration

	// DB Migration first
	postgres.NewStorage(DBConf)

	// Test Table Existing or Not
	// connect DB
	data := DBConf
	db := initDB(data)
	// checkTable exist or not
	result := db.Exec("SELECT * FROM REGISTRIES")
	if result.Error != nil {
		T.Errorf("No Table Exists %+v\n", result.Error)
	}
	result = db.Exec("SELECT * FROM PARAMETERS")
	if result.Error != nil {
		T.Errorf("No Table Exists %+v\n", result.Error)
	}
}

// TestTableParameter for checking Table in DB exist or not.
func TestTableParameter(T *testing.T) {
	// Test Insert Parameter
	var ParameterEngine engine.Parameter
	tf := &TestFactory{}
	tf.initializeApps()
	ParameterEngine = tf.eng.NewParameter()

	// Expected
	expected := model.Parameters{
		Key:   "haha",
		Value: "meeong",
	}

	// input value to the Case
	inp := engine.AddParamReq{
		ID:    "0",
		Key:   &expected.Key,
		Value: &expected.Value,
	}

	// Do the Test
	ParameterEngine.Add(&inp)

	// Actual
	actual := model.Parameters{}
	dbase := initDB(tf.dbconf)
	result := dbase.Find(&actual)
	if result.Error != nil {
		T.Errorf("Error Find %+v\n", result.Error)
	}

	assert.Equal(T, expected.Key, actual.Key)
	assert.Equal(T, expected.Value, actual.Value)

	// clean Test from DB
	dbase.Where("key = ?", expected.Key).Delete(&expected)

	// other test plan
	// 1. Test unique KEY
	// 2. Test no Table
	// 3. Test no Connection of Database

}

func (t *TestFactory) initializeApps() {

	// Initialize Application First
	t.dbconf = fileconfig.GetDBConfig()
	// Connect and Migrate DB
	t.dbhand = postgres.NewStorage(t.dbconf)
	// Prepare Engine for Use Case Logic
	t.eng = engine.NewEngine(t.dbhand)
}

func editDBconfValue(nan model.DBConfigurationModel) {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath("./")
	v.SetDefault("db", nan)
	v.WriteConfigAs("appconfig.json")
}

func checkdbconfvalue(expected, actual model.DBConfigurationModel, t *testing.T) {
	if expected != actual {
		t.Errorf("Not Same Configuration %+v. Got %+v\n", expected, actual)
	}
}

func initDB(data model.DBConfigurationModel) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.Dbname)
	dba, _ := gorm.Open("postgres", psqlInfo)
	return dba
}
