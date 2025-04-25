package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/anibaldeboni/gen/cmd"
	"github.com/spf13/cobra"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	// Create a new buffer for each test to isolate output
	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)

	// Ensure arguments are isolated for each test
	c.SetArgs(args)

	// Log the command and arguments for debugging flaky tests
	t.Logf("Executing command: %s with args: %v", c.Name(), args)

	// Execute the command
	err := c.Execute()

	// Return trimmed output and error
	return strings.TrimSpace(buf.String()), err
}

func TestCpfCmd(t *testing.T) {
	t.Run("generate cpf", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cpf")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := len(out), 36; got != want {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})

	t.Run("valid cpf", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cpf", "validate", "74821506335")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 74821506335 is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("invalid cpf", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cpf", "validate", "12345678910")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🔴 12345678910 Invalid value"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}

func TestUuidCmd(t *testing.T) {
	t.Run("generate uuid", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "-r", "uuid")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) != 36 {
			t.Fatalf("expected UUID of length 36, got %d", len(out))
		}
	})

	t.Run("validate valid uuid", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validUuid := "123e4567-e89b-12d3-a456-426614174000"
		out, err := execute(t, rootCmd, "-r", "uuid", "validate", validUuid)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "true"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid uuid", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidUuid := "invalid-uuid"
		out, err := execute(t, rootCmd, "uuid", "validate", invalidUuid)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🔴 "+invalidUuid+" invalid UUID length: 12"; !strings.Contains(got, want) {
			t.Fatalf("expected %q to contain %q", got, want)
		}
	})
}

func TestEmailCmd(t *testing.T) {
	t.Run("generate email", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "email")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !strings.Contains(out, "@example.com") {
			t.Fatalf("expected email to contain domain @example.com, got %q", out)
		}
	})
}

func TestNameCmd(t *testing.T) {
	t.Run("generate name", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "name")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty name, got %q", out)
		}
	})
}

func TestCnpjCmd(t *testing.T) {
	t.Run("generate cnpj", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cnpj")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty CNPJ, got %q", out)
		}
	})

	t.Run("validate valid cnpj", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validCnpj := "11222333000181"
		out, err := execute(t, rootCmd, "cnpj", "validate", validCnpj)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 "+validCnpj+" is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid cnpj", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidCnpj := "11111111111111"
		out, err := execute(t, rootCmd, "-r", "cnpj", "validate", invalidCnpj)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "false"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}

func TestCnhCmd(t *testing.T) {
	t.Run("generate cnh", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cnh")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty CNH, got %q", out)
		}
	})

	t.Run("validate valid cnh", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validCnh := "07437082110"
		out, err := execute(t, rootCmd, "cnh", "validate", validCnh)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 "+validCnh+" is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid cnh", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidCnh := "00000000001"
		out, err := execute(t, rootCmd, "-r", "cnh", "validate", invalidCnh)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "false"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}

func TestRenavamCmd(t *testing.T) {
	t.Run("generate renavam", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "renavam")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty RENAVAM, got %q", out)
		}
	})

	t.Run("validate valid renavam", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validRenavam := "12345678900"
		out, err := execute(t, rootCmd, "renavam", "validate", validRenavam)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 "+validRenavam+" is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid renavam", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidRenavam := "00000000000"
		out, err := execute(t, rootCmd, "-r", "renavam", "validate", invalidRenavam)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "false"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}

func TestVoterCmd(t *testing.T) {
	t.Run("generate voter registration", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "voter")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty voter registration number, got %q", out)
		}
	})

	t.Run("validate valid voter registration", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validVoter := "056277020949"
		out, err := execute(t, rootCmd, "voter", "validate", validVoter)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 "+validVoter+" is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid voter registration", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidVoter := "000000000000"
		out, err := execute(t, rootCmd, "-r", "voter", "validate", invalidVoter)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "false"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}

func TestCcCmd(t *testing.T) {
	t.Run("generate credit card", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		out, err := execute(t, rootCmd, "cc")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(out) == 0 {
			t.Fatalf("expected a non-empty credit card number, got %q", out)
		}
	})

	t.Run("validate valid credit card", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		validCc := "4858114879119"
		out, err := execute(t, rootCmd, "cc", "validate", validCc)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "🟢 "+validCc+" is valid"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("validate invalid credit card", func(t *testing.T) {
		var rootCmd = cmd.RootCmd()

		invalidCc := "1234567890123456"
		out, err := execute(t, rootCmd, "-r", "cc", "validate", invalidCc)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if got, want := out, "false"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}
