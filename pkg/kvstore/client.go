// Copyright 2016-2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvstore

import "fmt"

var (
	// defaultClient is the default client initialized by initClient
	defaultClient BackendOperations
)

func initClient(module backendModule) error {
	c, errChan := module.newClient()

	defaultClient = c

	go func() {
		err, isErr := <-errChan
		if isErr {
			log.WithError(err).Fatalf("Unable to connect to kvstore")
		}

		deleteLegacyPrefixes()
	}()

	return nil
}

// Client returns the global kvstore client or nil if the client is not configured yet
func Client() BackendOperations {
	return defaultClient
}

// NewClient returns a new kvstore client based on the configuration
func NewClient(selectedBackend string, opts map[string]string) (BackendOperations, chan error) {
	// Channel used to report immediate errors, module.newClient will
	// create and return a different channel, caller doesn't need to know
	errChan := make(chan error, 1)
	defer close(errChan)

	module := getBackend(selectedBackend)
	if module == nil {
		errChan <- fmt.Errorf("unknown key-value store type %q. See cilium.link/err-kvstore for details", selectedBackend)
		return nil, errChan
	}

	if err := module.setConfig(opts); err != nil {
		errChan <- err
		return nil, errChan
	}

	return module.newClient()
}
