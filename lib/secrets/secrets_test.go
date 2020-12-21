package secrets_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chef/automate/lib/io/fileutils"
	"github.com/chef/automate/lib/secrets"
)

func TestSecretNameFromString(t *testing.T) {
	t.Run("it parses chef_secrets style secrets specs", func(t *testing.T) {
		name, err := secrets.SecretNameFromString("group.secret_name")
		require.NoError(t, err)
		assert.Equal(t, secrets.SecretName{"group", "secret_name"}, name)
	})

	t.Run("it fails if too many parts are given", func(t *testing.T) {
		_, err := secrets.SecretNameFromString("group.secret_name.out")
		require.Error(t, err)
	})

	t.Run("it fails if too few parts are given", func(t *testing.T) {
		_, err := secrets.SecretNameFromString("secret_name")
		require.Error(t, err)
	})
}

func TestDiskSecretStore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "secret-store-tests")
	defer os.RemoveAll(tmpDir)

	require.NoError(t, err, "failed to create temporary directory")
	uid := os.Getuid()
	gid := os.Getgid()
	require.NoError(t, err, "failed to look up running user")
	dataDir := filepath.Join(tmpDir, "data-dir")

	dataStore := secrets.NewDiskStore(dataDir, uid, gid)
	err = dataStore.Initialize()
	require.NoError(t, err, "failed to initialize")

	t.Run("init creates the data directory with mode 0700", func(t *testing.T) {
		exists, _ := fileutils.PathExists(dataDir)
		assert.True(t, exists)

		stat, err := os.Stat(dataDir)
		require.NoError(t, err)
		assert.Equal(t, os.FileMode(os.ModeDir|0700), stat.Mode())
	})

	t.Run("set creates a secret with correct permissions", func(t *testing.T) {
		testSecretName := secrets.SecretName{"test", "test-secret-1"}
		testSecretContent := []byte("this-is-a-secret")
		err := dataStore.SetSecret(testSecretName, testSecretContent)
		require.NoError(t, err)

		stat, err := os.Stat(filepath.Join(dataDir, "test", "test-secret-1"+secrets.SecretFileExtension))
		require.NoError(t, err)
		assert.Equal(t, os.FileMode(0700), stat.Mode())
		assert.Equal(t, uint32(uid), stat.Sys().(*syscall.Stat_t).Uid)

	})

	t.Run("get and set can roundtrip a secret", func(t *testing.T) {
		testSecretName := secrets.SecretName{"test", "test-secret-1"}
		testSecretContent := []byte("this-is-a-secret")
		err := dataStore.SetSecret(testSecretName, testSecretContent)
		require.NoError(t, err)

		secretContent, err := dataStore.GetSecret(testSecretName)
		require.NoError(t, err)
		assert.Equal(t, testSecretContent, secretContent)
	})

	t.Run("Exists returns true if the secret exists", func(t *testing.T) {
		testSecretName := secrets.SecretName{"test", "test-secret-2"}
		testDNESecretName := secrets.SecretName{"test", "nothing-to-see-here"}
		testSecretContent := []byte("this-is-a-secret")
		err := dataStore.SetSecret(testSecretName, testSecretContent)
		require.NoError(t, err)

		exists, err := dataStore.Exists(testSecretName)
		require.NoError(t, err)
		assert.True(t, exists)

		exists, err = dataStore.Exists(testDNESecretName)
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

func TestSubPropertyBased(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 500
	properties := gopter.NewProperties(parameters)

	tmpDir, _ := ioutil.TempDir("", "secret-store-tests")
	defer os.RemoveAll(tmpDir)

	uid := os.Getuid()
	gid := os.Getgid()
	dataDir := filepath.Join(tmpDir, "data-dir")

	dataStore := secrets.NewDiskStore(dataDir, uid, gid)
	err := dataStore.Initialize()
	require.NoError(t, err, "failed to initialize")

	properties.Property("generate multiple random test cases to check if get and set can roundtrip a secret", prop.ForAll(
		func(a string, b string, c string) bool {
			testSecretName := secrets.SecretName{a, b}
			testSecretContent := []byte(c)
			err := dataStore.SetSecret(testSecretName, testSecretContent)
			if err != nil {
				return reportErrorAndYieldFalse(t, err)
			}
			secretContent, err := dataStore.GetSecret(testSecretName)
			if err != nil {
				return reportErrorAndYieldFalse(t, err)
			}
			return string(testSecretContent) == string(secretContent)
		},
		stringGen(10).WithLabel("a"),
		stringGen(10).WithLabel("b"),
		stringGen(10).WithLabel("c"),
	))
	properties.TestingRun(t)
}

// gopter helpers
func stringGen(size int) gopter.Gen {
	return gen.SliceOfN(size, gen.AlphaNumChar()).Map(func(r []rune) string {
		return string(r)
	}).WithShrinker(gen.StringShrinker)
}

func reportErrorAndYieldFalse(t *testing.T, err error) bool {
	t.Helper()
	t.Error(err.Error())
	return false
}
