package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_BuildDefaultsConfig = map[string]string{
	"":              "BuildDefaultsConfig controls the default information for Builds",
	"gitHTTPProxy":  "GitHTTPProxy is the location of the HTTPProxy for Git source",
	"gitHTTPSProxy": "GitHTTPSProxy is the location of the HTTPSProxy for Git source",
	"gitNoProxy":    "GitNoProxy is the list of domains for which the proxy should not be used",
	"env":           "Env is a set of default environment variables that will be applied to the build if the specified variables do not exist on the build",
	"sourceStrategyDefaults": "SourceStrategyDefaults are default values that apply to builds using the source strategy.",
}

func (BuildDefaultsConfig) SwaggerDoc() map[string]string {
	return map_BuildDefaultsConfig
}

var map_SourceStrategyDefaultsConfig = map[string]string{
	"":            "SourceStrategyDefaultsConfig contains values that apply to builds using the source strategy.",
	"incremental": "Incremental indicates if s2i build strategies should perform an incremental build or not",
}

func (SourceStrategyDefaultsConfig) SwaggerDoc() map[string]string {
	return map_SourceStrategyDefaultsConfig
}
