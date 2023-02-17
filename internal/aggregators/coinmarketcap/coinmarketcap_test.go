package coinmarketcap

//var l *logga.Logga
//var cfg *aggregatorConfig.Config
//var DB *sql.DB
//var tx *sql.Tx
//
//func setup() {
//	_ = os.Setenv("APP_ENV", "test")
//
//	l = logga.New()
//
//	// Setup aggregatorConfig.
//	var err error
//	cfg, err = config.New(l, "../../../aggregatorConfig.json.test")
//	if err != nil {
//		l.Lg.Error().Msgf("error starting main. %w", err.Error())
//		os.Exit(1)
//	}
//
//	l = logga.New()
//	DB, err = db.New(l, cfg.DB)
//	if err != nil {
//		l.Lg.Error().Msg(err.Error())
//		os.Exit(1)
//	}
//	tx, err = DB.Begin()
//	if err != nil {
//		l.Lg.Error().Msg(err.Error())
//		os.Exit(1)
//	}
//}
//
//func tearDown() {
//	tx.Rollback()
//}
//
//func TestRun(t *testing.T) {
//
//	setup()
//	defer tearDown()
//
//	// Setup test http server.
//	testJsonResponse, err := testfuncs.GetTestJsonResponse("coin_response.json")
//	if err != nil {
//		t.Fatalf("error getting server response. %s", err)
//	}
//	server := testfuncs.SetupTestServer(testJsonResponse)
//	defer server.Close()
//
//	cfg.CMC.Host = server.URL // Set URL to that of the test response
//
//	chm := New(cfg.CMC, l)
//
//	cmc := New(cfg.CMC, l, DB, chm)
//
//	agg := aggregatorengine.New(DB, l)
//	err = agg.UpdateLatestHistory(cmc)
//	if err != nil {
//		l.Lg.Error().Msg(err.Error())
//		os.Exit(1)
//	}
//
//	var currencies CurrencyData
//	err = json.Unmarshal(testJsonResponse, &currencies)
//	if err != nil {
//		t.Errorf("could not unmarshal JSON. %s", err)
//	}
//
//	jsonRec1 := currencies.Currencies[0]
//
//	// Test record in currency table.
//	var curDbRec1 database.Currency
//
//	cm := currency.New(DB, l)
//	err = cm.GetRecordBySymbol(&curDbRec1, jsonRec1.Symbol)
//	if err != nil {
//		t.Errorf("error with GetRecordBySymbol. %s", err)
//	}
//
//	if jsonRec1.Name != curDbRec1.Name {
//		t.Errorf("name not as expected. Got: %s; found: %s;", jsonRec1.Name, curDbRec1.Name)
//	}
//
//	if jsonRec1.Symbol != curDbRec1.Symbol {
//		t.Errorf("symbol not as expected. Got: %s; found: %s;", jsonRec1.Symbol, curDbRec1.Symbol)
//	}
//
//	err = cm.DeleteRecord(curDbRec1.ID)
//	if err != nil {
//		t.Errorf("Unable to remove test record from database")
//	}
//
//}
