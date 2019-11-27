package engine

type (
	// TestingEngineStruct Structure consist of Repository Interface.
	TestingEngineStruct struct {
		Param   ParameterRepository
		SysSett SystemSettingsRepository
	}
)

func (f *engineFactory) NewTestEngine() TestingEngineStruct {
	return TestingEngineStruct{
		Param:   f.NewParameterRepository(),
		SysSett: f.NewSystemSettingRespository(),
	}
}
