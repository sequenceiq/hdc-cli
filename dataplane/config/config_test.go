package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hortonworks/cb-cli/dataplane/common"
)

func TestWriteConfigToFileDirExists(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configwritetest")
	defer os.RemoveAll(tempDirName)
	os.MkdirAll(tempDirName+string(filepath.Separator)+common.Config_dir, 0700)

	WriteConfigToFile(tempDirName, "server", "output", "default", "workspace", "apikeyid", "privatekey")

	validateConfigContent(tempDirName, t)
}

func TestWriteConfigToFileDirNotExists(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configwritetest")
	defer os.RemoveAll(tempDirName)

	WriteConfigToFile(tempDirName, "server", "output", "default", "workspace", "apikeyid", "privatekey")

	validateConfigContent(tempDirName, t)
}

func validateConfigContent(tempDirName string, t *testing.T) {
	content, _ := ioutil.ReadFile(tempDirName + string(filepath.Separator) + common.Config_dir + string(filepath.Separator) + common.Config_file)

	expected := "default:\n  server: server\n  workspace: workspace\n  output: output\n  apikeyid: apikeyid\n  privatekey: privatekey\n"
	if string(content) != expected {
		t.Errorf("content not match %s == %s", expected, string(content))
	}
}

func TestReadConfig(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configreadtest")
	defer os.RemoveAll(tempDirName)

	os.MkdirAll(tempDirName+string(filepath.Separator)+common.Config_dir, 0700)
	password := "§±!@#$%^&*()_+-=[]{};'\\:\"/.,?><`~"
	ioutil.WriteFile(tempDirName+string(filepath.Separator)+common.Config_dir+string(filepath.Separator)+common.Config_file, []byte("default:\n  password: "+password+"\n  server: server\n  workspace: workspace\n  output: output\n"), 0700)

	config, err := ReadConfig(tempDirName, "default")

	if err != nil {
		t.Errorf("unable to read file: %s", err.Error())
	}
	if config.Server != "server" {
		t.Errorf("server not match server == %s", config.Server)
	}
	if config.Output != "output" {
		t.Errorf("output not match output == %s", config.Output)
	}
	if config.Workspace != "workspace" {
		t.Errorf("workspace not match workspace == %s", config.Workspace)
	}
}

func TestReadConfigReal(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configreadtest")
	defer os.RemoveAll(tempDirName)

	os.MkdirAll(tempDirName+string(filepath.Separator)+common.Config_dir, 0700)
	apiKey := "Y3JuOmFsdHVzOmlhbTp1cy13ZXN0LTE6ZGVmYXVsdDp1c2VyOmxuYXJkYWlAY2xvdWRlcmEuY29t"
	privateKey := "nHkdxgZR0BaNHaSYM3ooS6rIlpV5E+k1CIkr+jFId2g="
	ioutil.WriteFile(tempDirName+string(filepath.Separator)+common.Config_dir+string(filepath.Separator)+common.Config_file, []byte("default:\n  apikeyid: "+apiKey+"\n  privatekey: "+privateKey+"\n  server: server\n  workspace: workspace\n  output: output\n"), 0700)

	config, err := ReadConfig(tempDirName, "default")

	if err != nil {
		t.Errorf("unable to read file: %s", err.Error())
	}
	if config.Server != "server" {
		t.Errorf("server not match server == %s", config.Server)
	}
	if config.ApiKeyID != apiKey {
		t.Errorf("output not match output == %s", config.ApiKeyID)
	}
	if config.PrivateKey != privateKey {
		t.Errorf("output not match output == %s", config.PrivateKey)
	}
	if config.Output != "output" {
		t.Errorf("output not match output == %s", config.Output)
	}
	if config.Workspace != "workspace" {
		t.Errorf("workspace not match workspace == %s", config.Workspace)
	}
}
