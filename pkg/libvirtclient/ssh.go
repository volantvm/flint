package libvirtclient

import (
	"fmt"
	"os"
)

// SetupSSHEnvironment configures environment variables for SSH authentication
// This is necessary for libvirt-go's SSH connections to work properly
func SetupSSHEnvironment(keyPath, knownHostsPath string) error {
	// Validate that the SSH key exists
	if keyPath != "" {
		if _, err := os.Stat(keyPath); os.IsNotExist(err) {
			return fmt.Errorf("SSH key file not found: %s", keyPath)
		}

		// Check key file permissions (should be 600 or 400)
		info, err := os.Stat(keyPath)
		if err != nil {
			return fmt.Errorf("failed to stat SSH key file: %w", err)
		}

		mode := info.Mode().Perm()
		// Warn if permissions are too open (not 600 or 400)
		if mode&0077 != 0 {
			return fmt.Errorf("SSH key file has insecure permissions (%o). Expected 600 or 400", mode)
		}
	}

	// Validate known_hosts file if specified
	if knownHostsPath != "" {
		if _, err := os.Stat(knownHostsPath); os.IsNotExist(err) {
			// Create an empty known_hosts file if it doesn't exist
			file, err := os.OpenFile(knownHostsPath, os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				return fmt.Errorf("failed to create known_hosts file: %w", err)
			}
			file.Close()
		}
	}

	return nil
}

// ValidateSSHConnection attempts to validate SSH connectivity parameters
// This can be called before attempting to connect to libvirt over SSH
func ValidateSSHConnection(username, host string, port int, keyPath string) error {
	if username == "" {
		return fmt.Errorf("SSH username cannot be empty")
	}

	if host == "" {
		return fmt.Errorf("SSH host cannot be empty")
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid SSH port: %d", port)
	}

	if keyPath == "" {
		return fmt.Errorf("SSH key path cannot be empty")
	}

	// Validate key file exists and has correct permissions
	if err := SetupSSHEnvironment(keyPath, ""); err != nil {
		return fmt.Errorf("SSH environment validation failed: %w", err)
	}

	return nil
}
