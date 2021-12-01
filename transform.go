package main

type ExternalConfig interface {
	Transform()
}

// Environment variables file external configuration
type EnvFileExternalConfig struct {
	OutDir   string
	FileName string
	Values   map[string]interface{}
}

func (envFileCtx EnvFileExternalConfig) Transform() {
	ctx := TemplateExportContext{TYPE_ENVFILE, envFileCtx.OutDir, envFileCtx.FileName}
	ctx.RunTemplate(envFileCtx.Values)
}

// K8s ConfigMap file external configuration
type ConfigMapFileExternalConfig struct {
	OutDir   string
	FileName string
	Values   map[string]interface{}
}

func (cmFileCtx ConfigMapFileExternalConfig) Transform() {
	ctx := TemplateExportContext{TYPE_CONFIGMAP, cmFileCtx.OutDir, cmFileCtx.FileName}
	ctx.RunTemplate(cmFileCtx.Values)
}

// Helm ConfigMap
type HelmConfigMapFileExternalConfig struct {
	OutDir   string
	FileName string
	Values   map[string]interface{}
}

func (helmFileCtx HelmConfigMapFileExternalConfig) Transform() {
	// TODO implement
}

// Factory function
func CreateExternalConfig(cfg Config, values map[string]interface{}) ExternalConfig {
	fileName := ExtractFileName(cfg.fileName)
	switch cfg.outType {
	case "cm":
		return ConfigMapFileExternalConfig{cfg.outputDir, fileName, values}
	case "helm":
		return HelmConfigMapFileExternalConfig{cfg.outputDir, fileName, values}
	default:
		return EnvFileExternalConfig{cfg.outputDir, fileName, values}
	}
}
